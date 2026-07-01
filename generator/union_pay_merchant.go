package generator

func NewUnionPayMerchant(v *string) KHQRBuilder {
	return newBaseMerchantCode(unionPayMerchantCD, v)
}
