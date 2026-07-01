package util

import "testing"

func TestCutString(t *testing.T) {
	t.Run("valid tag-length-value", func(t *testing.T) {
		tag, value, remaining, err := CutString("000201rest")
		if err != nil {
			t.Fatalf("CutString() error = %v", err)
		}
		if tag != "00" {
			t.Errorf("tag = %q, want %q", tag, "00")
		}
		if value != "01" {
			t.Errorf("value = %q, want %q", value, "01")
		}
		if remaining != "rest" {
			t.Errorf("remaining = %q, want %q", remaining, "rest")
		}
	})

	t.Run("consumes entire string with no remainder", func(t *testing.T) {
		tag, value, remaining, err := CutString("5904test")
		if err != nil {
			t.Fatalf("CutString() error = %v", err)
		}
		if tag != "59" {
			t.Errorf("tag = %q, want %q", tag, "59")
		}
		if value != "test" {
			t.Errorf("value = %q, want %q", value, "test")
		}
		if remaining != "" {
			t.Errorf("remaining = %q, want empty string", remaining)
		}
	})

	t.Run("multi-byte runes counted as runes not bytes", func(t *testing.T) {
		// "αβ" is 2 runes but 4 bytes in UTF-8; the length prefix "02" must
		// be interpreted as a rune count for this to slice correctly.
		tag, value, remaining, err := CutString("0002αβrest")
		if err != nil {
			t.Fatalf("CutString() error = %v", err)
		}
		if tag != "00" {
			t.Errorf("tag = %q, want %q", tag, "00")
		}
		if value != "αβ" {
			t.Errorf("value = %q, want %q", value, "αβ")
		}
		if remaining != "rest" {
			t.Errorf("remaining = %q, want %q", remaining, "rest")
		}
	})

	t.Run("string too short for tag+length header errors", func(t *testing.T) {
		_, _, _, err := CutString("abc")
		if err == nil {
			t.Error("CutString() error = nil, want error")
		}
	})

	t.Run("non-numeric length errors", func(t *testing.T) {
		_, _, _, err := CutString("00XX1234")
		if err == nil {
			t.Error("CutString() error = nil, want error")
		}
	})

	t.Run("declared length exceeds available data errors", func(t *testing.T) {
		_, _, _, err := CutString("0099short")
		if err == nil {
			t.Error("CutString() error = nil, want error")
		}
	})

	t.Run("empty string errors", func(t *testing.T) {
		_, _, _, err := CutString("")
		if err == nil {
			t.Error("CutString() error = nil, want error")
		}
	})
}
