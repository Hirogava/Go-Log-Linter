package testdata

import "log/slog"

func TestSlogGood() {
	slog.Error("request processing failed")
	slog.Warn("invalid token in header received")
	slog.Info("database migration started successfully")
}
