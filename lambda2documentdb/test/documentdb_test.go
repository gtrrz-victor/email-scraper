package test

import (
	"fmt"
	"rft/lambda2documentdb/internal/store"
	"testing"
)

/**
type Email interface {
	Contact() string
	PhoneNumber() string
	BookingNumber() string
	NumPeople() string
	Name() string
	Tour() string
	Date() string
}
**/

func TestStoreEmptyString(t *testing.T) {
	var err error
	if err = store.Store(""); err == nil {
		t.Fatal("expected store empty json to throw an error")
	}
	const expectedError string = "invalid json format. Body: ''"
	if err.Error() != expectedError {
		t.Fatalf("expected error '%s' to be equals to %s", err, expectedError)
	}
}

func TestStoreInvalidJSON(t *testing.T) {
	const body string = "invalid json"
	var err error
	if err = store.Store(body); err == nil {
		t.Fatal("expected store empty json to throw an error")
	}
	var expectedError string = fmt.Sprintf("invalid json format. Body: '%s'", body)
	if err.Error() != expectedError {
		t.Fatalf("expected error '%s' to be equals to %s", err, expectedError)
	}
}
