package generator

import (
	"khqr/constants"
	"khqr/models"
)

type merchantInfoLangTemplate struct {
	langPre                         *string
	merchantNameAltLang             *string
	merchantInfoLangTemplateBuilder []KHQRBuilder
}

func NewMerchantInfoLangTemplate(langPre, merchantNameAltLang, merchantCityAltLang *string) KHQRBuilder {
	return &merchantInfoLangTemplate{
		langPre:             langPre,
		merchantNameAltLang: merchantNameAltLang,
		merchantInfoLangTemplateBuilder: []KHQRBuilder{
			newBaseMerchantCode(languagePreCD, langPre),
			newBaseMerchantCode(merchantNameAltLangCD, merchantNameAltLang),
			newBaseMerchantCode(merchantCityAltLangCD, merchantCityAltLang),
		},
	}
}

func (m *merchantInfoLangTemplate) String() string {

	sub := BatchStringify(m.merchantInfoLangTemplateBuilder)
	if sub == "" {
		return ""
	}

	return models.NewTagLengthValue(constants.MerchantInformationLanguageTemplate, &sub).ToString()
}

func (m *merchantInfoLangTemplate) Validate() *constants.ErrorCode {

	if m.langPre != nil && m.merchantNameAltLang == nil {
		return &constants.ErrMerchantNameAlternateLanguageRequired
	}

	return BatchValidate(m.merchantInfoLangTemplateBuilder)
}
