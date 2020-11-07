package romeo

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/a-skua/romeo/result"
	"github.com/a-skua/romeo/status"
	"github.com/a-skua/romeo/status/httpstatus"
)

func TestResponseWrite(t *testing.T) {
	w := httptest.NewRecorder()
	res := &response{
		log: log.New(os.Stdout, "", 0),
		status2http: func(s status.Status) int {
			return 200
		},
		wrapper:   DefaultWrap,
		setHeader: DefaultSetHeader,
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

func TestResponseWriteWithCustomSetHeaderFunc(t *testing.T) {
	res := &response{
		log: log.New(os.Stdout, "", 0),
		status2http: func(s status.Status) int {
			return 200
		},
		wrapper: DefaultWrap,
		setHeader: func(h http.Header) {
			h.Set("content-type", "application/json; charset=utf-8")
			h.Set("cache-control", "no-store")
			h.Set("pragma", "no-cache")
		},
	}
	w := httptest.NewRecorder()
	res.Write(w, result.New(httpstatus.OK(), nil, nil))
	r := w.Result()
	if h, w := r.Header.Get("Content-Type"), "application/json; charset=utf-8"; h != w {
		t.Errorf("have (%#v), want (%#v)", h, w)
	}
	if h, w := r.Header.Get("Cache-Control"), "no-store"; h != w {
		t.Errorf("have (%#v), want (%#v)", h, w)
	}
	if h, w := r.Header.Get("Pragma"), "no-cache"; h != w {
		t.Errorf("have (%#v), want (%#v)", h, w)
	}
}

func TestNewResponseWriter(t *testing.T) {
	if r := NewResponseWriter(&ResponseWriterConfigs{nil, nil, nil, nil}); r == nil {
		t.Fatal("this is not nil")
	}
}
