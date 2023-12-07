package jwt_secret_test

import (
	"api/util/jwt_secret"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	userId := uint(1)
	token, err := jwt_secret.GenerateToken(userId)

	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	if token == "" {
		t.Fatal("Expected token to be not empty")
	}
}
