package generator

import (
	"khqr/constants"
	"khqr/models"
)

type additionalAccountInfo struct {
	mainAcc, secondaryAcc, txnType *string
	addAccInfoBuilder              []KHQRBuilder
}

func NewAdditionalAccInfo(identifier, paymentRef, mainAcc, secondaryAcc, txnType *string) KHQRBuilder {

	return &additionalAccountInfo{

		mainAcc:      mainAcc,
		secondaryAcc: secondaryAcc,
		txnType:      txnType,
		addAccInfoBuilder: []KHQRBuilder{
			newBaseMerchantCode(addAccInfoIndentifier, identifier),
			newBaseMerchantCode(addAccInfoTxnRef, paymentRef),
			newBaseMerchantCode(addAccInfoMainAccCD, mainAcc),
			newBaseMerchantCode(addAccInfoSecondaryAccCD, secondaryAcc),
			newBaseMerchantCode(addAccInfoTxnTypeCD, txnType),
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

func (ad *additionalAccountInfo) Validate() *constants.ErrorCode {

	return BatchValidate(ad.addAccInfoBuilder)
}
