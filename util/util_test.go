package util

import "testing"

func TestPtr(t *testing.T) {
	v := Ptr("hello")

	if v == nil {
		t.Fatal("Ptr() returned nil")
	}
	if *v != "hello" {
		t.Errorf("*Ptr(%q) = %q, want %q", "hello", *v, "hello")
	}
}

func TestSafePtrDeref(t *testing.T) {
	t.Run("non-nil pointer returns pointed-to value", func(t *testing.T) {
		v := "hello"
		if got := SafePtrDeref(&v); got != "hello" {
			t.Errorf("SafePtrDeref(&v) = %q, want %q", got, "hello")
		}
	})

	t.Run("nil pointer returns zero value", func(t *testing.T) {
		var v *string
		if got := SafePtrDeref(v); got != "" {
			t.Errorf("SafePtrDeref(nil) = %q, want empty string", got)
		}
	})

	t.Run("nil int pointer returns zero value", func(t *testing.T) {
		var v *int
		if got := SafePtrDeref(v); got != 0 {
			t.Errorf("SafePtrDeref(nil) = %d, want 0", got)
		}
	})
}
