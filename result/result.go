package result

import (
	"github.com/a-skua/romeo/status"
)

// A Result is ther interface
// that wraps Data method to tell response
type Result interface {
	Status() status.Status
	Data() interface{}
	Error() error
}

// New return a Result
func New(s status.Status, data interface{}, err error) Result {
	return &result{
		status: s,
		data:   data,
		err:    err,
	}
}

type result struct {
	status status.Status
	data   interface{}
	err    error
}

// Data return a response data
func (r *result) Data() interface{} {
	return r.data
}

// Status return a status.Status
func (r *result) Status() status.Status {
	return r.status
}

// Error return an error
func (r *result) Error() error {
	return r.err
}
