package testdata

import "log/slog"

func TestSpecialBad() {
	logger := slog.Default()
	logger.Info("invalid @character in message") // want "log message contains special characters"
	logger.Warn("error #tag not allowed")        // want "log message contains special characters"
	logger.Error("request $failed unexpectedly") // want "log message contains special characters"
}
