package generator

import (
	"strconv"

	"khqr/constants"
	"khqr/models"
	"khqr/util"
)

type transactionCurrency struct {
	value int
}

func NewTransactionCurrency(v int) KHQRBuilder {
	return &transactionCurrency{
		value: v,
	}
}

func (c *transactionCurrency) String() string {

	return models.NewTagLengthValue(transactionCurrencyCD.Tag, util.Ptr(strconv.Itoa(c.value))).ToString()
}

func (c *transactionCurrency) Validate() *constants.ErrorCode {

	if c.value == 0 {
		return &constants.ErrCurrencyTypeRequired
	}

	if len(strconv.Itoa(c.value)) > 3 {
		return &constants.ErrTransactionCurrencyLengthInvalid
	}

	if c.value != constants.CurrencyKHR && c.value != constants.CurrencyUSD {
		return &constants.ErrUnsupportedCurrency
	}

	return nil
}
