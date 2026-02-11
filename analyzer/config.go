package analyzer

import "flag"

// Config содержит настройки для линтера
type Config struct {
	// AllowExclamation позволяет восклицательные знаки в конце сообщений
	AllowExclamation bool

	// CustomSensitive добавляет свои чувствительные ключевые слова
	CustomSensitive []string

	// AllowedSpecialChars дополнительные разрешённые спецсимволы
	AllowedSpecialChars string

	// StrictMode включает дополнительные проверки
	StrictMode bool

	// IgnoreNumbers игнорирует числа в сообщениях
	IgnoreNumbers bool
}

// DefaultConfig возвращает конфигурацию по умолчанию
func DefaultConfig() *Config {
	return &Config{
		AllowExclamation:    false,
		CustomSensitive:     []string{},
		AllowedSpecialChars: "",
		StrictMode:          false,
		IgnoreNumbers:       false,
	}
}

// RegisterFlags регистрирует флаги для конфигурации
func (c *Config) RegisterFlags(fs *flag.FlagSet) {
	fs.BoolVar(&c.AllowExclamation, "allow-exclamation", false,
		"Allow exclamation marks at the end of log messages")
	fs.StringVar(&c.AllowedSpecialChars, "allowed-special-chars", "",
		"Additional allowed special characters (comma-separated)")
	fs.BoolVar(&c.StrictMode, "strict", false,
		"Enable strict mode with additional checks")
	fs.BoolVar(&c.IgnoreNumbers, "ignore-numbers", false,
		"Ignore numbers in log messages")
}

// ===== RuleConfig Interface Implementation =====

// IsStrictMode возвращает режим strict
func (c *Config) IsStrictMode() bool {
	return c.StrictMode
}

// IsIgnoreNumbers возвращает флаг игнорирования чисел
func (c *Config) IsIgnoreNumbers() bool {
	return c.IgnoreNumbers
}

// IsAllowExclamation возвращает флаг разрешения восклицаний
func (c *Config) IsAllowExclamation() bool {
	return c.AllowExclamation
}

// GetAllowedSpecialChars возвращает дополнительные разрешённые спецсимволы
func (c *Config) GetAllowedSpecialChars() string {
	return c.AllowedSpecialChars
}

// GetCustomSensitive возвращает кастомные чувствительные ключевые слова
func (c *Config) GetCustomSensitive() []string {
	return c.CustomSensitive
}
