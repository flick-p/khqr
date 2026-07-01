package constants

// Constants
const (

	// EMV Constants
	PayloadFormatIndicator                  = "00"
	DefaultPayloadFormatIndicator           = "01"
	PointOfInitiationMethod                 = "01"
	StaticQR                                = "11"
	DynamicQR                               = "12"
	MerchantAccountInformationIndividual    = "29"
	MerchantAccountInformationMerchant      = "30"
	BakongAccountIdentifier                 = "00"
	MerchantAccountInformationMerchantID    = "01"
	IndividualAccountInformation            = "01"
	MerchantAccountInformationAcquiringBank = "02"
	MerchantCategoryCodeTag                 = "52"
	DefaultMerchantCategoryCode             = "5999"
	TransactionCurrencyTag                  = "53"
	TransactionAmountTag                    = "54"
	DefaultTransactionAmount                = "0"
	CountryCodeTag                          = "58"
	DefaultCountryCode                      = "KH"
	MerchantNameTag                         = "59"
	MerchantCityTag                         = "60"
	DefaultMerchantCity                     = "Phnom Penh"
	CRCTag                                  = "63"
	CRCLength                               = "04"
	AdditionalDataTag                       = "62"
	BillNumberTag                           = "01"
	AdditionalDataFieldMobileNumber         = "02"
	StoreLabelTag                           = "03"
	TerminalTag                             = "07"
	PurposeOfTransactionTag                 = "08"
	TimestampTag                            = "99"
	CreationTimestamp                       = "00"
	ExpirationTimestamp                     = "01"
	MerchantInformationLanguageTemplate     = "64"
	LanguagePreference                      = "00"
	MerchantNameAlternateLanguage           = "01"
	MerchantCityAlternateLanguage           = "02"
	UnionpayMerchantAccount                 = "15"
	AdditionalAccountInfoTag                = "40"
	AddAccInfoPaymentType                   = "00"
	AddAccInfoTxnRef                        = "01"
	AddAccInfoMainAcc                       = "02"
	AddAccInfoSecondaryAcc                  = "03"
	AddAccInfoTxnType                       = "04"

	// Currency codes
	CurrencyUSD = 840
	CurrencyKHR = 116

	// Max lengths
	MaxKHQRLength                          = 12
	MaxMerchantNameLength                  = 25
	MaxBakongAccountLength                 = 32
	MaxAmountLength                        = 13
	MaxCountryCodeLength                   = 3
	MaxMerchantCategoryCodeLength          = 4
	MaxMerchantCityLength                  = 15
	TimestampLength                        = 13
	MaxTransactionAmountLength             = 14
	MaxTransactionCurrencyLength           = 3
	MaxBillNumberLength                    = 25
	MaxStoreLabelLength                    = 25
	MaxTerminalLabelLength                 = 25
	MaxPurposeOfTransactionLength          = 25
	MaxMerchantIDLength                    = 32
	MaxAcquiringBankLength                 = 32
	MaxMobileNumberLength                  = 25
	MaxAccountInformationLength            = 32
	MaxMerchantInformationLanguageTemplate = 99
	MaxUPIMerchantLength                   = 99
	MaxLanguagePreferenceLength            = 2
	MaxMerchantNameAlternateLanguageLength = 25
	MaxMerchantCityAlternateLanguageLength = 15
	MaxCRCLength                           = 4
	MaxPayloadFormatIndicatorLength        = 2
	MaxPointOfInitMtdLength                = 2

	GlobalMaxLength = 99

	// Txn Type
	TxnTypeDual = "Dual"
)
