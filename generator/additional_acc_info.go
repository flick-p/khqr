package generator

import (
	"khqr/constants"
	"khqr/models"
)

type additionalAccountInfo struct {
	addAccInfoBuilder []KHQRBuilder
}

func NewAdditionalAccInfo(identifier, paymentRef, mainAcc, secondaryAcc, txnType *string) KHQRBuilder {

	return &additionalAccountInfo{
		addAccInfoBuilder: []KHQRBuilder{
			newBaseMerchantCode(addAccInfoIndentifier, identifier, false),
			newBaseMerchantCode(addAccInfoTxnRef, paymentRef, false),
			newBaseMerchantCode(addAccInfoMainAccCD, mainAcc, false),
			newBaseMerchantCode(addAccInfoSecondaryAccCD, secondaryAcc, false),
			newBaseMerchantCode(addAccInfoTxnTypeCD, txnType, false),
		},
	}
}

func (ad *additionalAccountInfo) String() string {

	sub := BatchStringify(ad.addAccInfoBuilder)

	if sub == "" {
		return ""
	}

	return models.NewTagLengthValue(constants.AdditionalAccountInfoTag, &sub).ToString()
}

func (ad *additionalAccountInfo) Validate() error {

	return BatchValidate(ad.addAccInfoBuilder)
}
