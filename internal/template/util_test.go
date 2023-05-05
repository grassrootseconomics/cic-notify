package template

import (
	"testing"
)

func TestTxNotifyTemplates_Prepare(t *testing.T) {
	tmpl := LoadTemplates()

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

	expectedReceived := "1234XXX Confirmed. You have received 1000 SRF from Sohail 0722000000 on 23/4/23 at 10:09 AM. New SRF Balance is 1100."
	result, err := tmpl.Prepare(SuccessReceivedTemplate, receivedPayload)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if result != expectedReceived {
		t.Errorf("Expected '%s', but got '%s'", expectedReceived, result)
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
	expectedSent := "5555XXX Confirmed. 1000 SRF sent to Sohail 0722000000 on 24/4/23 at 10:09 AM. New SRF Balance is 900."
	result, err = tmpl.Prepare(SuccessSentTemplate, sentPayload)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if result != expectedSent {
		t.Errorf("Expected '%s', but got '%s'", expectedSent, result)
	}

	// Test successful preparation of FailedTemeplate
	payload := struct{ FailReason string }{"insufficient funds"}
	expected := "Failed. insufficient funds."
	result, err = tmpl.Prepare(FailedTemeplate, payload)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if result != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, result)
	}
}
