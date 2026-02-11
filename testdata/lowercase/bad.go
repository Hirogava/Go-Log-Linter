package testdata

import "log/slog"

func TestLowercaseBad() {
	logger := slog.Default()
	logger.Info("Database connection failed") // want "log message must start with lowercase letter"
	logger.Warn("User not found")             // want "log message must start with lowercase letter"
	logger.Error("API request timeout")       // want "log message must start with lowercase letter"
}
