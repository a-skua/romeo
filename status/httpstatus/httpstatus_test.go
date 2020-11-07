package httpstatus

import (
	"fmt"
	"net/http"
	"testing"
)

func TestOK(t *testing.T) {
	status := OK()

	if h, w := status.Http(), http.StatusOK; h != w {
		t.Error(fmt.Sprintf("have (%#v), want (%#v)", h, w))
	}

	if h, w := status.Int(), StatusOK; h != w {
		t.Error(fmt.Sprintf("have (%#v), want (%#v)", h, w))
	}
}

func TestBadRequest(t *testing.T) {
	status := BadRequest()

	if h, w := status.Http(), http.StatusBadRequest; h != w {
		t.Error(fmt.Sprintf("have (%#v), want (%#v)", h, w))
	}

	if h, w := status.Int(), StatusBadRequest; h != w {
		t.Error(fmt.Sprintf("have (%#v), want (%#v)", h, w))
	}
}

func TestUnAuthorized(t *testing.T) {
	status := UnAuthorized()

	if h, w := status.Http(), http.StatusUnauthorized; h != w {
		t.Error(fmt.Sprintf("have (%#v), want (%#v)", h, w))
	}

	if h, w := status.Int(), StatusUnAuthorized; h != w {
		t.Error(fmt.Sprintf("have (%#v), want (%#v)", h, w))
	}
}

func TestForbidden(t *testing.T) {
	status := Forbidden()

	if h, w := status.Http(), http.StatusForbidden; h != w {
		t.Error(fmt.Sprintf("have (%#v), want (%#v)", h, w))
	}

	if h, w := status.Int(), StatusForbidden; h != w {
		t.Error(fmt.Sprintf("have (%#v), want (%#v)", h, w))
	}
}

func TestNotFound(t *testing.T) {
	status := NotFound()

	if h, w := status.Http(), http.StatusNotFound; h != w {
		t.Error(fmt.Sprintf("have (%#v), want (%#v)", h, w))
	}

	if h, w := status.Int(), StatusNotFound; h != w {
		t.Error(fmt.Sprintf("have (%#v), want (%#v)", h, w))
	}
}

func TestTeapot(t *testing.T) {
	status := Teapot()

	if h, w := status.Http(), http.StatusTeapot; h != w {
		t.Error(fmt.Sprintf("have (%#v), want (%#v)", h, w))
	}

	if h, w := status.Int(), StatusTeapot; h != w {
		t.Error(fmt.Sprintf("have (%#v), want (%#v)", h, w))
	}
}

func TestInternalError(t *testing.T) {
	status := InternalError()

	if h, w := status.Http(), http.StatusInternalServerError; h != w {
		t.Error(fmt.Sprintf("have (%#v), want (%#v)", h, w))
	}

	if h, w := status.Int(), StatusInternalError; h != w {
		t.Error(fmt.Sprintf("have (%#v), want (%#v)", h, w))
	}
}
