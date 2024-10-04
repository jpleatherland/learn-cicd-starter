package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKeyOk(t *testing.T) {
	headers := http.Header{}
	headers.Add("Authorization", "ApiKey testkey")
	expected := "testkey"
	got, err := GetAPIKey(headers)
	if err != nil {
		t.Fatal(err.Error())
	}
	if got != expected {
		t.Fatalf("failure - expected %s, got %s", expected, got)
	}
}

func TestGetAPIKeyMalformed(t *testing.T){
	headers := http.Header{}
	headers.Add("Authorization", "Bearer testkey")
	expected := errors.New("malformed authorization header")
	_, err := GetAPIKey(headers)
	if err.Error() != expected.Error() || err == nil {
		t.Fatalf("incorrect error - expected %s, got %s", expected, err)
	}
}

func TestGetAPIKeyMissing(t *testing.T){
	headers := http.Header{}
	expected := ErrNoAuthHeaderIncluded
	_, err := GetAPIKey(headers)
	if err.Error() != expected.Error() || err == nil {
		t.Fatalf("incorrect error - expected %s, got %s", expected, err)
	}
}

func TestBrokenTest(t *testing.T){
	headers := http.Header{}
	_, err := GetAPIKey(headers)
	if err != nil {
		t.Fatal("failed successfully")
	}
}
