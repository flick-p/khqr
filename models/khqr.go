package models

type DecodedKHQR struct {
	MerchantType                  string  `json:"merchantType"`
	BakongAccountID               string  `json:"bakongAccountID"`
	AccountInformation            *string `json:"accountInformation"`
	MerchantID                    *string `json:"merchantID"`
	AcquiringBank                 *string `json:"acquiringBank"`
	PayloadFormatIndicator        string  `json:"payloadFormatIndicator"`
	PointOfInitiationMethod       *string `json:"pointofInitiationMethod"`
	UnionPayMerchant              *string `json:"unionPayMerchant"`
	MerchantCategoryCode          string  `json:"merchantCategoryCode"`
	TransactionCurrency           string  `json:"transactionCurrency"`
	TransactionAmount             *string `json:"transactionAmount"`
	CountryCode                   string  `json:"countryCode"`
	MerchantName                  string  `json:"merchantName"`
	MerchantCity                  string  `json:"merchantCity"`
	BillNumber                    *string `json:"billNumber"`
	MobileNumber                  *string `json:"mobileNumber"`
	StoreLabel                    *string `json:"storeLabel"`
	TerminalLabel                 *string `json:"terminalLabel"`
	PurposeOfTransaction          *string `json:"purposeOfTransaction"`
	AddAccInfoIdentifier          *string `json:"accInfoIdentifier,omitempty"`
	AddAccInfoPaymentRef          *string `json:"accInfoPaymentRef,omitempty"`
	AddAccInfoMainAcc             *string `json:"accInfoMainAcc,omitempty"`
	AddAccInfoSecondaryAcc        *string `json:"accInfoSecondaryAcc,omitempty"`
	AddAccInfoTxnType             *string `json:"accInfoTxnType,omitempty"`
	LanguagePreference            *string `json:"languagePreference"`
	MerchantNameAlternateLanguage *string `json:"merchantNameAlternateLanguage"`
	MerchantCityAlternateLanguage *string `json:"merchantCityAlternateLanguage"`
	CreationTimestamp             *int64  `json:"creationTimestamp"`
	ExpirationTimestamp           *int64  `json:"expirationTimestamp"`
	CRC                           string  `json:"crc"`
}
