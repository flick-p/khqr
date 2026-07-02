package models

import (
	"fmt"
	"unicode/utf8"
)

type tagLengthValue struct {
	tag   string
	value *string
}

func NewTagLengthValue(tag string, value *string) *tagLengthValue {
	return &tagLengthValue{
		tag:   tag,
		value: value,
	}
}

func (tlv *tagLengthValue) ToString() string {

	if tlv.value == nil || *tlv.value == "" {
		return ""
	}

	length := utf8.RuneCountInString(*tlv.value)
	lengthStr := fmt.Sprintf("%02d", length)
	return tlv.tag + lengthStr + *tlv.value
}
