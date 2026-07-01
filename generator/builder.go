package generator

import (
	"strings"

	"khqr/constants"
)

type KHQRBuilder interface {
	String() string
	Validate() *constants.ErrorCode
}

func BatchStringify(v []KHQRBuilder) string {

	var out strings.Builder
	for i := range v {
		out.WriteString(v[i].String())
	}

	return out.String()
}

func BatchValidate(v []KHQRBuilder) *constants.ErrorCode {

	for i := range v {
		err := v[i].Validate()
		if err != nil {

			return err
		}
	}

	return nil
}
