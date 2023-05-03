package task

import "testing"

func TestFormatDate(t *testing.T) {
	expected := "2022-05-03 05:40:00"
	result := formatDate(1651545600, "Europe/Moscow")
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestFormatIdentifier(t *testing.T) {
	expected := "JOHN DOE 0720000000"
	result := formatIdentifier("John", "Doe", "0720000000", "0x0")
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}

	expected = "DOE 0720000000"
	result = formatIdentifier("", "Doe", "0720000000", "0x0")
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}

	expected = "0x0"
	result = formatIdentifier("", "", "", "0x0")
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestFormatShortHash(t *testing.T) {
	expected := "40BC8DA2"
	result := formatShortHash("0xc92eb48f99e5aa993dbe17b5f742f9a1aa6d57247796dfe8c9e99ba640bc8da2")
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
