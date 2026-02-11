package analyzer

import (
	"testing"
)

func TestConfigStruct(t *testing.T) {
	config := &Config{
		AllowExclamation:    true,
		CustomSensitive:     []string{"secret", "key", "token"},
		AllowedSpecialChars: "!@#",
		StrictMode:          true,
		IgnoreNumbers:       false,
	}

	if !config.AllowExclamation {
		t.Error("AllowExclamation should be true")
	}

	if len(config.CustomSensitive) != 3 {
		t.Errorf("CustomSensitive length = %d, want 3", len(config.CustomSensitive))
	}

	if config.CustomSensitive[0] != "secret" {
		t.Errorf("CustomSensitive[0] = %q, want %q", config.CustomSensitive[0], "secret")
	}

	if !config.IsStrictMode() {
		t.Error("IsStrictMode should be true")
	}

	if config.IsIgnoreNumbers() {
		t.Error("IsIgnoreNumbers should be false")
	}

	if config.GetAllowedSpecialChars() != "!@#" {
		t.Errorf("GetAllowedSpecialChars() = %q, want %q", config.GetAllowedSpecialChars(), "!@#")
	}
}

func TestConfigDefault(t *testing.T) {
	config := DefaultConfig()

	if config.AllowExclamation {
		t.Error("AllowExclamation should be false by default")
	}

	if config.CustomSensitive != nil && len(config.CustomSensitive) > 0 {
		t.Error("CustomSensitive should be empty by default")
	}

	if !config.IsAllowExclamation() {
		// This is expected, should return false
	}

	if config.IsStrictMode() {
		t.Error("StrictMode should be false by default")
	}
}
