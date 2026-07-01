package constants

type ErrorCode struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *ErrorCode) Error() string {
	return e.Message
}

var (
	ErrBakongAccountIDRequired                    = ErrorCode{1, "Bakong Account ID cannot be null or empty"}
	ErrMerchantNameRequired                       = ErrorCode{2, "Merchant name cannot be null or empty"}
	ErrBakongAccountIDInvalid                     = ErrorCode{3, "Bakong Account ID is invalid"}
	ErrTransactionAmountInvalid                   = ErrorCode{4, "Amount is invalid"}
	ErrMerchantTypeRequired                       = ErrorCode{5, "Merchant type cannot be null or empty"}
	ErrBakongAccountIDLengthInvalid               = ErrorCode{6, "Bakong Account ID Length is Invalid"}
	ErrMerchantNameLengthInvalid                  = ErrorCode{7, "Merchant Name Length is invalid"}
	ErrKHQRInvalid                                = ErrorCode{8, "KHQR provided is invalid"}
	ErrCurrencyTypeRequired                       = ErrorCode{9, "Currency type cannot be null or empty"}
	ErrBillNumberLengthInvalid                    = ErrorCode{10, "Bill Name Length is invalid"}
	ErrStoreLabelLengthInvalid                    = ErrorCode{11, "Store Label Length is invalid"}
	ErrTerminalLabelLengthInvalid                 = ErrorCode{12, "Terminal Label Length is invalid"}
	ErrConnectionTimeout                          = ErrorCode{13, "Cannot reach Bakong Open API service. Please check internet connection"}
	ErrInvalidDeepLinkSourceInfo                  = ErrorCode{14, "Source Info for Deep Link is invalid"}
	ErrInternalServer                             = ErrorCode{15, "Internal server error"}
	ErrPayloadFormatIndicatorLengthInvalid        = ErrorCode{16, "Payload Format indicator Length is invalid"}
	ErrPointInitiationLengthInvalid               = ErrorCode{17, "Point of initiation Length is invalid"}
	ErrMerchantCodeLengthInvalid                  = ErrorCode{18, "Merchant code Length is invalid"}
	ErrTransactionCurrencyLengthInvalid           = ErrorCode{19, "Transaction currency Length is invalid"}
	ErrCountryCodeLengthInvalid                   = ErrorCode{20, "Country code Length is invalid"}
	ErrMerchantCityLengthInvalid                  = ErrorCode{21, "Merchant city Length is invalid"}
	ErrCRCLengthInvalid                           = ErrorCode{22, "CRC Length is invalid"}
	ErrPayloadFormatIndicatorTagRequired          = ErrorCode{23, "Payload format indicator tag required"}
	ErrCRCTagRequired                             = ErrorCode{24, "CRC tag required"}
	ErrMerchantCategoryTagRequired                = ErrorCode{25, "Merchant category tag required"}
	ErrCountryCodeTagRequired                     = ErrorCode{26, "Country Code cannot be null or empty"}
	ErrMerchantCityTagRequired                    = ErrorCode{27, "Merchant City cannot be null or empty"}
	ErrUnsupportedCurrency                        = ErrorCode{28, "Unsupported currency"}
	ErrInvalidDeepLinkURL                         = ErrorCode{29, "Deep Link URL is not valid"}
	ErrMerchantIDRequired                         = ErrorCode{30, "Merchant ID cannot be null or empty"}
	ErrAcquiringBankRequired                      = ErrorCode{31, "Acquiring Bank cannot be null or empty"}
	ErrMerchantIDLengthInvalid                    = ErrorCode{32, "Merchant ID Length is invalid"}
	ErrAcquiringBankLengthInvalid                 = ErrorCode{33, "Acquiring Bank Length is invalid"}
	ErrMobileNumberLengthInvalid                  = ErrorCode{34, "Mobile Number Length is invalid"}
	ErrAccountInformationLengthInvalid            = ErrorCode{35, "Account Information Length is invalid"}
	ErrTagNotInOrder                              = ErrorCode{36, "Tag is not in order"}
	ErrLanguagePreferenceRequired                 = ErrorCode{37, "Language Preference cannot be null or empty"}
	ErrLanguagePreferenceLengthInvalid            = ErrorCode{38, "Language Preference Length is invalid"}
	ErrMerchantNameAlternateLanguageRequired      = ErrorCode{39, "Merchant Name Alternate Language cannot be null or empty"}
	ErrMerchantNameAlternateLanguageLengthInvalid = ErrorCode{40, "Merchant Name Alternate Language Length is invalid"}
	ErrMerchantCityAlternateLanguageLengthInvalid = ErrorCode{41, "Merchant City Alternate Language Length is invalid"}
	ErrPurposeOfTransactionLengthInvalid          = ErrorCode{42, "Purpose of Transaction Length is invalid"}
	ErrUPIAccountInformationLengthInvalid         = ErrorCode{43, "Upi Account Information Length is invalid"}
	ErrUPIAccountInformationInvalidCurrency       = ErrorCode{44, "Upi Account Information Length does not accept USD"}
	ErrExpirationTimestampRequired                = ErrorCode{45, "Expiration timestamp is required for dynamic KHQR"}
	ErrKHQRExpired                                = ErrorCode{46, "This dynamic KHQR has expired"}
	ErrInvalidDynamicKHQR                         = ErrorCode{47, "This dynamic KHQR has invalid field transaction amount"}
	ErrPointOfInitiationMethodInvalid             = ErrorCode{48, "Point of Initiation Method is invalid"}
	ErrExpirationTimestampLengthInvalid           = ErrorCode{49, "Expiration timestamp length is invalid"}
	ErrExpirationTimestampInThePast               = ErrorCode{50, "Expiration timestamp is in the past"}
	ErrInvalidMerchantCategoryCode                = ErrorCode{51, "Invalid merchant category code"}

	ErrInvalidLength = ErrorCode{-1, "Invalid length"}
)
