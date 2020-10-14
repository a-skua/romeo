package romeo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/a-skua/romeo/container"
)

// A RequestReader is the interface
// that wraps Read method to read Request.Body
type RequestReader interface {
	Read(*http.Request, interface{}) error
}

// a request is implements the RequestReader
type request struct{}

// A Read is read from request body
func (request) Read(r *http.Request, v interface{}) error {
	if have := r.Header.Get("Content-Type"); len(have) == 0 {
		return errors.New("Request-Header has not Content-Type")
	} else if want := "application/json"; have != want {
		return errors.New(fmt.Sprintf("Content-Type(%s) is not Content-Type(%s)", have, want))
	}
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, v)
}

// ResponseWriter is the interface
// that wraps Write method to write Response
type ResponseWriter interface {
	Write(http.ResponseWriter, container.Result)
}
