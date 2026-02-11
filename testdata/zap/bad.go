package testdata

import "go.uber.org/zap"

func TestZapBad() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Info("Database connection failed")           // want "log message must start with lowercase letter"
	logger.Warn("User not found")                       // want "log message must start with lowercase letter"
	logger.Sugar().Error("API request timeout")         // want "log message must start with lowercase letter"
}
