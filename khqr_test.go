package khqr

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"khqr/constants"
	"khqr/models"
	"khqr/util"
)

// tlv builds a single EMVCo tag-length-value segment for hand-crafting raw
// KHQR strings in tests that need to reach decode-only validation branches
// (e.g. malformed/expired QRs) that GenerateIndividual/GenerateMerchant can
// never produce because they enforce those invariants at generation time.
func tlv(tag, value string) string {
	return fmt.Sprintf("%s%02d%s", tag, len([]rune(value)), value)
}

// withCRC appends a valid CRC tag/value computed over the given payload,
// mirroring what a real KHQR generator would produce.
func withCRC(payload string) string {
	base := payload + constants.CRCTag + constants.CRCLength
	return base + util.CalculateCRC16(base)
}

func TestGenerateIndividual(t *testing.T) {
	t.Run("static QR without amount", func(t *testing.T) {
		res, err := GenerateIndividual(models.IndividualInfo{
			BakongAccountID:    "receivekhqr@dvpy",
			Currency:           constants.CurrencyUSD,
			MerchantName:       "test",
			AccountInformation: util.Ptr("12312312"),
		})
		if err != nil {
			t.Fatalf("GenerateIndividual() error = %v", err)
		}
		if res.QR == "" {
			t.Fatal("QR string is empty")
		}
		if res.MD5 != util.CalculateMD5(res.QR) {
			t.Errorf("MD5 = %q, does not match CalculateMD5(QR)", res.MD5)
		}

		decoded := DecodeKHQR(res.QR)
		if decoded.PointOfInitiationMethod == nil || *decoded.PointOfInitiationMethod != constants.StaticQR {
			t.Errorf("PointOfInitiationMethod = %v, want %q (static)", decoded.PointOfInitiationMethod, constants.StaticQR)
		}
		if decoded.TransactionAmount != nil {
			t.Errorf("TransactionAmount = %v, want nil for static QR", decoded.TransactionAmount)
		}
		if decoded.BakongAccountID != "receivekhqr@dvpy" {
			t.Errorf("BakongAccountID = %q, want %q", decoded.BakongAccountID, "receivekhqr@dvpy")
		}
		if decoded.MerchantCity != constants.DefaultMerchantCity {
			t.Errorf("MerchantCity = %q, want default %q", decoded.MerchantCity, constants.DefaultMerchantCity)
		}

		if _, err := DecodeKHQRValidation(res.QR); err != nil {
			t.Errorf("DecodeKHQRValidation() error = %v, want valid QR", err)
		}
	})

	t.Run("dynamic QR with USD amount", func(t *testing.T) {
		exp := time.Now().Add(time.Hour).UnixMilli()
		res, err := GenerateIndividual(models.IndividualInfo{
			BakongAccountID:     "receivekhqr@dvpy",
			Currency:            constants.CurrencyUSD,
			MerchantName:        "test",
			MerchantCity:        "Siem Reap",
			Amount:              util.Ptr("10.5"),
			ExpirationTimestamp: util.Ptr(exp),
		})
		if err != nil {
			t.Fatalf("GenerateIndividual() error = %v", err)
		}

		decoded := DecodeKHQR(res.QR)
		if decoded.PointOfInitiationMethod == nil || *decoded.PointOfInitiationMethod != constants.DynamicQR {
			t.Errorf("PointOfInitiationMethod = %v, want %q (dynamic)", decoded.PointOfInitiationMethod, constants.DynamicQR)
		}
		if decoded.TransactionAmount == nil || *decoded.TransactionAmount != "10.50" {
			t.Errorf("TransactionAmount = %v, want %q", decoded.TransactionAmount, "10.50")
		}
		if decoded.MerchantCity != "Siem Reap" {
			t.Errorf("MerchantCity = %q, want %q", decoded.MerchantCity, "Siem Reap")
		}
		if decoded.ExpirationTimestamp == nil || *decoded.ExpirationTimestamp != exp {
			t.Errorf("ExpirationTimestamp = %v, want %d", decoded.ExpirationTimestamp, exp)
		}

		if _, err := DecodeKHQRValidation(res.QR); err != nil {
			t.Errorf("DecodeKHQRValidation() error = %v, want valid QR", err)
		}
	})

	t.Run("dynamic QR with KHR integer amount", func(t *testing.T) {
		res, err := GenerateIndividual(models.IndividualInfo{
			BakongAccountID: "receivekhqr@dvpy",
			Currency:        constants.CurrencyKHR,
			MerchantName:    "test",
			Amount:          util.Ptr("5000"),
		})
		if err != nil {
			t.Fatalf("GenerateIndividual() error = %v", err)
		}

		decoded := DecodeKHQR(res.QR)
		if decoded.TransactionAmount == nil || *decoded.TransactionAmount != "5000" {
			t.Errorf("TransactionAmount = %v, want %q", decoded.TransactionAmount, "5000")
		}
		if decoded.TransactionCurrency != "116" {
			t.Errorf("TransactionCurrency = %q, want %q", decoded.TransactionCurrency, "116")
		}
	})

	t.Run("with additional data and language template", func(t *testing.T) {
		res, err := GenerateIndividual(models.IndividualInfo{
			BakongAccountID:               "receivekhqr@dvpy",
			Currency:                      constants.CurrencyUSD,
			MerchantName:                  "test",
			BillNumber:                    util.Ptr("BILL01"),
			StoreLabel:                    util.Ptr("store1"),
			TerminalLabel:                 util.Ptr("term1"),
			MobileNumber:                  util.Ptr("85512345678"),
			PurposeOfTransaction:          util.Ptr("shopping"),
			LanguagePreference:            util.Ptr("km"),
			MerchantNameAlternateLanguage: util.Ptr("merchant-km"),
			MerchantCityAlternateLanguage: util.Ptr("city-km"),
		})
		if err != nil {
			t.Fatalf("GenerateIndividual() error = %v", err)
		}

		decoded := DecodeKHQR(res.QR)
		if decoded.BillNumber == nil || *decoded.BillNumber != "BILL01" {
			t.Errorf("BillNumber = %v, want %q", decoded.BillNumber, "BILL01")
		}
		if decoded.StoreLabel == nil || *decoded.StoreLabel != "store1" {
			t.Errorf("StoreLabel = %v, want %q", decoded.StoreLabel, "store1")
		}
		if decoded.TerminalLabel == nil || *decoded.TerminalLabel != "term1" {
			t.Errorf("TerminalLabel = %v, want %q", decoded.TerminalLabel, "term1")
		}
		if decoded.MobileNumber == nil || *decoded.MobileNumber != "85512345678" {
			t.Errorf("MobileNumber = %v, want %q", decoded.MobileNumber, "85512345678")
		}
		if decoded.PurposeOfTransaction == nil || *decoded.PurposeOfTransaction != "shopping" {
			t.Errorf("PurposeOfTransaction = %v, want %q", decoded.PurposeOfTransaction, "shopping")
		}
		if decoded.LanguagePreference == nil || *decoded.LanguagePreference != "km" {
			t.Errorf("LanguagePreference = %v, want %q", decoded.LanguagePreference, "km")
		}
		if decoded.MerchantNameAlternateLanguage == nil || *decoded.MerchantNameAlternateLanguage != "merchant-km" {
			t.Errorf("MerchantNameAlternateLanguage = %v, want %q", decoded.MerchantNameAlternateLanguage, "merchant-km")
		}
	})

	t.Run("errors", func(t *testing.T) {
		tests := []struct {
			name     string
			data     models.IndividualInfo
			wantCode int
		}{
			{
				name: "missing @ in Bakong account id",
				data: models.IndividualInfo{
					BakongAccountID: "no-at-sign",
					Currency:        constants.CurrencyUSD,
					MerchantName:    "test",
				},
				wantCode: constants.ErrBakongAccountIDInvalid.Code,
			},
			{
				name: "Bakong account id too long",
				data: models.IndividualInfo{
					BakongAccountID: strings.Repeat("a", 40) + "@bank",
					Currency:        constants.CurrencyUSD,
					MerchantName:    "test",
				},
				wantCode: constants.ErrBakongAccountIDLengthInvalid.Code,
			},
			{
				name: "merchant name too long",
				data: models.IndividualInfo{
					BakongAccountID: "receivekhqr@dvpy",
					Currency:        constants.CurrencyUSD,
					MerchantName:    strings.Repeat("a", 26),
				},
				wantCode: constants.ErrMerchantNameLengthInvalid.Code,
			},
			{
				name: "currency required",
				data: models.IndividualInfo{
					BakongAccountID: "receivekhqr@dvpy",
					MerchantName:    "test",
				},
				wantCode: constants.ErrCurrencyTypeRequired.Code,
			},
			{
				name: "unsupported currency",
				data: models.IndividualInfo{
					BakongAccountID: "receivekhqr@dvpy",
					Currency:        999,
					MerchantName:    "test",
				},
				wantCode: constants.ErrUnsupportedCurrency.Code,
			},
			{
				name: "negative amount",
				data: models.IndividualInfo{
					BakongAccountID: "receivekhqr@dvpy",
					Currency:        constants.CurrencyUSD,
					MerchantName:    "test",
					Amount:          util.Ptr("-5"),
				},
				wantCode: constants.ErrTransactionAmountInvalid.Code,
			},
			{
				name: "non-integer KHR amount",
				data: models.IndividualInfo{
					BakongAccountID: "receivekhqr@dvpy",
					Currency:        constants.CurrencyKHR,
					MerchantName:    "test",
					Amount:          util.Ptr("50.5"),
				},
				wantCode: constants.ErrTransactionAmountInvalid.Code,
			},
			{
				name: "invalid merchant category code",
				data: models.IndividualInfo{
					BakongAccountID:      "receivekhqr@dvpy",
					Currency:             constants.CurrencyUSD,
					MerchantName:         "test",
					MerchantCategoryCode: util.Ptr("abcd"),
				},
				wantCode: constants.ErrInvalidMerchantCategoryCode.Code,
			},
			{
				name: "expiration timestamp in the past",
				data: models.IndividualInfo{
					BakongAccountID:     "receivekhqr@dvpy",
					Currency:            constants.CurrencyUSD,
					MerchantName:        "test",
					Amount:              util.Ptr("10"),
					ExpirationTimestamp: util.Ptr(time.Now().Add(-time.Hour).UnixMilli()),
				},
				wantCode: constants.ErrExpirationTimestampInThePast.Code,
			},
			{
				name: "language preference without alternate merchant name",
				data: models.IndividualInfo{
					BakongAccountID:    "receivekhqr@dvpy",
					Currency:           constants.CurrencyUSD,
					MerchantName:       "test",
					LanguagePreference: util.Ptr("km"),
				},
				wantCode: constants.ErrMerchantNameAlternateLanguageRequired.Code,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				_, err := GenerateIndividual(tt.data)
				if err == nil {
					t.Fatal("GenerateIndividual() error = nil, want error")
				}

				code, ok := err.(*constants.ErrorCode)
				if !ok {
					t.Fatalf("error type = %T, want *constants.ErrorCode", err)
				}
				if code.Code != tt.wantCode {
					t.Errorf("error code = %d (%q), want %d", code.Code, code.Message, tt.wantCode)
				}
			})
		}
	})
}

func TestGenerateMerchant(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		res, err := GenerateMerchant(models.MerchantInfo{
			MerchantID: "M123456",
			IndividualInfo: models.IndividualInfo{
				BakongAccountID: "merchant@dvpy",
				Currency:        constants.CurrencyUSD,
				MerchantName:    "Merchant Co",
				AcquiringBank:   util.Ptr("Acme Bank"),
			},
		})
		if err != nil {
			t.Fatalf("GenerateMerchant() error = %v", err)
		}

		decoded := DecodeKHQR(res.QR)
		if decoded.MerchantType != constants.MerchantAccountInformationMerchant {
			t.Errorf("MerchantType = %q, want %q", decoded.MerchantType, constants.MerchantAccountInformationMerchant)
		}
		if decoded.MerchantID == nil || *decoded.MerchantID != "M123456" {
			t.Errorf("MerchantID = %v, want %q", decoded.MerchantID, "M123456")
		}
		if decoded.AcquiringBank == nil || *decoded.AcquiringBank != "Acme Bank" {
			t.Errorf("AcquiringBank = %v, want %q", decoded.AcquiringBank, "Acme Bank")
		}

		if _, err := DecodeKHQRValidation(res.QR); err != nil {
			t.Errorf("DecodeKHQRValidation() error = %v, want valid QR", err)
		}
	})

	t.Run("missing acquiring bank is required for merchant QR", func(t *testing.T) {
		_, err := GenerateMerchant(models.MerchantInfo{
			MerchantID: "M123456",
			IndividualInfo: models.IndividualInfo{
				BakongAccountID: "merchant@dvpy",
				Currency:        constants.CurrencyUSD,
				MerchantName:    "Merchant Co",
			},
		})
		if err == nil {
			t.Fatal("GenerateMerchant() error = nil, want error")
		}

		code, ok := err.(*constants.ErrorCode)
		if !ok {
			t.Fatalf("error type = %T, want *constants.ErrorCode", err)
		}
		if code.Code != constants.ErrAcquiringBankRequired.Code {
			t.Errorf("error code = %d (%q), want %d", code.Code, code.Message, constants.ErrAcquiringBankRequired.Code)
		}
	})
}

func TestDecodeKHQR(t *testing.T) {
	raw := "00020101021129320016receivekhqr@dvpy0108123123125204599953038405802KH5904test6010Phnom Penh63042356"

	decoded := DecodeKHQR(raw)

	if decoded.BakongAccountID != "receivekhqr@dvpy" {
		t.Errorf("BakongAccountID = %q, want %q", decoded.BakongAccountID, "receivekhqr@dvpy")
	}
	if decoded.MerchantName != "test" {
		t.Errorf("MerchantName = %q, want %q", decoded.MerchantName, "test")
	}
	if decoded.CRC != "2356" {
		t.Errorf("CRC = %q, want %q", decoded.CRC, "2356")
	}
}

func TestDecodeKHQRValidation(t *testing.T) {
	validRaw := "00020101021129320016receivekhqr@dvpy0108123123125204599953038405802KH5904test6010Phnom Penh63042356"

	t.Run("valid KHQR passes", func(t *testing.T) {
		decoded, err := DecodeKHQRValidation(validRaw)
		if err != nil {
			t.Fatalf("DecodeKHQRValidation() error = %v", err)
		}
		if decoded.BakongAccountID != "receivekhqr@dvpy" {
			t.Errorf("BakongAccountID = %q, want %q", decoded.BakongAccountID, "receivekhqr@dvpy")
		}
	})

	t.Run("empty string is invalid", func(t *testing.T) {
		if _, err := DecodeKHQRValidation(""); err == nil {
			t.Error("DecodeKHQRValidation(\"\") error = nil, want error")
		}
	})

	t.Run("too short to contain a CRC is invalid", func(t *testing.T) {
		if _, err := DecodeKHQRValidation("00"); err == nil {
			t.Error("DecodeKHQRValidation() error = nil, want error")
		}
	})

	t.Run("tampered CRC is invalid", func(t *testing.T) {
		tampered := validRaw[:len(validRaw)-4] + "0000"

		_, err := DecodeKHQRValidation(tampered)
		if err == nil {
			t.Fatal("DecodeKHQRValidation() error = nil, want error for tampered CRC")
		}
		if err.Error() != constants.ErrKHQRInvalid.Message {
			t.Errorf("error = %q, want %q", err.Error(), constants.ErrKHQRInvalid.Message)
		}
	})

	t.Run("missing Bakong account id is invalid", func(t *testing.T) {
		payload := tlv(constants.PayloadFormatIndicator, "01") +
			tlv(constants.PointOfInitiationMethod, constants.StaticQR) +
			tlv(constants.MerchantCategoryCodeTag, "5999") +
			tlv(constants.TransactionCurrencyTag, "840") +
			tlv(constants.CountryCodeTag, "KH") +
			tlv(constants.MerchantNameTag, "test") +
			tlv(constants.MerchantCityTag, "Phnom Penh")

		raw := withCRC(payload)

		_, err := DecodeKHQRValidation(raw)
		if err == nil {
			t.Fatal("DecodeKHQRValidation() error = nil, want error")
		}
		if err.Error() != constants.ErrBakongAccountIDRequired.Message {
			t.Errorf("error = %q, want %q", err.Error(), constants.ErrBakongAccountIDRequired.Message)
		}
	})

	t.Run("missing merchant name is invalid", func(t *testing.T) {
		gui := tlv(constants.BakongAccountIdentifier, "receivekhqr@dvpy")

		payload := tlv(constants.PayloadFormatIndicator, "01") +
			tlv(constants.PointOfInitiationMethod, constants.StaticQR) +
			tlv(constants.MerchantAccountInformationIndividual, gui) +
			tlv(constants.MerchantCategoryCodeTag, "5999") +
			tlv(constants.TransactionCurrencyTag, "840") +
			tlv(constants.CountryCodeTag, "KH") +
			tlv(constants.MerchantCityTag, "Phnom Penh")

		raw := withCRC(payload)

		_, err := DecodeKHQRValidation(raw)
		if err == nil {
			t.Fatal("DecodeKHQRValidation() error = nil, want error")
		}
		if err.Error() != constants.ErrMerchantNameRequired.Message {
			t.Errorf("error = %q, want %q", err.Error(), constants.ErrMerchantNameRequired.Message)
		}
	})

	t.Run("dynamic QR without transaction amount is invalid", func(t *testing.T) {
		gui := tlv(constants.BakongAccountIdentifier, "receivekhqr@dvpy")

		payload := tlv(constants.PayloadFormatIndicator, "01") +
			tlv(constants.PointOfInitiationMethod, constants.DynamicQR) +
			tlv(constants.MerchantAccountInformationIndividual, gui) +
			tlv(constants.MerchantCategoryCodeTag, "5999") +
			tlv(constants.TransactionCurrencyTag, "840") +
			tlv(constants.CountryCodeTag, "KH") +
			tlv(constants.MerchantNameTag, "test") +
			tlv(constants.MerchantCityTag, "Phnom Penh")

		raw := withCRC(payload)

		_, err := DecodeKHQRValidation(raw)
		if err == nil {
			t.Fatal("DecodeKHQRValidation() error = nil, want error")
		}
		if err.Error() != constants.ErrInvalidDynamicKHQR.Message {
			t.Errorf("error = %q, want %q", err.Error(), constants.ErrInvalidDynamicKHQR.Message)
		}
	})

	t.Run("expired dynamic QR is invalid", func(t *testing.T) {
		gui := tlv(constants.BakongAccountIdentifier, "receivekhqr@dvpy")
		pastExpiry := "1700000000000" // 13-digit timestamp safely in the past

		payload := tlv(constants.PayloadFormatIndicator, "01") +
			tlv(constants.PointOfInitiationMethod, constants.DynamicQR) +
			tlv(constants.MerchantAccountInformationIndividual, gui) +
			tlv(constants.MerchantCategoryCodeTag, "5999") +
			tlv(constants.TransactionCurrencyTag, "840") +
			tlv(constants.TransactionAmountTag, "10.00") +
			tlv(constants.CountryCodeTag, "KH") +
			tlv(constants.MerchantNameTag, "test") +
			tlv(constants.MerchantCityTag, "Phnom Penh") +
			tlv(constants.TimestampTag, tlv(constants.ExpirationTimestamp, pastExpiry))

		raw := withCRC(payload)

		_, err := DecodeKHQRValidation(raw)
		if err == nil {
			t.Fatal("DecodeKHQRValidation() error = nil, want error")
		}
		if err.Error() != constants.ErrKHQRExpired.Message {
			t.Errorf("error = %q, want %q", err.Error(), constants.ErrKHQRExpired.Message)
		}
	})
}

func TestGenerateThenDecodeThenValidate_RoundTrip(t *testing.T) {
	exp := time.Now().Add(24 * time.Hour).UnixMilli()

	res, err := GenerateIndividual(models.IndividualInfo{
		BakongAccountID:     "receivekhqr@dvpy",
		Currency:            constants.CurrencyUSD,
		MerchantName:        "Round Trip Shop",
		MerchantCity:        "Phnom Penh",
		Amount:              util.Ptr("19.99"),
		ExpirationTimestamp: util.Ptr(exp),
	})
	if err != nil {
		t.Fatalf("GenerateIndividual() error = %v", err)
	}

	decoded, err := DecodeKHQRValidation(res.QR)
	if err != nil {
		t.Fatalf("DecodeKHQRValidation() error = %v", err)
	}

	if decoded.MerchantName != "Round Trip Shop" {
		t.Errorf("MerchantName = %q, want %q", decoded.MerchantName, "Round Trip Shop")
	}
	if decoded.TransactionAmount == nil || *decoded.TransactionAmount != "19.99" {
		t.Errorf("TransactionAmount = %v, want %q", decoded.TransactionAmount, "19.99")
	}
}
