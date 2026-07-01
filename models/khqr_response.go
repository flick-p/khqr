package models

import (
	"fmt"
	"unicode/utf8"
)

// type KHQRResponse struct {
// 	Status struct {
// 		Code      int     `json:"code"`
// 		ErrorCode *int    `json:"errorCode"`
// 		Message   *string `json:"message"`
// 	} `json:"status"`
// 	Data interface{} `json:"data"`
// }

type IndividualInfo struct {
	BakongAccountID               string  `json:"bakongAccountID"`
	AccountInformation            *string `json:"accountInformation,omitempty"`
	AcquiringBank                 *string `json:"acquiringBank,omitempty"`
	Currency                      int     `json:"currency"`
	Amount                        *string `json:"amount,omitempty"`
	MerchantName                  string  `json:"merchantName"`
	MerchantCity                  string  `json:"merchantCity"`
	BillNumber                    *string `json:"billNumber,omitempty"`
	StoreLabel                    *string `json:"storeLabel,omitempty"`
	TerminalLabel                 *string `json:"terminalLabel,omitempty"`
	MobileNumber                  *string `json:"mobileNumber,omitempty"`
	PurposeOfTransaction          *string `json:"purposeOfTransaction,omitempty"`
	LanguagePreference            *string `json:"languagePreference,omitempty"`
	MerchantNameAlternateLanguage *string `json:"merchantNameAlternateLanguage,omitempty"`
	MerchantCityAlternateLanguage *string `json:"merchantCityAlternateLanguage,omitempty"`
	UPIMerchantAccount            *string `json:"upiMerchantAccount,omitempty"`
	ExpirationTimestamp           *int64  `json:"expirationTimestamp,omitempty"`
	MerchantCategoryCode          *string `json:"merchantCategoryCode,omitempty"`
	AddAccInfoIdentifier          *string `json:"accInfoIdentifier,omitempty"`
	AddAccInfoPaymentRef          *string `json:"accInfoPaymentRef,omitempty"`
	AddAccInfoMainAcc             *string `json:"accInfoMainAcc,omitempty"`
	AddAccInfoSecondaryAcc        *string `json:"accInfoSecondaryAcc,omitempty"`
	AddAccInfoTxnType             *string `json:"accInfoTxnType,omitempty"`
}

type MerchantInfo struct {
	IndividualInfo
	MerchantID string `json:"merchantID"`
}

type KHQRData struct {
	QR  string `json:"qr"`
	MD5 string `json:"md5"`
}

type CRCValidation struct {
	IsValid bool `json:"isValid"`
}

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

// func CreateResponse(data interface{}, errorCode *constants.ErrorCode) KHQRResponse {
// 	response := KHQRResponse{
// 		Data: data,
// 	}

// 	if errorCode != nil {
// 		response.Status.Code = 1
// 		response.Status.ErrorCode = &errorCode.Code
// 		response.Status.Message = &errorCode.Message
// 	} else {
// 		response.Status.Code = 0
// 	}

// 	return response
// }
