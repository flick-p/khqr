package khqr

import (
	"encoding/json"
	"fmt"
	"testing"

	"khqr/constants"
	"khqr/models"
	"khqr/util"
)

func TestIndividualQR(t *testing.T) {

	res, err := GenerateIndividual(models.IndividualInfo{
		BakongAccountID:    "receivekhqr@dvpy",
		Currency:           constants.CurrencyUSD,
		MerchantName:       "test",
		AccountInformation: util.Ptr("12312312"),
	})
	if err != nil {
		t.Error("errGen: ", err)
	}

	stringifyPrint("data", res)
}

func TestDecoder(t *testing.T) {

	res := DecodeKHQR("00020101021129320016receivekhqr@dvpy0108123123125204599953038405802KH5904test6010Phnom Penh63042356")

	stringifyPrint("Decoded", res)
}

func stringifyPrint(key string, v any) {

	data, _ := json.Marshal(v)

	fmt.Printf("key: %s", key)
	fmt.Printf("\nvalue: %s", string(data))
}
