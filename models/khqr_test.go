package models

import "testing"

func TestTagLengthValue_ToString(t *testing.T) {
	str := func(s string) *string { return &s }

	tests := []struct {
		name  string
		tag   string
		value *string
		want  string
	}{
		{"nil value produces empty string", "59", nil, ""},
		{"empty value produces empty string", "59", str(""), ""},
		{"ASCII value encodes tag and 2-digit length", "59", str("test"), "5904test"},
		{"length is zero-padded", "60", str("KH"), "6002KH"},
		{"multi-byte runes counted as runes not bytes", "59", str("ភ្នំពេញ"), "5907ភ្នំពេញ"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tlv := NewTagLengthValue(tt.tag, tt.value)
			if got := tlv.ToString(); got != tt.want {
				t.Errorf("ToString() = %q, want %q", got, tt.want)
			}
		})
	}
}
