package generator

import (
	"khqr/constants"
	"khqr/models"
)

type pointOfInitMtd struct {
	value *string
}

func NewPointOfInitMtd(v *string) KHQRBuilder {

	return &pointOfInitMtd{
		value: v,
	}
}

func (p *pointOfInitMtd) String() string {

	return models.NewTagLengthValue(pointInitMtdCD.Tag, p.value).ToString()
}

func (p *pointOfInitMtd) Validate() *constants.ErrorCode {

	err := ValidateLength(pointInitMtdCD, p.value)
	if err != nil {
		return err
	}

	if p.value != nil && *p.value != constants.StaticQR && *p.value != constants.DynamicQR {
		return &constants.ErrPointOfInitiationMethodInvalid
	}

	return nil
}
