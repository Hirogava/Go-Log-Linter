package testdata

import "log/slog"

func TestSensitiveGood() {
	logger := slog.Default()
	logger.Info("user authentication required")
	logger.Error("authorization failed for user")
	logger.Warn("session validation required")
}
