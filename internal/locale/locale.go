package locale

import (
	"github.com/kataras/i18n"
)

type (
	TemplateType string

	Templates struct {
		i18n *i18n.I18n
	}
)

const (
	FailedTemeplate         TemplateType = "failed"
	SuccessReceivedTemplate TemplateType = "successReceived"
	SuccessSentTemplate     TemplateType = "successSent"
)

func InitTemplates() (*Templates, error) {
	i18N, err := i18n.New(i18n.KV(localeMap), "en-US", "sw-KE")
	if err != nil {
		return nil, err
	}

	return &Templates{
		i18n: i18N,
	}, nil
}

func (l *Templates) PrepareLocale(template TemplateType, lang string, templatePayload interface{}) string {
	var preparedTemplate string

	switch template {
	case FailedTemeplate:
		preparedTemplate = l.i18n.Tr(langCode[lang], "failed", templatePayload)
	case SuccessReceivedTemplate:
		preparedTemplate = l.i18n.Tr(langCode[lang], "successReceived", templatePayload)
	case SuccessSentTemplate:
		preparedTemplate = l.i18n.Tr(langCode[lang], "successSent", templatePayload)
	}

	return preparedTemplate
}
