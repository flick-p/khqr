package generator

func NewPayloadFormatInd(v *string) KHQRBuilder {
	return newBaseMerchantCode(payloadFormatIndCD, v, true)
}
