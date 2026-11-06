package ecl

import (
	"reflect"
	"testing"
)

func TestVRIDAsState(t *testing.T) {
	cases := []struct {
		vrid     interface{}
		expected string
	}{
		{vrid: nil, expected: "null"},
		{vrid: float64(123), expected: "123"},
		{vrid: 123, expected: "123"},
		{vrid: "123", expected: "123"},
		{vrid: "", expected: "null"},
	}

	for i, c := range cases {
		if actual := vridAsState(c.vrid); actual != c.expected {
			t.Errorf("case %d: expected %q, got %q", i, c.expected, actual)
		}
	}
}

func TestVRIDFromState(t *testing.T) {
	cases := []struct {
		vrid      string
		expected  interface{}
		expectErr bool
	}{
		{vrid: "null", expected: nil},
		{vrid: "123", expected: 123},
		{vrid: "abc", expectErr: true},
	}

	for i, c := range cases {
		actual, err := vridFromState(c.vrid)
		if c.expectErr {
			if err == nil {
				t.Errorf("case %d: expected error, got nil", i)
			}
			continue
		}
		if err != nil {
			t.Errorf("case %d: unexpected error: %s", i, err)
		}
		if !reflect.DeepEqual(actual, c.expected) {
			t.Errorf("case %d: expected %#v, got %#v", i, c.expected, actual)
		}
	}
}
