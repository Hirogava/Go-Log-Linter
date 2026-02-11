package testdata

import "log/slog"

func TestSlogBad() {
	slog.Error("Ошибка при обработке запроса") // want "log message must be in English"
	slog.Warn("Invalid token in header")       // want "log message must start with lowercase letter"
	slog.Info("Database migration started")    // want "log message must start with lowercase letter"
}
