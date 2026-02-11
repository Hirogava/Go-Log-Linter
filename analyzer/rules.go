package analyzer

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

type LowercaseRule struct{}

type EnglishRule struct{}

type SpecialCharRule struct{}

type SensitiveRule struct{}

type Rule interface {
	Check(message string) error
}

var rules = []Rule{
	LowercaseRule{},
	EnglishRule{},
	SpecialCharRule{},
	SensitiveRule{},
}

func (r LowercaseRule) Check(msg string) error {
	if len(msg) == 0 {
		return nil
	}

	first := rune(msg[0])
	if unicode.IsUpper(first) {
		return errors.New("log message must start with lowercase letter")
	}

	return nil
}

func (r EnglishRule) Check(msg string) error {
    for _, r := range msg {
		if unicode.In(r, unicode.Cyrillic) {
			return errors.New("log message must be in English")
		}
	}
	
	return nil
}

func (r SpecialCharRule) Check(msg string) error {
    for _, r := range msg {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.IsSpace(r) {
			continue
		}

		return errors.New("log message contains special characters")
	}

	return nil
}

var sensitiveKeywords = []string{
    "password",
    "token",
    "api_key",
    "secret",
}

func (r SensitiveRule) Check(msg string) error {
    lower := strings.ToLower(msg)
	for _, word := range sensitiveKeywords {
		if strings.Contains(lower, word) {
			return fmt.Errorf("log message contains sensitive keyword: %s", word)
		}
	}

	return nil
}
