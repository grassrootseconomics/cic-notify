package locale

import (
	"github.com/kamikazechaser/locale"
)

// All message translation strings should be defined in the map here.
var (
	localeMap = locale.LangMap{
		"eng": locale.Map{
			"failed":          "Failed {{ .FailReason }} Contact 0757628885",
			"successReceived": "{{ .ShortHash }} Confirmed {{ .TransferValue }} {{ .VoucherSymbol }} received from {{ .ReceivedFrom }} {{ .DateString }} Balance {{ .CurrentBalance }} {{ .VoucherSymbol }}",
			"successSent":     "{{ .ShortHash }} Confirmed {{ .TransferValue }} {{ .VoucherSymbol }} sent to {{ .SentTo }} {{ .DateString }} Balance {{ .CurrentBalance }} {{ .VoucherSymbol }}",
		},
		"swa": locale.Map{
			"failed":          "Imeshindwa {{ .FailReason }} Wasiliana na 0757628885",
			"successReceived": "{{ .ShortHash }} Imethibitishwa {{ .TransferValue }} {{ .VoucherSymbol }} imepokelewa kutoka kwa {{ .ReceivedFrom }} {{ .DateString }} Salio {{ .CurrentBalance }} {{ .VoucherSymbol }}",
			"successSent":     "{{ .ShortHash }} Imethibitishwa {{ .TransferValue }} {{ .VoucherSymbol }} imetumwa kwa {{ .SentTo }} {{ .DateString }} Salio {{ .CurrentBalance }} {{ .VoucherSymbol }}",
		},
		"mij": locale.Map{
			"failed":          "Vishindikana {{ .FailReason }} Pigira 0757628885",
			"successReceived": "{{ .ShortHash }} Kuthibitiswa {{ .TransferValue }} {{ .VoucherSymbol }} phokera kulaa {{ .ReceivedFrom }} {{ .DateString }} Saliyo {{ .CurrentBalance }} {{ .VoucherSymbol }}",
			"successSent":     "{{ .ShortHash }} kuthibitiswa {{ .TransferValue }} {{ .VoucherSymbol }} ihumwa kwa {{ .SentTo }} {{ .DateString }} saliyo {{ .CurrentBalance }} {{ .VoucherSymbol }}",
		},
		"kam": locale.Map{
			"failed":          "Ndenetekelwa {{ .FailReason }} Kunia 0757628885",
			"successReceived": "{{ .ShortHash }} Nimbeteklye{{ .TransferValue }} {{ .VoucherSymbol }} niwakwata kuma {{ .ReceivedFrom }} {{ .DateString }} Balanci {{ .CurrentBalance }} {{ .VoucherSymbol }}",
			"successSent":     "{{ .ShortHash }} Nwakwata {{ .TransferValue }} {{ .VoucherSymbol }} toma kwa {{ .SentTo }} {{ .DateString }} Balanci {{ .CurrentBalance }} {{ .VoucherSymbol }}",
		},
		"luo": locale.Map{
			"failed":          "Otamre {{ .FailReason }} Goch ni 0757628885",
			"successReceived": "{{ .ShortHash }} Oyudre {{ .TransferValue}} {{ .VoucherSymbol }} yudo kwuom {{ .ReceivedFrom }} {{ .DateString }} Omenda modong {{ .CurrentBalance }} {{ .VoucherSymbol }}",
			"successSent":     "{{ .ShortHash }} Oyudre {{ .TransferValue }} {{ .VoucherSymbol }} oro ni {{ .SentTo }} {{ .DateString }} Omenda modong {{ .CurrentBalance }} {{ .VoucherSymbol }}",
		},
		"kik": locale.Map{
			"failed":          "Niyarema {{ .FailReason }} Hurira 0757628885",
			"successReceived": "{{ .ShortHash }} Gwitikirika {{ .TransferValue }} {{ .VoucherSymbol }} Kuuma kwi {{ .ReceivedFrom }} {{ .DateString }} Matigari {{ .CurrentBalance }} {{ .VoucherSymbol }}",
			"successSent":     "{{ .ShortHash }} Gwitikirika {{ .TransferValue }} {{ .VoucherSymbol }} gutuma kwi {{ .SentTo }} {{ .DateString }} Matigari {{ .CurrentBalance }} {{ .VoucherSymbol }}",
		},
		"gax": locale.Map{
			"failed":          "Kufe {{ .FailReason }} Contact 0757628885",
			"successReceived": "{{ .ShortHash }} Mirkanesse {{ .TransferValue }} {{ .VoucherSymbol }} irraa argame {{ .ReceivedFrom }} {{ .DateString }} Balansi {{ .CurrentBalance }} {{ .VoucherSymbol }}",
			"successSent":     "{{ .ShortHash }} Mirkanesse {{ .TransferValue }} {{ .VoucherSymbol }} Ya ergan garaa {{ .SentTo }} {{ .DateString }} Balansi {{ .CurrentBalance }} {{ .VoucherSymbol }}",
		},
	}
)
