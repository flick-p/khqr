package generator

import (
	"khqr/models"
)

// this base code create only for merchant code that have only length validation

type baseMerchantCode struct {
	code               KHQRCodeDict
	value              *string
	isRequiredValidate bool
}

func newBaseMerchantCode(code KHQRCodeDict, v *string, isRequiredValidate bool) KHQRBuilder {

	return &baseMerchantCode{
		code:               code,
		value:              v,
		isRequiredValidate: isRequiredValidate,
	}
}

// convert value to string
func (b *baseMerchantCode) String() string {
	return models.NewTagLengthValue(b.code.Tag, b.value).ToString()
}

// validation

func (b *baseMerchantCode) Validate() error {

	if b.isRequiredValidate && b.value == nil && b.code.ErrRequire != nil {
		return b.code.ErrRequire
	}

	if err := ValidateLength(b.code, b.value); err != nil {
		return err
	}

	return nil
}
