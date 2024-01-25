package locale

import (
	"github.com/kamikazechaser/locale"
)

type (
	TemplateType string

	Templates struct {
		locale *locale.Locale
	}
)

const (
	FailedTemeplate         TemplateType = "failed"
	SuccessReceivedTemplate TemplateType = "successReceived"
	SuccessSentTemplate     TemplateType = "successSent"
)

func InitTemplates() (*Templates, error) {
	l, err := locale.NewLocale(localeMap, "swa")
	if err != nil {
		return nil, err
	}

	return &Templates{
		locale: l,
	}, nil
}

func (l *Templates) PrepareLocale(template TemplateType, lang string, templatePayload interface{}) (string, error) {
	return l.locale.Render(string(template), locale.WithLangCode(lang), locale.WithPayload(templatePayload))
}
