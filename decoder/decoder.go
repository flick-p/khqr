package decoder

import (
	"strconv"

	"khqr/constants"
	"khqr/models"
	"khqr/util"
)

type decoder struct {
	khqr    string
	decoded models.DecodedKHQR
}

func NewDecoder(khqr string) *decoder {
	return &decoder{
		khqr: khqr,
	}
}

func (d *decoder) Decode() models.DecodedKHQR {
	remaining := d.khqr

	merchantType := constants.MerchantAccountInformationIndividual
	var lastTag string

	for len(remaining) > 0 {
		tag, value, newRemaining, err := util.CutString(remaining)
		if err != nil {
			break
		}

		if tag == lastTag {
			break
		}

		// Handle merchant vs individual
		isMerchant := tag == constants.MerchantAccountInformationMerchant
		if isMerchant {
			merchantType = constants.MerchantAccountInformationMerchant
			tag = "29" // Normalize to 29 for processing
		}

		// Process each tag
		switch tag {
		case constants.PayloadFormatIndicator:
			d.decoded.PayloadFormatIndicator = value
		case constants.PointOfInitiationMethod:
			d.decoded.PointOfInitiationMethod = &value
		case constants.UnionpayMerchantAccount:
			d.decoded.UnionPayMerchant = &value
		case "29": // Global unique identifier
			d.decodeGlobalUniqueIdentifier(value, isMerchant)

		case constants.MerchantCategoryCodeTag:
			d.decoded.MerchantCategoryCode = value
		case constants.TransactionCurrencyTag:
			d.decoded.TransactionCurrency = value
		case constants.TransactionAmountTag:
			d.decoded.TransactionAmount = &value
		case constants.CountryCodeTag:
			d.decoded.CountryCode = value
		case constants.MerchantNameTag:
			d.decoded.MerchantName = value
		case constants.MerchantCityTag:
			d.decoded.MerchantCity = value
		case constants.AdditionalDataTag:
			d.decodeAdditionalData(value)

		case constants.AdditionalAccountInfoTag:

			d.decodeAddtionalAccInfo(value)

		case constants.MerchantInformationLanguageTemplate:
			d.decodeLanguageTemplate(value)

		case constants.TimestampTag:
			d.decodeTimestamp(value)

		case constants.CRCTag:
			d.decoded.CRC = value
		}

		remaining = newRemaining
		lastTag = tag
	}

	d.decoded.MerchantType = merchantType

	return d.decoded
}

func (d *decoder) decodeGlobalUniqueIdentifier(value string, isMerchant bool) {
	remaining := value

	for len(remaining) > 0 {
		tag, val, newRemaining, err := util.CutString(remaining)
		if err != nil {
			break
		}

		switch tag {
		case constants.BakongAccountIdentifier:
			d.decoded.BakongAccountID = val
		case constants.MerchantAccountInformationMerchantID:
			if isMerchant {
				d.decoded.MerchantID = &val
			} else {
				d.decoded.AccountInformation = &val
			}
		case constants.MerchantAccountInformationAcquiringBank:
			d.decoded.AcquiringBank = &val
		}

		remaining = newRemaining
	}
}

func (d *decoder) decodeAdditionalData(value string) {
	remaining := value

	for len(remaining) > 0 {
		tag, val, newRemaining, err := util.CutString(remaining)
		if err != nil {
			break
		}

		switch tag {
		case constants.BillNumberTag:
			d.decoded.BillNumber = &val
		case constants.AdditionalDataFieldMobileNumber:
			d.decoded.MobileNumber = &val
		case constants.StoreLabelTag:
			d.decoded.StoreLabel = &val
		case constants.TerminalTag:
			d.decoded.TerminalLabel = &val
		case constants.PurposeOfTransactionTag:
			d.decoded.PurposeOfTransaction = &val
		}

		remaining = newRemaining
	}

}

func (d *decoder) decodeAddtionalAccInfo(value string) {

	remaining := value

	for len(remaining) > 0 {
		tag, val, newRemaining, err := util.CutString(remaining)
		if err != nil {
			break
		}

		switch tag {
		case constants.AddAccInfoMainAcc:
			d.decoded.AddAccInfoMainAcc = &val
		case constants.AddAccInfoSecondaryAcc:
			d.decoded.AddAccInfoSecondaryAcc = &val
		case constants.AddAccInfoTxnType:
			d.decoded.AddAccInfoTxnType = &val
		}

		remaining = newRemaining
	}
}

func (d *decoder) decodeLanguageTemplate(value string) {
	remaining := value

	for len(remaining) > 0 {
		tag, val, newRemaining, err := util.CutString(remaining)
		if err != nil {
			break
		}

		switch tag {
		case constants.LanguagePreference:
			d.decoded.LanguagePreference = &val
		case constants.MerchantNameAlternateLanguage:
			d.decoded.MerchantNameAlternateLanguage = &val
		case constants.MerchantCityAlternateLanguage:
			d.decoded.MerchantCityAlternateLanguage = &val
		}

		remaining = newRemaining
	}

}

func (d *decoder) decodeTimestamp(value string) error {
	remaining := value

	for len(remaining) > 0 {
		tag, val, newRemaining, err := util.CutString(remaining)
		if err != nil {
			break
		}

		switch tag {
		case constants.CreationTimestamp:
			if timestamp, err := strconv.ParseInt(val, 10, 64); err == nil {
				d.decoded.CreationTimestamp = &timestamp
			}
		case constants.ExpirationTimestamp:
			if timestamp, err := strconv.ParseInt(val, 10, 64); err == nil {
				d.decoded.ExpirationTimestamp = &timestamp
			}
		}

		remaining = newRemaining
	}

	return nil
}
