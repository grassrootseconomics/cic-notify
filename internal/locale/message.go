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
		"fr-FR": i18n.Map{
			"failed":          "Échoué {{ .FailReason }} Contactez 0757628885",
			"successReceived": "{{ .ShortHash }} Confirmé {{ .TransferValue }} {{ .VoucherSymbol }} reçu de {{ .ReceivedFrom }} {{ .DateString }} Solde {{ .CurrentBalance }} {{ .VoucherSymbol }}",
			"successSent":     "{{ .ShortHash }} Confirmé {{ .TransferValue }} {{ .VoucherSymbol }} envoyé à {{ .SentTo }} {{ .DateString }} Solde {{ .CurrentBalance }} {{ .VoucherSymbol }}",
		},
		"kam-KE": i18n.Map{
			"failed":          "ndenetekelwa {{ .FailReason }} kunia 0757628885",
			"successReceived": "{{ .ShortHash }}  {{ .TransferValue }} {{ .VoucherSymbol }} niwakwata kuma {{ .ReceivedFrom }} {{ .DateString }} Balanci {{ .CurrentBalance }} {{ .VoucherSymbol }}",
			"successSent":     "{{ .ShortHash }} niwakwata {{ .TransferValue }} {{ .VoucherSymbol }} toma kwa {{ .SentTo }} {{ .DateString }} Balanci {{ .CurrentBalance }} {{ .VoucherSymbol }}",
		},
		"kik-KE": i18n.Map{
			"failed":          "Niyarema {{ .FailReason }} Hurira 0757628885",
			"successReceived": "{{ .ShortHash }} Gwitikirika {{ .TransferValue }} {{ .VoucherSymbol }} Kuuma kwi {{ .ReceivedFrom }} {{ .DateString }} Matigari {{ .CurrentBalance }} {{ .VoucherSymbol }}",
			"successSent":     "{{ .ShortHash }} Gwitikirika {{ .TransferValue }} {{ .VoucherSymbol }} gutuma kwi{{ .SentTo }} {{ .DateString }} Matigari {{ .CurrentBalance }} {{ .VoucherSymbol }}",
		},
		"luo-KE": i18n.Map{
			"failed":          "Otamre {{ .FailReason}} goch ni 0757628885",
			"successReceived": "{{ .ShortHash }} oyudre {{ .TransferValue }} {{ .Vouchersymbol }} yudo kwuom {{ .ReceivedFrom }} {{ .DateString }} Omenda modong {{ .CurrentBalance }} {{ .Vouchersymbol }}",
			"successSent":     "{{ .ShortHash }} oyudre {{ .TransferValue }} {{ .VoucherSymbol }} Oro ni {{ .SentTo }} {{ .DateString }} Omenda modong {{ .CurrentBalance }} {{ .Vouchersymbol }}",
		},
		"mij-KE": i18n.Map{
			"failed":          "Vishindikana {{ .FailReason }} Pigira 0757628885",
			"successReceived": "{{ .ShortHash }} Kuthibitiswa {{ .TransferValue }} {{ .VoucherSymbol }} phokera kulaa {{ .ReceivedFrom }} {{ .DateString }} Saliyo {{ .CurrentBalance }} {{ .VoucherSymbol }}",
			"successSent":     "{{ .ShortHash }} kuthibitiswa {{ .TransferValue }} {{ .VoucherSymbol }} ihumwa kwa {{ .SentTo }} {{ .DateString }} saliyo .{{ .CurrentBalance }} {{ .VoucherSymbol }}",
		},
		"gax-KE": i18n.Map{
			"failed":          "Kufe {{ .FailReason }} Contact 0757628885",
			"successReceived": "{{ .ShortHash }} Mirkanesse{{ .TransferValue }} {{ .VoucherSymbol }} irraa argame {{ .ReceivedFrom }} {{ .DateString }} Balansi {{ .CurrentBalance }} {{ .VoucherSymbol }}",
			"successSent":     "{{ .ShortHash }} Mirkanesse{{ .TransferValue }} {{ .VoucherSymbol }} Ya ergan garaa {{ .SentTo }} {{ .DateString }} Balansi {{ .CurrentBalance }} {{ .VoucherSymbol }}",
		},
	}
)
