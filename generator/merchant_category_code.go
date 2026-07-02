package generator

import (
	"regexp"
	"strconv"

	"khqr/constants"
	"khqr/models"
)

type merchantCategoryCode struct {
	value *string
}

func NewMerchantCategoryCode(v *string) KHQRBuilder {
	return &merchantCategoryCode{
		value: v,
	}
}

func (m *merchantCategoryCode) String() string {
	return models.NewTagLengthValue(merchantCatCodeCD.Tag, m.value).ToString()
}

func (m *merchantCategoryCode) Validate() error {

	err := newBaseMerchantCode(merchantCatCodeCD, m.value, true).Validate()
	if err != nil {
		return err
	}

	// Check if it's numeric
	if matched, _ := regexp.MatchString(`^\d+$`, *m.value); !matched {
		return &constants.ErrInvalidMerchantCategoryCode
	}

	// Check range
	if mcc, err := strconv.Atoi(*m.value); err != nil || mcc < 0 || mcc > 9999 {
		return &constants.ErrInvalidMerchantCategoryCode
	}

	return nil
}
