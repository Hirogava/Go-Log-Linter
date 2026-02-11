package testdata

import "log/slog"

func TestSensitiveBad() {
	logger := slog.Default()
	logger.Info("user password is incorrect")        // want "log message contains sensitive keyword: password"
	logger.Error("api_key validation failed")        // want "log message contains sensitive keyword: api_key"
	logger.Warn("secret token not found in request") // want "log message contains sensitive keyword: secret"
}
