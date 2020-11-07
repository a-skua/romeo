package romeo

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestRequestImpl(t *testing.T) {
	r := &request{}

	testdata := []struct {
		req *http.Request
		err error
	}{
		{
			&http.Request{
				Body: ioutil.NopCloser(strings.NewReader(
					`{"id":"foo","name":"bar"}`,
				)),
			},
			errors.New("Request-Header has not Content-Type"),
		},
		{
			&http.Request{
				Body: ioutil.NopCloser(strings.NewReader(
					`{"id":"foo","name":"bar"}`,
				)),
				Header: func() http.Header {
					h := http.Header{}
					h.Set("Content-Type", "text/plain")
					return h
				}(),
			},
			errors.New("Content-Type(text/plain) is not Content-Type(application/json)"),
		},
		{
			&http.Request{
				Body: ioutil.NopCloser(strings.NewReader(
					`{"id":"foo","name":"bar"}`,
				)),
				Header: func() http.Header {
					h := http.Header{}
					h.Set("Content-Type", "application/json")
					return h
				}(),
			},
			nil,
		},
	}
	for i, d := range testdata {
		i += 1
		val := &struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		}{}

		if err := r.Read(d.req, val);
		// err is not d.err
		(err != nil && d.err != nil && err.Error() != d.err.Error()) ||
			// !(err == d.err == nil)
			(err != nil && d.err == nil) || (err == nil && d.err != nil) {
			t.Fatalf("%02d: have (%v), want (%v)", i, err, d.err)
		}

		// read has not error
		if d.err == nil {
			if have, want := val.ID, "foo"; have != want {
				t.Fatalf("%02d: have (%v), want (%v)", i, have, want)
			} else if have, want := val.Name, "bar"; have != want {
				t.Fatalf("%02d: have (%v), want (%v)", i, have, want)
			}
		}
	}
}

func TestNewRequestReader(t *testing.T) {
	if r := NewRequestReader(&RequestReaderConfigs{}); r == nil {
		t.Fatal("this is not nil")
	}
}
