package testdata

import "go.uber.org/zap"

func TestZapGood() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
}
