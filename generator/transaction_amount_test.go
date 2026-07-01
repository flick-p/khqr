package generator

import (
	"testing"

	"khqr/constants"
	"khqr/util"
)

func TestTransactionAmount_String(t *testing.T) {
	tests := []struct {
		name     string
		amount   *string
		currency int
		want     string
	}{
		{"nil amount produces empty string", nil, constants.CurrencyUSD, ""},
		{"zero amount produces empty string", util.Ptr("0"), constants.CurrencyUSD, ""},
		{"empty amount produces empty string", util.Ptr(""), constants.CurrencyUSD, ""},
		{"USD amount formatted to 2 decimals", util.Ptr("10"), constants.CurrencyUSD, "540510.00"},
		{"USD amount with decimals preserved", util.Ptr("10.5"), constants.CurrencyUSD, "540510.50"},
		{"KHR amount left as-is", util.Ptr("5000"), constants.CurrencyKHR, "54045000"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewTransactionAmount(tt.amount, tt.currency)
			if got := a.String(); got != tt.want {
				t.Errorf("String() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestTransactionAmount_Validate(t *testing.T) {
	tests := []struct {
		name     string
		amount   *string
		currency int
		wantErr  bool
	}{
		{"nil amount is valid", nil, constants.CurrencyUSD, false},
		{"empty amount is valid", util.Ptr(""), constants.CurrencyUSD, false},
		{"positive USD amount is valid", util.Ptr("10.50"), constants.CurrencyUSD, false},
		{"positive KHR integer amount is valid", util.Ptr("5000"), constants.CurrencyKHR, false},
		{"zero amount is invalid", util.Ptr("0"), constants.CurrencyUSD, true},
		{"negative amount is invalid", util.Ptr("-5"), constants.CurrencyUSD, true},
		{"non-numeric amount is invalid", util.Ptr("abc"), constants.CurrencyUSD, true},
		{"KHR amount with decimals is invalid", util.Ptr("50.5"), constants.CurrencyKHR, true},
		{"USD amount with more than 2 decimals is invalid", util.Ptr("10.555"), constants.CurrencyUSD, true},
		{"amount longer than max length is invalid", util.Ptr("123456789012345"), constants.CurrencyUSD, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewTransactionAmount(tt.amount, tt.currency)
			err := a.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTransactionCurrency(t *testing.T) {
	t.Run("String encodes numeric currency code", func(t *testing.T) {
		c := NewTransactionCurrency(constants.CurrencyUSD)
		if got, want := c.String(), "5303840"; got != want {
			t.Errorf("String() = %q, want %q", got, want)
		}
	})

	tests := []struct {
		name    string
		value   int
		wantErr bool
	}{
		{"zero currency is invalid", 0, true},
		{"unsupported currency is invalid", 999, true},
		{"USD is valid", constants.CurrencyUSD, false},
		{"KHR is valid", constants.CurrencyKHR, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewTransactionCurrency(tt.value)
			err := c.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
