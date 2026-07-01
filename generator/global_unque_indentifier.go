package generator

import (
	"strings"

	"khqr/constants"
	"khqr/models"
	"khqr/util"
)

type InitGlobalIdentifier struct {
	BakongAccountID    *string
	MerchantID         *string
	AcquiringBank      *string
	AccountInformation *string
	IsMerchant         bool
}
type globalUniqueIdentifier struct {
	InitGlobalIdentifier
	guiBuilder []KHQRBuilder
}

type bakongAccountID struct {
	value *string
}

// initialize
func NewGlobalUniqueIdentifier(data InitGlobalIdentifier) KHQRBuilder {

	builder := []KHQRBuilder{
		newBakongAccountID(data.BakongAccountID),
	}

	if data.IsMerchant {

		builder = append(builder, []KHQRBuilder{
			newBaseMerchantCode(merchantIdCD, data.MerchantID, true),
			newBaseMerchantCode(acquiringBankCD, data.AcquiringBank, true),
		}...)
	} else {

		if data.AccountInformation != nil {
			builder = append(builder, newBaseMerchantCode(accountInformationCD, data.AccountInformation))
		}

		if data.AcquiringBank != nil {
			builder = append(builder, newBaseMerchantCode(acquiringBankCD, data.AcquiringBank, true))
		}
	}

	return &globalUniqueIdentifier{
		InitGlobalIdentifier: data,
		guiBuilder:           builder,
	}
}

func newBakongAccountID(v *string) KHQRBuilder {
	return &bakongAccountID{
		value: v,
	}
}

// convert to string
func (g *globalUniqueIdentifier) String() string {

	tag := constants.MerchantAccountInformationIndividual

	if g.IsMerchant {
		tag = constants.MerchantAccountInformationMerchant
	}

	return models.NewTagLengthValue(tag, util.Ptr(BatchStringify(g.guiBuilder))).ToString()
}

func (b *bakongAccountID) String() string {

	return models.NewTagLengthValue(bakongAccountIDCD.Tag, b.value).ToString()
}

// validate

func (g *globalUniqueIdentifier) Validate() *constants.ErrorCode {

	return BatchValidate(g.guiBuilder)
}
func (b *bakongAccountID) Validate() *constants.ErrorCode {

	err := newBaseMerchantCode(bakongAccountIDCD, b.value, true).Validate()
	if err != nil {
		return err
	}

	if len(strings.Split(*b.value, "@")) < 2 {
		return &constants.ErrBakongAccountIDInvalid
	}

	return nil
}
