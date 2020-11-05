package result

import (
	"errors"
	"fmt"
	"testing"

	"github.com/a-skua/romeo/status"
)

func TestResultData(t *testing.T) {
	testdata := []interface{}{
		nil,
		"text",
		&struct {
			text string
		}{"foo text"},
	}

	for i, data := range testdata {
		i += 1
		r := &result{
			data: data,
		}
		if h, w := r.Data(), data; h != w {
			t.Error(i, fmt.Sprintf("have (%#v), want (%#v)", h, w))
		}
	}
}

func TestResultStatus(t *testing.T) {
	testdata := []status.Status{
		status.New(1),
		status.New(2),
		status.New(4),
	}

	for i, data := range testdata {
		i += 1
		r := &result{
			status: data,
		}
		if h, w := r.Status(), data; h != w {
			t.Error(i, fmt.Sprintf("have (%#v), want (%#v)", h, w))
		}
	}
}

func TestResultError(t *testing.T) {
	testdata := []error{
		nil,
		errors.New("test error"),
	}
	for i, data := range testdata {
		i += 1
		r := &result{
			err: data,
		}
		if h, w := r.Error(), data; h != w {
			t.Error(i, fmt.Sprintf("have (%#v), want (%#v)", h, w))
		}
	}
}

func TestNew(t *testing.T) {
	s, d, e := status.New(1), "test data", errors.New("test error")
	r := New(s, d, e)

	if h, w := r.Status(), s; h != w {
		t.Error("Status", fmt.Sprintf("have (%#v), want (%#v)", h, w))
	}

	if h, w := r.Data(), d; h != w {
		t.Error("Data", fmt.Sprintf("have (%#v), want (%#v)", h, w))
	}

	if h, w := r.Error(), e; h != w {
		t.Error("Error", fmt.Sprintf("have (%#v), want (%#v)", h, w))
	}
}
