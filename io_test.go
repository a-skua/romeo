package romeo

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/a-skua/romeo/result"
	"github.com/a-skua/romeo/status"
	"github.com/a-skua/romeo/status/httpstatus"
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

func TestResponseWrite(t *testing.T) {
	w := httptest.NewRecorder()
	res := &response{
		log: log.New(os.Stdout, "", 0),
		status2http: func(s status.Status) int {
			return 200
		},
		wrapper: DefaultWrap,
	}

	// status ok
	res.Write(w, result.New(httpstatus.OK(), struct {
		User string `json:"user"`
	}{"romeo"}, nil))
	r := w.Result()
	if body, err := ioutil.ReadAll(r.Body); err != nil {
		t.Fatal(err)
	} else {
		if h, w := r.StatusCode, 200; h != w {
			t.Errorf("have (%#v), want (%#v)", h, w)
		}
		if h, w := string(body), `{"data":{"user":"romeo"}}`; h != w {
			t.Errorf("have (%#v), want (%#v)", h, w)
		}
		if h, w := r.Header.Get("Content-Type"), "application/json"; h != w {
			t.Errorf("have (%#v), want (%#v)", h, w)
		}
	}

	// status internal
	w = httptest.NewRecorder()
	res.Write(w, result.New(httpstatus.InternalError(), nil, errors.New("test error")))
	// Output: test error
	r = w.Result()
	if body, err := ioutil.ReadAll(r.Body); err != nil {
		t.Fatal(err)
	} else {
		if h, w := r.StatusCode, 500; h != w {
			t.Errorf("have (%#v), want (%#v)", h, w)
		}
		if h, w := string(body), `{"data":null}`; h != w {
			t.Errorf("have (%#v), want (%#v)", h, w)
		}
		if h, w := r.Header.Get("Content-Type"), "application/json"; h != w {
			t.Errorf("have (%#v), want (%#v)", h, w)
		}
	}

	// error
	w = httptest.NewRecorder()
	res = nil
	res.Write(w, result.New(httpstatus.OK(), nil, nil))
	r = w.Result()
	if body, err := ioutil.ReadAll(r.Body); err != nil {
		t.Fatal(err)
	} else {
		if h, w := r.StatusCode, 500; h != w {
			t.Errorf("have (%#v), want (%#v)", h, w)
		}
		if h, w := string(body), "Internal Server Error\n"; h != w {
			t.Errorf("have (%#v), want (%#v)", h, w)
		}
		if h, w := r.Header.Get("Content-Type"), "text/plain; charset=utf-8"; h != w {
			t.Errorf("have (%#v), want (%#v)", h, w)
		}
	}
}

func TestNewResponseWriter(t *testing.T) {
	if r := NewResponseWriter(&ResponseWriterConfigs{nil, nil, nil}); r == nil {
		t.Fatal("this is not nil")
	}
}
