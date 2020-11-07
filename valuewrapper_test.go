package romeo

import (
	"testing"
)

func TestDefaultWrap(t *testing.T) {
	testdata := []interface{}{
		"foo",
		1,
		struct{}{},
	}
	for i, data := range testdata {
		i++
		d, ok := DefaultWrap(&data).(*defaultValue)
		if !ok {
			t.Fatalf("return value can't cast to *defaultValue")
		}
		if h, w := d.Data, &data; h != w {
			t.Errorf("%02d have (%#v), want (%#v)", i, h, w)
		}
	}
}

func TestNoWrap(t *testing.T) {
	testdata := []interface{}{
		"foo",
		1,
		struct{}{},
	}
	for i, data := range testdata {
		i++
		if h, w := NoWrap(&data), &data; h != w {
			t.Errorf("%02d have (%#v), want (%#v)", i, h, w)
		}
	}
}
