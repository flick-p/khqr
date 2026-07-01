package generator

import (
	"khqr/constants"
	"khqr/models"
)

// struct
type additionalData struct {
	BillNumberInput      *string
	MobileNumberInput    *string
	StoreLabelInput      *string
	TerminalLabelInput   *string
	PurposeOfTransaction *string
	addDataBuilder       []KHQRBuilder
}

func NewAdditionalData(billNum, mobileNum, storeLabel, terminalLabel, pot *string) KHQRBuilder {

	return &additionalData{
		BillNumberInput:      billNum,
		MobileNumberInput:    mobileNum,
		StoreLabelInput:      storeLabel,
		TerminalLabelInput:   terminalLabel,
		PurposeOfTransaction: pot,
		addDataBuilder: []KHQRBuilder{
			newBaseMerchantCode(billNumberCD, billNum),
			newBaseMerchantCode(mobileNumberCD, mobileNum),
			newBaseMerchantCode(storeLabelCD, storeLabel),
			newBaseMerchantCode(terminalLabelCD, terminalLabel),
			newBaseMerchantCode(purposeOfTxnCD, pot),
		},
	}
}

func (ad *additionalData) String() string {

	sub := BatchStringify(ad.addDataBuilder)

	if sub == "" {
		return ""
	}

	return models.NewTagLengthValue(constants.AdditionalDataTag, &sub).ToString()
}

func (ad *additionalData) Validate() *constants.ErrorCode {

	return BatchValidate(ad.addDataBuilder)
}
