package rules

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

// RuleConfig интерфейс для конфигурации линтера (из пакета analyzer)
type RuleConfig interface {
	IsStrictMode() bool
	IsIgnoreNumbers() bool
	IsAllowExclamation() bool
	GetAllowedSpecialChars() string
	GetCustomSensitive() []string
}

// LowercaseRule проверяет что сообщение начинается со строчной буквы
type LowercaseRule struct {
	allowExclamation bool
}

// EnglishRule проверяет что сообщение на английском
type EnglishRule struct{}

// SpecialCharRule проверяет спецсимволы
type SpecialCharRule struct {
	allowedChars map[rune]bool
}

// SensitiveRule проверяет чувствительные данные
type SensitiveRule struct {
	keywords map[string]bool
}

// Rule интерфейс для всех правил проверки
type Rule interface {
	Check(message string) error
	GetSuggestedFix(message string) string
}

// Checker создает правила на основе конфига
type Checker struct {
	config RuleConfig
	rules  []Rule
}

// NewChecker создает новый Checker с конфигом
func NewChecker(cfg interface{}) *Checker {
	config, ok := cfg.(RuleConfig)
	if !ok {
		// Если не RuleConfig, создаем с умолчиями
		config = &defaultConfig{}
	}

	return &Checker{
		config: config,
		rules: []Rule{
			&LowercaseRule{allowExclamation: config.IsAllowExclamation()},
			&EnglishRule{},
			NewSpecialCharRule(config.GetAllowedSpecialChars()),
			NewSensitiveRule(config.GetCustomSensitive()),
		},
	}
}

// Rules возвращает список правил
func (c *Checker) Rules() []Rule {
	return c.rules
}

// NewSpecialCharRule создает правило для специальных символов с доп. разрешенными
func NewSpecialCharRule(extraChars string) Rule {
	allowed := map[rune]bool{
		'.': true,
		'-': true,
		'_': true,
		':': true,
	}

	// Добавляем дополнительные символы из конфига
	for _, r := range extraChars {
		allowed[r] = true
	}

	return &SpecialCharRule{allowedChars: allowed}
}

// NewSensitiveRule создает правило для чувствительных данных с доп. keywords
func NewSensitiveRule(customKeywords []string) Rule {
	keywords := map[string]bool{
		"password": true,
		"token":    true,
		"api_key":  true,
		"secret":   true,
	}

	// Добавляем кастомные ключевые слова
	for _, kw := range customKeywords {
		keywords[strings.ToLower(kw)] = true
	}

	return &SensitiveRule{keywords: keywords}
}

// ===== LOWERCASE RULE =====

func (r *LowercaseRule) Check(msg string) error {
	if len(msg) == 0 {
		return nil
	}

	first := rune(msg[0])
	if unicode.IsUpper(first) {
		return errors.New("log message must start with lowercase letter")
	}

	return nil
}

func (r *LowercaseRule) GetSuggestedFix(msg string) string {
	if len(msg) == 0 {
		return msg
	}

	return strings.ToLower(string(msg[0])) + msg[1:]
}

// ===== ENGLISH RULE =====

func (r *EnglishRule) Check(msg string) error {
	for _, ch := range msg {
		if unicode.In(ch, unicode.Cyrillic) {
			return errors.New("log message must be in English")
		}
	}

	return nil
}

func (r *EnglishRule) GetSuggestedFix(msg string) string {
	return msg
}

// ===== SPECIAL CHAR RULE =====

func (r *SpecialCharRule) Check(msg string) error {
	for _, ch := range msg {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) || unicode.IsSpace(ch) {
			continue
		}

		if r.allowedChars[ch] {
			continue
		}

		return errors.New("log message contains special characters")
	}

	return nil
}

func (r *SpecialCharRule) GetSuggestedFix(msg string) string {
	return msg
}

// ===== SENSITIVE RULE =====

func (r *SensitiveRule) Check(msg string) error {
	lower := strings.ToLower(msg)
	for keyword := range r.keywords {
		if strings.Contains(lower, keyword) {
			return fmt.Errorf("log message contains sensitive keyword: %s", keyword)
		}
	}

	return nil
}

func (r *SensitiveRule) GetSuggestedFix(msg string) string {
	return msg
}

// ===== BACKWARD COMPATIBILITY =====

// Rules хранит старые правила для обратной совместимости
var Rules = []Rule{
	&LowercaseRule{},
	&EnglishRule{},
	NewSpecialCharRule(""),
	NewSensitiveRule(nil),
}

// defaultConfig для использования когда конфиг не передан
type defaultConfig struct{}

func (c *defaultConfig) IsStrictMode() bool             { return false }
func (c *defaultConfig) IsIgnoreNumbers() bool          { return false }
func (c *defaultConfig) IsAllowExclamation() bool       { return false }
func (c *defaultConfig) GetAllowedSpecialChars() string { return "" }
func (c *defaultConfig) GetCustomSensitive() []string   { return nil }
