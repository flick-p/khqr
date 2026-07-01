package generator

import (
	"testing"

	"khqr/constants"
	"khqr/util"
)

func TestValidateLength(t *testing.T) {
	code := KHQRCodeDict{
		Tag:              "01",
		MaxLength:        5,
		ErrInvalidLength: constants.ErrMerchantNameLengthInvalid,
	}

	tests := []struct {
		name    string
		value   *string
		wantErr *constants.ErrorCode
	}{
		{"nil value is valid", nil, nil},
		{"value within max length is valid", util.Ptr("abcde"), nil},
		{"value exceeding max length is invalid", util.Ptr("abcdef"), &constants.ErrMerchantNameLengthInvalid},
		{"empty value is valid", util.Ptr(""), nil},
		{"multi-byte runes counted as runes not bytes", util.Ptr("សួស្"), nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidateLength(code, tt.value)

			if (got == nil) != (tt.wantErr == nil) {
				t.Fatalf("ValidateLength() = %v, want %v", got, tt.wantErr)
			}
			if got != nil && got.Code != tt.wantErr.Code {
				t.Errorf("ValidateLength() code = %d, want %d", got.Code, tt.wantErr.Code)
			}
		})
	}
}

func TestIsValidRequiredStrValue(t *testing.T) {
	tests := []struct {
		name  string
		value *string
		want  bool
	}{
		{"nil pointer is invalid", nil, false},
		{"empty string is invalid", util.Ptr(""), false},
		{"non-empty string is valid", util.Ptr("value"), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidRequiredStrValue(tt.value); got != tt.want {
				t.Errorf("IsValidRequiredStrValue(%v) = %v, want %v", tt.value, got, tt.want)
			}
		})
	}
}
