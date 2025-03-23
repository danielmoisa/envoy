package model

import (
	"strings"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestUser_HashPassword(t *testing.T) {
	tests := []struct {
		name          string
		plainPassword string
		wantErr       bool
		errMessage    string
	}{
		{
			name:          "Valid password",
			plainPassword: "MySecurePass123!",
			wantErr:       false,
		},
		{
			name:          "Empty password",
			plainPassword: "",
			wantErr:       true,
			errMessage:    "password cannot be empty",
		},
		{
			name:          "Password with whitespace",
			plainPassword: "  MySecurePass123!  ",
			wantErr:       false,
		},
		{
			name:          "Very long password",
			plainPassword: strings.Repeat("a", 72), // bcrypt max length
			wantErr:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{}
			hash1, err1 := u.HashPassword(tt.plainPassword)

			if (err1 != nil) != tt.wantErr {
				t.Errorf("HashPassword() error = %v, wantErr %v", err1, tt.wantErr)
				return
			}

			if tt.wantErr {
				if err1.Error() != tt.errMessage {
					t.Errorf("HashPassword() error message = %v, want %v", err1.Error(), tt.errMessage)
				}
				return
			}

			// Verify the hash is valid
			err := bcrypt.CompareHashAndPassword([]byte(hash1), []byte(strings.TrimSpace(tt.plainPassword)))
			if err != nil {
				t.Errorf("HashPassword() generated invalid hash: %v", err)
			}

			// Verify salt is working (two hashes of same password should be different)
			hash2, _ := u.HashPassword(tt.plainPassword)
			if hash1 == hash2 {
				t.Error("HashPassword() generated identical hashes for same password")
			}
		})
	}
}
