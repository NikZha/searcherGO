package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()
	homeHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("got %d, want %d", rr.Code, http.StatusOK)
	}
}

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

func TestGetBody(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok","site":"example"}`))
	}))
	defer server.Close()
	statusCode, body := getBody(server.URL)
	if statusCode != 200 {
		t.Errorf("status code = %d, want 200", statusCode)
	}
	isExample := strings.Contains(string(body), "example")

	if !isExample {
		t.Errorf("body = %s, want contains 'example", body)
	}
}
