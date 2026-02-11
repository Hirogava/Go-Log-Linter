package services

import (
	"go.uber.org/zap"
)

type RequestHandler struct {
	logger *zap.Logger
	sugar  *zap.SugaredLogger
}

func NewRequestHandler(logger *zap.Logger) *RequestHandler {
	return &RequestHandler{
		logger: logger,
		sugar:  logger.Sugar(),
	}
}

// Примеры с zap структурированным логированием
func (rh *RequestHandler) HandleRequest(requestID string) error {
	// ОШИБКА: Заглавная буква
	rh.logger.Info("Processing HTTP request",
		zap.String("request_id", requestID),
	)

	// ПРАВИЛЬНО: нижний регистр
	rh.logger.Info("request started",
		zap.String("request_id", requestID),
	)

	return nil
}

// Примеры с zap Sugar API (форматированное логирование)
func (rh *RequestHandler) ValidateInput(data string) error {
	// ОШИБКА: Кириллица
	rh.sugar.Infof("Валидация входных данных: %s", data)

	// ОШИБКА: Конфиденциальные данные
	rh.sugar.Infof("api_key validation in progress: %s", data)

	// ПРАВИЛЬНО
	rh.sugar.Infof("input validation started for data: %s", data)
	rh.sugar.Infof("validation completed successfully")

	return nil
}

func (rh *RequestHandler) LogError(err error) {
	// ОШИБКА: Спецсимволы @
	rh.sugar.Errorf("Error @occurred: %v", err)

	// ПРАВИЛЬНО
	rh.sugar.Errorf("error occurred: %v", err)
}
