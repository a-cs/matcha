package utils

import (
	"testing"
)

func TestIsValidPassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		expected bool
	}{
		{
			name:     "Valid password with all requirements",
			password: "StrongP@ss1",
			expected: true,
		},
		{
			name:     "Valid password with different special characters",
			password: "MyP@ssw0rd!",
			expected: true,
		},
		{
			name:     "Password too short",
			password: "Abc1@",
			expected: false,
		},
		{
			name:     "Password without uppercase",
			password: "myp@ssw0rd",
			expected: false,
		},
		{
			name:     "Password without lowercase",
			password: "MYP@SSW0RD",
			expected: false,
		},
		{
			name:     "Password without number",
			password: "MyP@ssword",
			expected: false,
		},
		{
			name:     "Password without special character",
			password: "MyPassword1",
			expected: false,
		},
		{
			name:     "Common English word - password",
			password: "Password1@",
			expected: false,
		},
		{
			name:     "Common English word - hello",
			password: "Hello123!",
			expected: false,
		},
		{
			name:     "Common English word - admin",
			password: "Admin123!",
			expected: false,
		},
		{
			name:     "Empty password",
			password: "",
			expected: false,
		},
		{
			name:     "Password with spaces",
			password: "My P@ss 1",
			expected: true,
		},
		{
			name:     "Password with unicode characters",
			password: "Myp@ss1ñ",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidPassword(tt.password)
			if result != tt.expected {
				t.Errorf("IsValidPassword(%q) = %v, expected %v", tt.password, result, tt.expected)
			}
		})
	}
}
