package testdata

import "log/slog"

func TestLowercaseGood() {
	logger := slog.Default()
	logger.Info("database connection established")
	logger.Warn("user not found in cache")
	logger.Error("api request timeout after 30s")
}
