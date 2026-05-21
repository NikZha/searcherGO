package main

import "testing"

func TestGetEmail(t *testing.T) {
	htmlBody := "<a href=\"mailto:test@domen.test\">Send email</a>:"
	arrayEmails := getEmail(htmlBody)
	if len(arrayEmails) != 1 {
		t.Errorf("len(arrayEmails) = %d, want 1", len(arrayEmails))
	}
	if arrayEmails[0] != "test@domen.test" {
		t.Errorf("arrayEmails[0] = %s; want test@domen.test", arrayEmails[0])
	}
}

func TestClearCoincidencesEmails(t *testing.T) {
	arrayEmails := []string{"test@domen.test", "test@domen.test", "1@1.1", "1@1.1"}
	arrayEmails = clearCoincidencesEmails(arrayEmails)
	if len(arrayEmails) != 2 {
		t.Errorf("len(arrayEmails) = %d, want 2", len(arrayEmails))
	}
	if arrayEmails[0] != "test@domen.test" {
		t.Errorf("arrayEmails[0] = %s; want test@domen.test", arrayEmails[0])
	}
	if arrayEmails[1] != "1@1.1" {
		t.Errorf("arrayEmails[1] = %s; want 1@1.1", arrayEmails[1])
	}
}
