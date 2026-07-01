package generator

func NewMerchantName(v *string) KHQRBuilder {
	return newBaseMerchantCode(merchantNameCD, v, true)
}
