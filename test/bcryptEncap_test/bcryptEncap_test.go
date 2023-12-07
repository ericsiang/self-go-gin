package bcryptEncap_test

import (
	"api/util/bcryptEncap"
	"testing"
)

func TestGenerateFromPassword(t *testing.T) {
	password := "testPassword"
	hash, err := bcryptEncap.GenerateFromPassword(password)
	if err != nil {
		t.Fatalf("Failed to generate hash from password: %v", err)
	}

	if len(hash) == 0 {
		t.Fatal("Expected hash to be not empty")
	}
}

func TestCompareHashAndPassword(t *testing.T) {
	password := "testPassword"
	hash, _ := bcryptEncap.GenerateFromPassword(password)

	err := bcryptEncap.CompareHashAndPassword(hash, []byte(password))
	if err != nil {
		t.Fatalf("Failed to compare hash and password: %v", err)
	}
}
