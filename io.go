package romeo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/a-skua/romeo/result"
	"github.com/a-skua/romeo/status"
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

// RequestReaderConfigs is argument by NewRequestReader
type RequestReaderConfigs struct{}

// NewRequestReader return A RequestReader
func NewRequestReader(conf *RequestReaderConfigs) RequestReader {
	return &request{}
}

// ResponseWriter is the interface
// that wraps Write method to write Response
type ResponseWriter interface {
	Write(http.ResponseWriter, result.Result)
}

// a response is implements the ResponseWriter
type response struct {
	status2http status.Status2HTTPConverter
	log         *log.Logger
	wrapper     ResponseValueWrapper
}

// A Write is written to response body
func (r *response) Write(w http.ResponseWriter, res result.Result) {
	if r == nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	if err := res.Error(); err != nil {
		log.Println(err.Error())
	}

	val := r.wrapper(res.Data())
	body, err := json.Marshal(val)
	if err != nil {
		r.log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
		return
	}

	// TODO: outsourcing
	w.Header().Set("Content-Type", "application/json")

	// write http status
	h, ok := res.Status().(status.HttpStatus)
	if ok {
		w.WriteHeader(h.Http())
	} else {
		w.WriteHeader(r.status2http(res.Status()))
	}

	w.Write(body)
}

// ResponseWriterConfigs is argument by NewResponseWriter
type ResponseWriterConfigs struct {
	logger     *log.Logger
	statusConv status.Status2HTTPConverter
	wrapper    ResponseValueWrapper
}

// NewResponseWriter return A ResponseWriter
func NewResponseWriter(conf *ResponseWriterConfigs) ResponseWriter {
	if conf == nil {
		conf = &ResponseWriterConfigs{}
	}

	if conf.logger == nil {
		conf.logger = log.New(os.Stderr, "", log.LstdFlags|log.LUTC|log.Lmicroseconds)
	}
	if conf.statusConv == nil {
		// TODO: anything
	}
	if conf.wrapper == nil {
		conf.wrapper = NoWrap
	}
	return &response{
		status2http: conf.statusConv,
		log:         conf.logger,
		wrapper:     conf.wrapper,
	}
}
