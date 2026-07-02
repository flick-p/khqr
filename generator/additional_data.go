package generator

import (
	"khqr/constants"
	"khqr/models"
)

type additionalData struct {
	addDataBuilder []KHQRBuilder
}

func NewAdditionalData(billNum, mobileNum, storeLabel, terminalLabel, pot *string) KHQRBuilder {

	return &additionalData{
		addDataBuilder: []KHQRBuilder{
			newBaseMerchantCode(billNumberCD, billNum, false),
			newBaseMerchantCode(mobileNumberCD, mobileNum, false),
			newBaseMerchantCode(storeLabelCD, storeLabel, false),
			newBaseMerchantCode(terminalLabelCD, terminalLabel, false),
			newBaseMerchantCode(purposeOfTxnCD, pot, false),
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

func (ad *additionalData) Validate() error {

	return BatchValidate(ad.addDataBuilder)
}
