package util

import (
	"errors"
	"strconv"
)

func CutString(s string) (tag, value, remaining string, err error) {
	r := []rune(s)

	if len(r) < 4 {
		return "", "", "", errors.New("invalid string length")
	}

	tag = string(r[:2])

	lengthStr := string(r[2:4])
	length, err := strconv.Atoi(lengthStr)
	if err != nil {
		return "", "", "", err
	}

	if len(r) < 4+length {
		return "", "", "", errors.New("invalid value length")
	}

	value = string(r[4 : 4+length])
	remaining = string(r[4+length:])

	return tag, value, remaining, nil
}
