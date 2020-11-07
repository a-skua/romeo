package status

type Status2HTTPConverter func(s Status) int

// A Status is the interface
type Status interface {
	Int() int
}

// A HttpStatus is the interface
// that extends Status interface with Http Method
type HttpStatus interface {
	Status
	Http() int
}

// New return Status
func New(code int) Status {
	return &status{code}
}

type status struct {
	code int
}

// Int return status 2 int
func (s *status) Int() int {
	return s.code
}

// NewHttpStatus return HttpStatus
func NewHttpStatus(code, http int) HttpStatus {
	return &httpStatus{code, http}
}

type httpStatus struct {
	code int
	http int
}

// Int return status 2 int
func (s *httpStatus) Int() int {
	return s.code
}

// Http return http status code
func (s *httpStatus) Http() int {
	return s.http
}
