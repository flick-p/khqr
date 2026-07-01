package generator

import (
	"khqr/constants"
	"khqr/models"
)

// this base code create only for merchant code that have only length validation

type baseMerchantCode struct {
	code               KHQRCodeDict
	value              *string
	isRequiredValidate bool
}

func newBaseMerchantCode(code KHQRCodeDict, v *string, isValidateRequiredInput ...bool) KHQRBuilder {

	var isRequiredValidate bool

	if len(isValidateRequiredInput) > 0 {
		isRequiredValidate = isValidateRequiredInput[0]
	}

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

func (b *baseMerchantCode) Validate() *constants.ErrorCode {

	if b.isRequiredValidate && b.value == nil && b.code.ErrRequire != nil {
		return b.code.ErrRequire
	}

	return ValidateLength(b.code, b.value)
}
