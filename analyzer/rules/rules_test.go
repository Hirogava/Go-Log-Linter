package rules

import (
	"testing"
)

func TestLowercaseRule_Check(t *testing.T) {
	tests := []struct {
		name    string
		msg     string
		wantErr bool
	}{
		{
			name:    "lowercase message",
			msg:     "database connection established",
			wantErr: false,
		},
		{
			name:    "uppercase first letter",
			msg:     "Database connection failed",
			wantErr: true,
		},
		{
			name:    "empty message",
			msg:     "",
			wantErr: false,
		},
		{
			name:    "single uppercase letter",
			msg:     "A",
			wantErr: true,
		},
		{
			name:    "single lowercase letter",
			msg:     "a",
			wantErr: false,
		},
	}

	rule := LowercaseRule{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := rule.Check(tt.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Check() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLowercaseRule_GetSuggestedFix(t *testing.T) {
	tests := []struct {
		name     string
		msg      string
		expected string
	}{
		{
			name:     "uppercase first letter",
			msg:      "Database connection",
			expected: "database connection",
		},
		{
			name:     "already lowercase",
			msg:      "database connection",
			expected: "database connection",
		},
		{
			name:     "empty message",
			msg:      "",
			expected: "",
		},
		{
			name:     "single uppercase letter",
			msg:      "A",
			expected: "a",
		},
	}

	rule := LowercaseRule{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rule.GetSuggestedFix(tt.msg)
			if got != tt.expected {
				t.Errorf("GetSuggestedFix() = %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestEnglishRule_Check(t *testing.T) {
	tests := []struct {
		name    string
		msg     string
		wantErr bool
	}{
		{
			name:    "english only",
			msg:     "database connection established",
			wantErr: false,
		},
		{
			name:    "contains cyrillic",
			msg:     "Ошибка при работе",
			wantErr: true,
		},
		{
			name:    "mixed english and cyrillic",
			msg:     "error: Ошибка",
			wantErr: true,
		},
		{
			name:    "with numbers and special chars",
			msg:     "request-123: finished successfully",
			wantErr: false,
		},
		{
			name:    "empty message",
			msg:     "",
			wantErr: false,
		},
	}

	rule := EnglishRule{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := rule.Check(tt.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Check() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSpecialCharRule_Check(t *testing.T) {
	tests := []struct {
		name    string
		msg     string
		wantErr bool
	}{
		{
			name:    "allowed chars only",
			msg:     "process-name_123: completed",
			wantErr: false,
		},
		{
			name:    "forbidden @ symbol",
			msg:     "invalid @char",
			wantErr: true,
		},
		{
			name:    "forbidden # symbol",
			msg:     "error #tag",
			wantErr: true,
		},
		{
			name:    "forbidden $ symbol",
			msg:     "price $100",
			wantErr: true,
		},
		{
			name:    "allowed dot",
			msg:     "completed successfully.",
			wantErr: false,
		},
		{
			name:    "empty message",
			msg:     "",
			wantErr: false,
		},
	}

	rule := NewSpecialCharRule("")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := rule.Check(tt.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Check() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSensitiveRule_Check(t *testing.T) {
	tests := []struct {
		name    string
		msg     string
		wantErr bool
	}{
		{
			name:    "contains password",
			msg:     "user password incorrect",
			wantErr: true,
		},
		{
			name:    "contains token",
			msg:     "invalid token provided",
			wantErr: true,
		},
		{
			name:    "contains api_key",
			msg:     "api_key validation failed",
			wantErr: true,
		},
		{
			name:    "contains secret",
			msg:     "secret key not found",
			wantErr: true,
		},
		{
			name:    "safe message",
			msg:     "user authentication required",
			wantErr: false,
		},
		{
			name:    "case insensitive - PASSWORD",
			msg:     "user PASSWORD incorrect",
			wantErr: true,
		},
		{
			name:    "empty message",
			msg:     "",
			wantErr: false,
		},
	}

	rule := NewSensitiveRule(nil)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := rule.Check(tt.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Check() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
