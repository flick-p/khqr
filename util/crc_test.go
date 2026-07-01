package util

import "testing"

func TestCalculateCRC16(t *testing.T) {
	tests := []struct {
		name string
		data string
		want string
	}{
		{
			name: "known KHQR payload",
			data: "00020101021129320016receivekhqr@dvpy0108123123125204599953038405802KH5904test6010Phnom Penh6304",
			want: "2356",
		},
		{
			name: "empty string",
			data: "",
			want: "FFFF",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateCRC16(tt.data); got != tt.want {
				t.Errorf("CalculateCRC16(%q) = %q, want %q", tt.data, got, tt.want)
			}
		})
	}
}

func TestCalculateCRC16_Deterministic(t *testing.T) {
	data := "0002010102122932001 6receivekhqr@dvpy"

	first := CalculateCRC16(data)
	second := CalculateCRC16(data)

	if first != second {
		t.Errorf("CalculateCRC16 is not deterministic: %q != %q", first, second)
	}
}

func TestCalculateCRC16_AlwaysFourHexDigits(t *testing.T) {
	inputs := []string{"", "a", "ab", "00020101021229"}

	for _, in := range inputs {
		got := CalculateCRC16(in)
		if len(got) != 4 {
			t.Errorf("CalculateCRC16(%q) = %q, want length 4", in, got)
		}
		for _, r := range got {
			isHex := (r >= '0' && r <= '9') || (r >= 'A' && r <= 'F')
			if !isHex {
				t.Errorf("CalculateCRC16(%q) = %q, contains non-uppercase-hex rune %q", in, got, r)
			}
		}
	}
}
