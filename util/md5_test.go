package util

import "testing"

func TestCalculateMD5(t *testing.T) {
	tests := []struct {
		name string
		data string
		want string
	}{
		{"empty string", "", "d41d8cd98f00b204e9800998ecf8427e"},
		{"known string", "hello", "5d41402abc4b2a76b9719d911017c592"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateMD5(tt.data); got != tt.want {
				t.Errorf("CalculateMD5(%q) = %q, want %q", tt.data, got, tt.want)
			}
		})
	}
}

func TestCalculateMD5_Deterministic(t *testing.T) {
	data := "00020101021129320016receivekhqr@dvpy"

	if CalculateMD5(data) != CalculateMD5(data) {
		t.Error("CalculateMD5 is not deterministic")
	}
}
