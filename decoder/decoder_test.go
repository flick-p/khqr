package decoder

import (
	"fmt"
	"testing"

	"khqr/constants"
)

func tlv(tag, value string) string {
	return fmt.Sprintf("%s%02d%s", tag, len([]rune(value)), value)
}

func TestDecode_IndividualStaticQR(t *testing.T) {
	raw := "00020101021129320016receivekhqr@dvpy0108123123125204599953038405802KH5904test6010Phnom Penh63042356"

	decoded := NewDecoder(raw).Decode()

	if decoded.MerchantType != constants.MerchantAccountInformationIndividual {
		t.Errorf("MerchantType = %q, want %q", decoded.MerchantType, constants.MerchantAccountInformationIndividual)
	}
	if decoded.BakongAccountID != "receivekhqr@dvpy" {
		t.Errorf("BakongAccountID = %q, want %q", decoded.BakongAccountID, "receivekhqr@dvpy")
	}
	if decoded.AccountInformation == nil || *decoded.AccountInformation != "12312312" {
		t.Errorf("AccountInformation = %v, want %q", decoded.AccountInformation, "12312312")
	}
	if decoded.MerchantName != "test" {
		t.Errorf("MerchantName = %q, want %q", decoded.MerchantName, "test")
	}
	if decoded.MerchantCity != "Phnom Penh" {
		t.Errorf("MerchantCity = %q, want %q", decoded.MerchantCity, "Phnom Penh")
	}
	if decoded.CountryCode != "KH" {
		t.Errorf("CountryCode = %q, want %q", decoded.CountryCode, "KH")
	}
	if decoded.CRC != "2356" {
		t.Errorf("CRC = %q, want %q", decoded.CRC, "2356")
	}
}

func TestDecode_MerchantTag(t *testing.T) {
	gui := tlv(constants.BakongAccountIdentifier, "acc@bank") +
		tlv(constants.MerchantAccountInformationMerchantID, "merchant123")

	raw := tlv(constants.PayloadFormatIndicator, "01") +
		tlv(constants.MerchantAccountInformationMerchant, gui) +
		tlv(constants.MerchantCategoryCodeTag, "5999") +
		tlv(constants.TransactionCurrencyTag, "840") +
		tlv(constants.CountryCodeTag, "KH") +
		tlv(constants.MerchantNameTag, "merchant") +
		tlv(constants.MerchantCityTag, "Phnom Penh")

	decoded := NewDecoder(raw).Decode()

	if decoded.MerchantType != constants.MerchantAccountInformationMerchant {
		t.Errorf("MerchantType = %q, want %q", decoded.MerchantType, constants.MerchantAccountInformationMerchant)
	}
	if decoded.BakongAccountID != "acc@bank" {
		t.Errorf("BakongAccountID = %q, want %q", decoded.BakongAccountID, "acc@bank")
	}
	if decoded.MerchantID == nil || *decoded.MerchantID != "merchant123" {
		t.Errorf("MerchantID = %v, want %q", decoded.MerchantID, "merchant123")
	}
	if decoded.AccountInformation != nil {
		t.Errorf("AccountInformation = %v, want nil for merchant QR", decoded.AccountInformation)
	}
}

func TestDecode_AdditionalData(t *testing.T) {
	additional := tlv(constants.BillNumberTag, "BILL01") +
		tlv(constants.AdditionalDataFieldMobileNumber, "85512345678") +
		tlv(constants.StoreLabelTag, "store1") +
		tlv(constants.TerminalTag, "term1") +
		tlv(constants.PurposeOfTransactionTag, "shopping")

	raw := tlv(constants.AdditionalDataTag, additional)

	decoded := NewDecoder(raw).Decode()

	if decoded.BillNumber == nil || *decoded.BillNumber != "BILL01" {
		t.Errorf("BillNumber = %v, want %q", decoded.BillNumber, "BILL01")
	}
	if decoded.MobileNumber == nil || *decoded.MobileNumber != "85512345678" {
		t.Errorf("MobileNumber = %v, want %q", decoded.MobileNumber, "85512345678")
	}
	if decoded.StoreLabel == nil || *decoded.StoreLabel != "store1" {
		t.Errorf("StoreLabel = %v, want %q", decoded.StoreLabel, "store1")
	}
	if decoded.TerminalLabel == nil || *decoded.TerminalLabel != "term1" {
		t.Errorf("TerminalLabel = %v, want %q", decoded.TerminalLabel, "term1")
	}
	if decoded.PurposeOfTransaction == nil || *decoded.PurposeOfTransaction != "shopping" {
		t.Errorf("PurposeOfTransaction = %v, want %q", decoded.PurposeOfTransaction, "shopping")
	}
}

func TestDecode_LanguageTemplate(t *testing.T) {
	lang := tlv(constants.LanguagePreference, "km") +
		tlv(constants.MerchantNameAlternateLanguage, "ឈ្មោះ") +
		tlv(constants.MerchantCityAlternateLanguage, "ភ្នំពេញ")

	raw := tlv(constants.MerchantInformationLanguageTemplate, lang)

	decoded := NewDecoder(raw).Decode()

	if decoded.LanguagePreference == nil || *decoded.LanguagePreference != "km" {
		t.Errorf("LanguagePreference = %v, want %q", decoded.LanguagePreference, "km")
	}
	if decoded.MerchantNameAlternateLanguage == nil || *decoded.MerchantNameAlternateLanguage != "ឈ្មោះ" {
		t.Errorf("MerchantNameAlternateLanguage = %v, want %q", decoded.MerchantNameAlternateLanguage, "ឈ្មោះ")
	}
	if decoded.MerchantCityAlternateLanguage == nil || *decoded.MerchantCityAlternateLanguage != "ភ្នំពេញ" {
		t.Errorf("MerchantCityAlternateLanguage = %v, want %q", decoded.MerchantCityAlternateLanguage, "ភ្នំពេញ")
	}
}

func TestDecode_Timestamp(t *testing.T) {
	ts := tlv(constants.CreationTimestamp, "1700000000000") +
		tlv(constants.ExpirationTimestamp, "1800000000000")

	raw := tlv(constants.TimestampTag, ts)

	decoded := NewDecoder(raw).Decode()

	if decoded.CreationTimestamp == nil || *decoded.CreationTimestamp != 1700000000000 {
		t.Errorf("CreationTimestamp = %v, want %d", decoded.CreationTimestamp, 1700000000000)
	}
	if decoded.ExpirationTimestamp == nil || *decoded.ExpirationTimestamp != 1800000000000 {
		t.Errorf("ExpirationTimestamp = %v, want %d", decoded.ExpirationTimestamp, 1800000000000)
	}
}

func TestDecode_EmptyString(t *testing.T) {
	decoded := NewDecoder("").Decode()

	if decoded.MerchantType != constants.MerchantAccountInformationIndividual {
		t.Errorf("MerchantType = %q, want default individual type", decoded.MerchantType)
	}
	if decoded.BakongAccountID != "" {
		t.Errorf("BakongAccountID = %q, want empty", decoded.BakongAccountID)
	}
}

func TestDecode_MalformedStringDoesNotPanic(t *testing.T) {
	malformed := []string{
		"0",
		"00",
		"0001",
		"000199",
		"00AA",
	}

	for _, raw := range malformed {
		t.Run(raw, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Fatalf("Decode() panicked on input %q: %v", raw, r)
				}
			}()
			NewDecoder(raw).Decode()
		})
	}
}

func TestDecode_DuplicateTagStopsParsing(t *testing.T) {
	raw := tlv(constants.PayloadFormatIndicator, "01") + tlv(constants.PayloadFormatIndicator, "02") + tlv(constants.MerchantNameTag, "test")

	decoded := NewDecoder(raw).Decode()

	if decoded.MerchantName != "" {
		t.Errorf("MerchantName = %q, want empty because parsing should stop at the duplicate tag", decoded.MerchantName)
	}
}
