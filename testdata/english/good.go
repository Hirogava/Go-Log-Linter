package testdata

import "log/slog"

func TestEnglishGood() {
	logger := slog.Default()
	logger.Info("database error occurred")
	logger.Warn("user session expired")
	logger.Error("api request failed")
}
