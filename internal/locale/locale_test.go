package locale

import (
	"testing"
)

func Test_Templates_PrepareLocale(t *testing.T) {
	tmpl, err := InitTemplates()
	if err != nil {
		t.Fatalf("Failed to init i18n templates %v", err)
	}

	// Test successful preparation of SuccessReceivedTemplate
	receivedPayload := struct {
		ShortHash      string
		TransferValue  uint64
		VoucherSymbol  string
		ReceivedFrom   string
		DateString     string
		CurrentBalance uint64
	}{
		"1234XXX",
		1000,
		"SRF",
		"Sohail 0722000000",
		"23/4/23 at 10:09 AM",
		1100,
	}

	expectedReceived := "1234XXX Confirmed received 1000 SRF from Sohail 0722000000 23/4/23 at 10:09 AM Balance 1100 SRF"
	result := tmpl.PrepareLocale(SuccessReceivedTemplate, "eng", receivedPayload)
	if result != expectedReceived {
		t.Errorf("Expected '%s', but got '%s'", expectedReceived, result)
	}

	// Test sw-KE sample template
	expectedReceivedSw := "1234XXX Limekubalika umepokea 1000 SRF kutoka kwa Sohail 0722000000 23/4/23 at 10:09 AM Salio 1100 SRF"
	resultSw := tmpl.PrepareLocale(SuccessReceivedTemplate, "swa", receivedPayload)
	if resultSw != expectedReceivedSw {
		t.Errorf("Expected '%s', but got '%s'", expectedReceivedSw, resultSw)
	}

	// Test successful preparation of SuccessSentTemplate
	sentPayload := struct {
		ShortHash      string
		TransferValue  uint64
		VoucherSymbol  string
		SentTo         string
		DateString     string
		CurrentBalance uint64
	}{
		"5555XXX",
		1000,
		"SRF",
		"Sohail 0722000000",
		"24/4/23 at 10:09 AM",
		900,
	}
	expectedSent := "5555XXX Confirmed 1000 SRF sent to Sohail 0722000000 24/4/23 at 10:09 AM Balance 900 SRF"
	result = tmpl.PrepareLocale(SuccessSentTemplate, "eng", sentPayload)
	if result != expectedSent {
		t.Errorf("Expected '%s', but got '%s'", expectedSent, result)
	}

	// Test successful preparation of FailedTemeplate
	payload := struct{ FailReason string }{"insufficient funds"}
	expected := "Transaction failed. insufficient funds"
	result = tmpl.PrepareLocale(FailedTemeplate, "eng", payload)
	if result != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, result)
	}
}
