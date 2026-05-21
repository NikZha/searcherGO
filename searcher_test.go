package main

import "testing"

func TestGetEmail(t *testing.T) {
	htmlBody := "<a href=\"mailto:test@domen.test\">Send email</a>:"
	arrayEmails := getEmail(htmlBody)
	if len(arrayEmails) != 1 {
		t.Errorf("len(arrayEmails) = %d; want 1\n", len(arrayEmails))
	}
	if arrayEmails[0] != "test@domen.test" {
		t.Errorf("arrayEmails[0] = %s; want test@domen.test\n", arrayEmails[0])
	}
}

func TestGetEmail_BrokenEmail(t *testing.T) {
	htmlBody := "test@domen.t"
	arrayEmails := getEmail(htmlBody)
	if len(arrayEmails) != 0 {
		t.Errorf("len(arrayEmails) = %d; want 0\n", len(arrayEmails))
	}
}

func TestClearCoincidencesEmails(t *testing.T) {
	arrayEmails := []string{"test@domen.test", "test@domen.test", "1@1.1", "1@1.1"}
	arrayEmails = clearCoincidencesEmails(arrayEmails)
	if len(arrayEmails) != 2 {
		t.Errorf("len(arrayEmails) = %d; want 2\n", len(arrayEmails))
	}
	emailTest := false
	emailOne := false
	for _, email := range arrayEmails {
		if email == "test@domen.test" {
			emailTest = true
		}
		if email == "1@1.1" {
			emailOne = true
		}
	}
	if !emailTest {
		t.Errorf("expected 'test@domen.test' not found in %v", arrayEmails)
	}
	if !emailOne {
		t.Errorf("expected '1@1.1' not found in %v", arrayEmails)
	}
}

func TestGetPort(t *testing.T) {
	startPort := 9000
	gotPort := getPort(startPort)

	if gotPort < startPort {
		t.Errorf("getPort(%d) = %d; want >= %d", startPort, gotPort, startPort)
	}

	if gotPort > startPort+1000 {
		t.Errorf("getPort(%d) = %d; want <= %d (or adjust limit)", startPort, gotPort, startPort+1000)
	}
}
