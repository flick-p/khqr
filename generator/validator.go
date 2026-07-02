package generator

import (
	"khqr/constants"
)

func ValidateLength(khqrCode KHQRCodeDict, value *string) *constants.ErrorCode {
	if value != nil && (len([]rune(*value)) > khqrCode.MaxLength) {
		return &khqrCode.ErrInvalidLength
	}

	return nil
}
