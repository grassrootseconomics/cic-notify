package locale

import "github.com/kataras/i18n"

// All message translation strings should be defined in the map here.
var (
	localeMap = i18n.LangMap{
		"en-US": i18n.Map{
			"failed":          "Transaction failed. {{ .FailReason }}",
			"successReceived": "{{ .ShortHash }} Confirmed received {{ .TransferValue }} {{ .VoucherSymbol }} from {{ .ReceivedFrom }} {{ .DateString }} Balance {{ .CurrentBalance }} {{ .VoucherSymbol }}",
			"successSent":     "{{ .ShortHash }} Confirmed {{ .TransferValue }} {{ .VoucherSymbol }} sent to {{ .SentTo }} {{ .DateString }} Balance {{ .CurrentBalance }} {{ .VoucherSymbol }}",
		},
		"sw-KE": i18n.Map{
			"failed":          "Samahani ombi lako haijakamiliki. {{ .FailReason}}",
			"successReceived": "{{ .ShortHash }} Limekubalika umepokea {{ .TransferValue }} {{ .VoucherSymbol }} kutoka kwa {{ .ReceivedFrom }} {{ .DateString }} Salio {{ .CurrentBalance }} {{ .VoucherSymbol }}",
			"successSent":     "{{ .ShortHash }} Limekubalika umetuma {{ .TransferValue }} {{ .VoucherSymbol }} kwenda kwa {{ .SentTo }} {{ .DateString }} Salio {{ .CurrentBalance }} {{ .VoucherSymbol }}",
		},
	}
)
