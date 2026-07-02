package generator

import (
	"strings"
)

type KHQRBuilder interface {
	String() string
	Validate() error
}

func BatchStringify(v []KHQRBuilder) string {

	var out strings.Builder
	for i := range v {
		out.WriteString(v[i].String())
	}

	return out.String()
}

func BatchValidate(v []KHQRBuilder) error {

	for i := range v {
		err := v[i].Validate()
		if err != nil {

			return err
		}
	}

	return nil
}
