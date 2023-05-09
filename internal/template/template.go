package template

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
	failedMsg          = "Failed. {{ .FailReason }}."
	successReceivedMsg = "{{ .ShortHash }} Confirmed {{ .TransferValue }} {{ .VoucherSymbol }} received from {{ .ReceivedFrom }} {{ .DateString }} Balance {{ .CurrentBalance }} {{ .VoucherSymbol }}"
	successSentMsg     = "{{ .ShortHash }} Confirmed {{ .TransferValue }} {{ .VoucherSymbol }} sent to {{ .SentTo }} {{ .DateString }} Balance {{ .CurrentBalance }} {{ .VoucherSymbol }}"

	FailedTemeplate         TemplateType = "failed"
	SuccessReceivedTemplate TemplateType = "successReceived"
	SuccessSentTemplate     TemplateType = "successSent"
)

func LoadTemplates() *TxNotifyTemplates {
	failedTmpl := template.Must(template.New("failed").Parse(failedMsg))
	successReceivedTmpl := template.Must(template.New("successReceived").Parse(successReceivedMsg))
	successSentTmpl := template.Must(template.New("successSent").Parse(successSentMsg))

	return &TxNotifyTemplates{
		failedTmpl:          failedTmpl,
		successReceivedTmpl: successReceivedTmpl,
		successSentTmpl:     successSentTmpl,
	}
}
