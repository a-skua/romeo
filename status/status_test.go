package status

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestStatusInt(t *testing.T) {
	testdata := []int{
		rand.Int(),
		rand.Int(),
		rand.Int(),
	}

	for i, data := range testdata {
		i++
		r := status{data}
		if h, w := r.Int(), data; h != w {
			t.Error(i, fmt.Sprintf("have (%#v), want (%#v)", h, w))
		}
	}
}

func TestHttpStatusInt(t *testing.T) {
	testdata := []int{
		rand.Int(),
		rand.Int(),
		rand.Int(),
	}

	for i, data := range testdata {
		i++
		r := httpStatus{code: data}
		if h, w := r.Int(), data; h != w {
			t.Error(i, fmt.Sprintf("have (%#v), want (%#v)", h, w))
		}
	}
}

func TestHttpStatusHttp(t *testing.T) {
	testdata := []int{
		rand.Int(),
		rand.Int(),
		rand.Int(),
	}

	for i, data := range testdata {
		i++
		r := httpStatus{http: data}
		if h, w := r.Http(), data; h != w {
			t.Error(i, fmt.Sprintf("have (%#v), want (%#v)", h, w))
		}
	}
}
