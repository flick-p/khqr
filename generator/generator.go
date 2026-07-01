package generator

import (
	"strings"

	"khqr/constants"
	"khqr/models"
	"khqr/util"
)

type khqrGenerator struct {
	data       models.MerchantInfo
	isMerchant bool
}

func NewKHQRGenerator(data models.MerchantInfo, isMerchant bool) *khqrGenerator {
	return &khqrGenerator{
		data:       data,
		isMerchant: isMerchant,
	}
}

func (g *khqrGenerator) Generate() (string, *constants.ErrorCode) {
	var (
		data            = g.data
		qrType          = constants.DynamicQR
		merchantCatCode = constants.DefaultMerchantCategoryCode
		out             strings.Builder
		merchantCity    = constants.DefaultMerchantCity
	)

	// Determine QR type
	if data.Amount == nil || *data.Amount == "" || *data.Amount == "0" {
		qrType = constants.StaticQR
	}

	// Fallback to default Merchant Category Code if not provided
	if data.MerchantCategoryCode != nil {

		merchantCatCode = *data.MerchantCategoryCode
	}

	if data.MerchantCity != "" {
		merchantCity = data.MerchantCity
	}

	// Build KHQR components
	builder := []KHQRBuilder{
		NewPayloadFormatInd(util.Ptr(constants.DefaultPayloadFormatIndicator)),
		NewPointOfInitMtd(&qrType),
		NewUnionPayMerchant(data.UPIMerchantAccount),
		NewGlobalUniqueIdentifier(InitGlobalIdentifier{
			BakongAccountID:    &data.BakongAccountID,
			MerchantID:         &data.MerchantID,
			AcquiringBank:      data.AcquiringBank,
			AccountInformation: data.AccountInformation,
			IsMerchant:         g.isMerchant,
		}),
		NewMerchantCategoryCode(&merchantCatCode),
		NewTransactionCurrency(data.Currency),
		NewTransactionAmount(data.Amount, data.Currency),
		NewCountryCode(util.Ptr(constants.DefaultCountryCode)),
		NewMerchantName(&data.MerchantName),
		NewMerchantCity(&merchantCity),
		NewAdditionalData(
			data.BillNumber,
			data.MobileNumber,
			data.StoreLabel,
			data.TerminalLabel,
			data.PurposeOfTransaction,
		),
		NewAdditionalAccInfo(
			data.AddAccInfoIdentifier,
			data.AddAccInfoPaymentRef,
			data.AddAccInfoMainAcc,
			data.AddAccInfoSecondaryAcc,
			data.AddAccInfoTxnType,
		),
		NewMerchantInfoLangTemplate(
			data.LanguagePreference,
			data.MerchantNameAlternateLanguage,
			data.MerchantCityAlternateLanguage,
		),
		NewTimestamp(data.ExpirationTimestamp),
	}

	// Validate KHQR fields
	if err := BatchValidate(builder); err != nil {
		return "", err
	}

	// Generate KHQR string (without CRC)
	out.WriteString(BatchStringify(builder))
	out.WriteString(constants.CRCTag)
	out.WriteString(constants.CRCLength)

	// Calculate and validate CRC
	crc := util.CalculateCRC16(out.String())
	if err := NewCrc(&crc).Validate(); err != nil {
		return "", err
	}

	// Append CRC and return
	out.WriteString(crc)
	return out.String(), nil
}
