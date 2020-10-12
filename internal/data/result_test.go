package data

import (
	"testing"
)

func TestResultTypeBothFields(t *testing.T) {
	cases := []struct {
		in, want Result
	}{
		{Result{1, "hi", "12"}, Result{1, "hi", "12"}},
		{Result{1, "test2", "12"}, Result{1, "test2", "12"}},
	}

	for _, c := range cases {
		got := Result{Index: c.in.Index, Result: c.in.Result}

		if got.Index != c.want.Index {
			t.Errorf("Index in == %q, got: %q want %q", c.in.Index, got, c.want.Index)
		}

		if got.Result != c.want.Result {
			t.Errorf("Index in == %q, got: %q want %q", c.in.Index, got, c.want.Index)
		}
	}
}
