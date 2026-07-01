package generator

func NewMerchantCity(v *string) KHQRBuilder {
	return newBaseMerchantCode(merchantCityCD, v, true)
}
