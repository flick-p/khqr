package generator

import "khqr/constants"

type KHQRCodeDict struct {
	Tag              string
	MaxLength        int
	ErrInvalidLength constants.ErrorCode
	ErrRequire       *constants.ErrorCode
}

var (
	billNumberCD          = KHQRCodeDict{Tag: constants.BillNumberTag, MaxLength: constants.MaxBillNumberLength, ErrInvalidLength: constants.ErrBillNumberLengthInvalid}
	storeLabelCD          = KHQRCodeDict{Tag: constants.StoreLabelTag, MaxLength: constants.MaxStoreLabelLength, ErrInvalidLength: constants.ErrStoreLabelLengthInvalid}
	terminalLabelCD       = KHQRCodeDict{Tag: constants.TerminalTag, MaxLength: constants.MaxTerminalLabelLength, ErrInvalidLength: constants.ErrTerminalLabelLengthInvalid}
	mobileNumberCD        = KHQRCodeDict{Tag: constants.AdditionalDataFieldMobileNumber, MaxLength: constants.MaxMobileNumberLength, ErrInvalidLength: constants.ErrMobileNumberLengthInvalid}
	purposeOfTxnCD        = KHQRCodeDict{Tag: constants.PurposeOfTransactionTag, MaxLength: constants.MaxPurposeOfTransactionLength, ErrInvalidLength: constants.ErrPurposeOfTransactionLengthInvalid}
	countryCodeCD         = KHQRCodeDict{Tag: constants.CountryCodeTag, MaxLength: constants.MaxCountryCodeLength, ErrInvalidLength: constants.ErrCountryCodeLengthInvalid, ErrRequire: &constants.ErrCountryCodeTagRequired}
	crcCD                 = KHQRCodeDict{Tag: constants.CRCTag, MaxLength: constants.MaxCRCLength, ErrInvalidLength: constants.ErrCRCLengthInvalid, ErrRequire: &constants.ErrCRCTagRequired}
	bakongAccountIDCD     = KHQRCodeDict{Tag: constants.BakongAccountIdentifier, MaxLength: constants.MaxBakongAccountLength, ErrInvalidLength: constants.ErrBakongAccountIDLengthInvalid, ErrRequire: &constants.ErrBakongAccountIDRequired}
	accountInformationCD  = KHQRCodeDict{Tag: constants.IndividualAccountInformation, MaxLength: constants.MaxAccountInformationLength, ErrInvalidLength: constants.ErrAccountInformationLengthInvalid}
	merchantIdCD          = KHQRCodeDict{Tag: constants.MerchantAccountInformationMerchantID, MaxLength: constants.MaxMerchantIDLength, ErrInvalidLength: constants.ErrMerchantIDLengthInvalid, ErrRequire: &constants.ErrMerchantIDRequired}
	acquiringBankCD       = KHQRCodeDict{Tag: constants.MerchantAccountInformationAcquiringBank, MaxLength: constants.MaxAcquiringBankLength, ErrInvalidLength: constants.ErrAcquiringBankLengthInvalid, ErrRequire: &constants.ErrAcquiringBankRequired}
	merchantCatCodeCD     = KHQRCodeDict{Tag: constants.MerchantCategoryCodeTag, MaxLength: constants.MaxMerchantCategoryCodeLength, ErrInvalidLength: constants.ErrInvalidMerchantCategoryCode, ErrRequire: &constants.ErrMerchantCategoryTagRequired}
	merchantCityCD        = KHQRCodeDict{Tag: constants.MerchantCityTag, MaxLength: constants.MaxMerchantCityLength, ErrInvalidLength: constants.ErrMerchantCityLengthInvalid, ErrRequire: &constants.ErrMerchantCityTagRequired}
	languagePreCD         = KHQRCodeDict{Tag: constants.LanguagePreference, MaxLength: constants.MaxLanguagePreferenceLength, ErrInvalidLength: constants.ErrLanguagePreferenceLengthInvalid}
	merchantCityAltLangCD = KHQRCodeDict{Tag: constants.MerchantCityAlternateLanguage, MaxLength: constants.MaxMerchantCityAlternateLanguageLength, ErrInvalidLength: constants.ErrMerchantCityAlternateLanguageLengthInvalid}
	merchantNameCD        = KHQRCodeDict{Tag: constants.MerchantNameTag, MaxLength: constants.MaxMerchantNameLength, ErrInvalidLength: constants.ErrMerchantNameLengthInvalid, ErrRequire: &constants.ErrMerchantNameRequired}
	merchantNameAltLangCD = KHQRCodeDict{Tag: constants.MerchantNameAlternateLanguage, MaxLength: constants.MaxMerchantNameAlternateLanguageLength, ErrInvalidLength: constants.ErrMerchantNameAlternateLanguageLengthInvalid}
	payloadFormatIndCD    = KHQRCodeDict{Tag: constants.PayloadFormatIndicator, MaxLength: constants.MaxPayloadFormatIndicatorLength, ErrInvalidLength: constants.ErrPayloadFormatIndicatorLengthInvalid, ErrRequire: &constants.ErrPayloadFormatIndicatorTagRequired}
	pointInitMtdCD        = KHQRCodeDict{Tag: constants.PointOfInitiationMethod, MaxLength: constants.MaxPointOfInitMtdLength, ErrInvalidLength: constants.ErrPointInitiationLengthInvalid}
	transactionAmtCD      = KHQRCodeDict{Tag: constants.TransactionAmountTag, MaxLength: constants.MaxTransactionAmountLength, ErrInvalidLength: constants.ErrTransactionAmountInvalid}
	transactionCurrencyCD = KHQRCodeDict{Tag: constants.TransactionCurrencyTag, MaxLength: constants.MaxTransactionAmountLength, ErrInvalidLength: constants.ErrTransactionCurrencyLengthInvalid}
	unionPayMerchantCD    = KHQRCodeDict{Tag: constants.UnionpayMerchantAccount, MaxLength: constants.MaxUPIMerchantLength, ErrInvalidLength: constants.ErrUPIAccountInformationLengthInvalid}

	// TODO: might need to define max length later
	addAccInfoIndentifier = KHQRCodeDict{constants.AddAccInfoPaymentType, constants.GlobalMaxLength, constants.ErrInvalidLength, nil}
	addAccInfoTxnRef      = KHQRCodeDict{constants.AddAccInfoTxnRef, constants.GlobalMaxLength, constants.ErrInvalidLength, nil}

	addAccInfoMainAccCD      = KHQRCodeDict{constants.AddAccInfoMainAcc, constants.GlobalMaxLength, constants.ErrInvalidLength, nil}
	addAccInfoSecondaryAccCD = KHQRCodeDict{constants.AddAccInfoSecondaryAcc, constants.GlobalMaxLength, constants.ErrInvalidLength, nil}
	addAccInfoTxnTypeCD      = KHQRCodeDict{constants.AddAccInfoTxnType, constants.GlobalMaxLength, constants.ErrInvalidLength, nil}
)
