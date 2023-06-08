package locale

import "github.com/kataras/i18n"

// All message translation strings should be defined in the map here.
var (
	localeMap = i18n.LangMap{
		"en-US": i18n.Map{
			"failed":          "Failed {{ .FailReason }} Contact 0757628885",
			"successReceived": "{{ .ShortHash }} Confirmed {{ .TransferValue }} {{ .VoucherSymbol }} received from {{ .ReceivedFrom }} {{ .DateString }} Balance {{ .CurrentBalance }} {{ .VoucherSymbol }}",
			"successSent":     "{{ .ShortHash }} Confirmed {{ .TransferValue }} {{ .VoucherSymbol }} sent to {{ .SentTo }} {{ .DateString }} Balance {{ .CurrentBalance }} {{ .VoucherSymbol }}",
		},
		"sw-KE": i18n.Map{
			"failed":          "Imeshindwa {{ .FailReason }} Wasiliana na 0757628885",
			"successReceived": "{{ .ShortHash }} Imethibitishwa {{ .TransferValue }} {{ .VoucherSymbol }} imepokelewa kutoka kwa {{ .ReceivedFrom }} {{ .DateString }} Salio {{ .CurrentBalance }} {{ .VoucherSymbol }}",
			"successSent":     "{{ .ShortHash }} Imethibitishwa {{ .TransferValue }} {{ .VoucherSymbol }} imetumwa kwa {{ .SentTo }} {{ .DateString }} Salio {{ .CurrentBalance }} {{ .VoucherSymbol }}",
		},
	}
)
