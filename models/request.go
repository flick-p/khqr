package models

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
