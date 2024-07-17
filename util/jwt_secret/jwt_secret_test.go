package jwt_secret

import (
	"api/common/common_const"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	SetJwtSecret("your_secret_key")
	token, err := GenerateToken(common_const.LoginUser, 1)
	if err != nil {
		t.Error(err)
	}

	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	if token == "" {
		t.Fatal("Expected token to be not empty")
	}
}

func TestParseToken(t *testing.T) {
	SetJwtSecret("your_secret_key")
	token, _ := GenerateToken(common_const.LoginUser, 1)
	claims, err := ParseToken(token)
	if err != nil {
		t.Error(err)
	}
	t.Logf("UserID: %d", claims.UserID)

	token, _ = GenerateToken(common_const.LoginAdmin, 1)
	claims, err = ParseToken(token)
	if err != nil {
		t.Error(err)
	}
	t.Logf("AdminID: %d", claims.AdminID)
}

func TestParseInvalidToken(t *testing.T) {
	SetJwtSecret("your_secret_key")
	invalidJwtToken :="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjowLCJhZG1pbl9pZCI6MSwiTWFwQ2xhaW1zIjp7ImV4cCI6MTcyMTIwNzI1MywiaXNzIjoiZ2luLWJsb2cifX0.y4Ku16plzvIUUPoCnF08xSG9JAOFgijv83ZNerxjjjo"
	_, err := ParseToken(invalidJwtToken)
	if err == nil {
		t.Error("Expected error parsing invalid token, but got nil")
	}
}
