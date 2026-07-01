package khqr

import (
	"errors"
	"strings"
	"time"

	"khqr/constants"
	"khqr/decoder"
	"khqr/generator"
	"khqr/models"
	"khqr/util"
)

const crcValueLength = 4

func GenerateIndividual(data models.IndividualInfo) (*models.KHQRData, error) {

	qr, err := generator.NewKHQRGenerator(models.MerchantInfo{
		IndividualInfo: data,
	}, false).Generate()
	if err != nil {
		return nil, err
	}

	return &models.KHQRData{
		QR:  qr,
		MD5: util.CalculateMD5(qr),
	}, nil
}

func GenerateMerchant(data models.MerchantInfo) (*models.KHQRData, error) {
	qr, err := generator.NewKHQRGenerator(data, true).Generate()
	if err != nil {
		return nil, err
	}

	return &models.KHQRData{
		QR:  qr,
		MD5: util.CalculateMD5(qr),
	}, nil
}

func DecodeKHQR(khqrString string) models.DecodedKHQR {

	return decoder.NewDecoder(khqrString).Decode()
}

func DecodeKHQRValidation(khqrString string) (*models.DecodedKHQR, error) {

	decoded := DecodeKHQR(khqrString)

	// Validate CRC to detect corrupted or tampered KHQR strings
	if decoded.CRC == "" || len(khqrString) < crcValueLength {
		return nil, errors.New(constants.ErrKHQRInvalid.Message)
	}

	payload := khqrString[:len(khqrString)-crcValueLength]
	expectedCRC := util.CalculateCRC16(payload)
	if !strings.EqualFold(expectedCRC, decoded.CRC) {
		return nil, errors.New(constants.ErrKHQRInvalid.Message)
	}

	// Validate required fields
	if decoded.BakongAccountID == "" {
		return nil, errors.New(constants.ErrBakongAccountIDRequired.Message)
	}

	if decoded.MerchantName == "" {
		return nil, errors.New(constants.ErrMerchantNameRequired.Message)
	}

	// Validate dynamic QR requirements
	if decoded.PointOfInitiationMethod != nil && *decoded.PointOfInitiationMethod == constants.DynamicQR {
		if decoded.TransactionAmount == nil {
			return nil, errors.New(constants.ErrInvalidDynamicKHQR.Message)
		}
		if decoded.ExpirationTimestamp != nil && *decoded.ExpirationTimestamp < time.Now().UnixMilli() {
			return nil, errors.New(constants.ErrKHQRExpired.Message)
		}
	}

	return &decoded, nil
}
