package generator

func NewCountryCode(v *string) KHQRBuilder {

	return newBaseMerchantCode(countryCodeCD, v, true)
}
