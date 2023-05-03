package templates

import "text/template"

type (
	TemplateType string

	TxNotifyTemplates struct {
		failedTmpl          *template.Template
		successReceivedTmpl *template.Template
		successSentTmpl     *template.Template
	}
)

const (
	FailedTemeplate         TemplateType = "failed.tmpl"
	SuccessReceivedTemplate TemplateType = "success_received.tmpl"
	SuccessSentTemplate     TemplateType = "success_sent.tmpl"
)

func LoadTemplates() *TxNotifyTemplates {
	failedTmpl := template.Must(template.New(string(FailedTemeplate)).ParseFiles(string(FailedTemeplate)))
	successReceivedTmpl := template.Must(template.New(string(SuccessReceivedTemplate)).ParseFiles(string(SuccessReceivedTemplate)))
	successSentTmpl := template.Must(template.New(string(SuccessSentTemplate)).ParseFiles(string(SuccessSentTemplate)))

	return &TxNotifyTemplates{
		failedTmpl:          failedTmpl,
		successReceivedTmpl: successReceivedTmpl,
		successSentTmpl:     successSentTmpl,
	}
}
