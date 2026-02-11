package testdata

import "log/slog"

func TestSpecialGood() {
	logger := slog.Default()
	logger.Info("user operation completed successfully")
	logger.Warn("process-name: database_connection_pool-size: 10")
	logger.Error("request timed out: 30s. retry in 60s.")
}
