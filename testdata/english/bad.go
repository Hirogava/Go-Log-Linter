package testdata

import "log/slog"

func TestEnglishBad() {
	logger := slog.Default()
	logger.Info("Ошибка подключения к базе") // want "log message must be in English"
	logger.Warn("Пользователь не найден")    // want "log message must be in English"
	logger.Error("Ошибка при запросе API")   // want "log message must be in English"
}
