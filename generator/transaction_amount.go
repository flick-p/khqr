package generator

import (
	"fmt"
	"strconv"
	"strings"

	"khqr/constants"
	"khqr/models"
)

type transactionAmount struct {
	amount   *string
	currency int
}

func NewTransactionAmount(amount *string, currency int) KHQRBuilder {
	return &transactionAmount{
		amount:   amount,
		currency: currency,
	}
}

func (t *transactionAmount) String() string {

	if t.amount == nil || *t.amount == "0" || *t.amount == "" {
		return ""
	}

	amount := *t.amount

	if t.currency == constants.CurrencyUSD {
		// Format to 2 decimal places for USD
		if amountFloat, err := strconv.ParseFloat(amount, 64); err == nil {
			amount = fmt.Sprintf("%.2f", amountFloat)
		}
	}

	return models.NewTagLengthValue(transactionAmtCD.Tag, &amount).ToString()
}

func (t *transactionAmount) Validate() *constants.ErrorCode {

	if !t.isValidAmount() {
		return &constants.ErrTransactionAmountInvalid
	}

	return nil
}

func (t *transactionAmount) isValidAmount() bool {

	if t.amount == nil || *t.amount == "" {
		return true
	}
	currency := t.currency
	amount := *t.amount

	if len(amount) > constants.MaxTransactionAmountLength || strings.Contains(amount, "-") {
		return false
	}

	amountFloat, err := strconv.ParseFloat(amount, 64)
	if err != nil || amountFloat <= 0 {
		return false
	}

	// For KHR, only integers are allowed
	if currency == constants.CurrencyKHR {
		if amountFloat != float64(int64(amountFloat)) {
			return false
		}
	} else {
		// For USD, check decimal precision
		parts := strings.Split(amount, ".")
		if len(parts) > 1 && len(parts[1]) > 2 {
			return false
		}
	}

	return true
}
