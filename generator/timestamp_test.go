package generator

import (
	"testing"
	"time"

	"khqr/util"
)

func TestTimestamp_String(t *testing.T) {
	t.Run("nil expiration produces empty string", func(t *testing.T) {
		ts := NewTimestamp(nil)
		if got := ts.String(); got != "" {
			t.Errorf("String() = %q, want empty string", got)
		}
	})

	t.Run("set expiration produces creation and expiration TLVs", func(t *testing.T) {
		exp := time.Now().Add(time.Hour).UnixMilli()
		ts := NewTimestamp(&exp)

		got := ts.String()
		if got == "" {
			t.Fatal("String() = empty string, want non-empty")
		}
		if got[:2] != "99" {
			t.Errorf("String() tag = %q, want %q", got[:2], "99")
		}
	})
}

func TestTimestamp_Validate(t *testing.T) {
	now := time.Now().UnixMilli()

	tests := []struct {
		name    string
		expTime *int64
		wantErr bool
	}{
		{"nil expiration is valid", nil, false},
		{"zero expiration is invalid", util.Ptr(int64(0)), true},
		{"expiration before creation is invalid", util.Ptr(now - 1000), true},
		{"expiration wrong digit length is invalid", util.Ptr(int64(1234567890)), true},
		{"future expiration is valid", util.Ptr(time.Now().Add(time.Hour).UnixMilli()), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := NewTimestamp(tt.expTime)
			err := ts.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPointOfInitMtd_Validate(t *testing.T) {
	tests := []struct {
		name    string
		value   *string
		wantErr bool
	}{
		{"nil value is valid", nil, false},
		{"static QR is valid", util.Ptr("11"), false},
		{"dynamic QR is valid", util.Ptr("12"), false},
		{"unknown value is invalid", util.Ptr("99"), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewPointOfInitMtd(tt.value)
			err := p.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMerchantCategoryCode_Validate(t *testing.T) {
	tests := []struct {
		name    string
		value   *string
		wantErr bool
	}{
		{"nil value is invalid (required)", nil, true},
		{"valid numeric code", util.Ptr("5999"), false},
		{"non-numeric code is invalid", util.Ptr("abcd"), true},
		{"out-of-range code is invalid", util.Ptr("99999"), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMerchantCategoryCode(tt.value)
			err := m.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
