package generator

func NewCrc(v *string) KHQRBuilder {
	return newBaseMerchantCode(crcCD, v, true)
}
