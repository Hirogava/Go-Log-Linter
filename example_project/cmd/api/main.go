package main

import (
	"log/slog"
	"net/http"
)

func main() {
	logger := slog.Default()

	// ОШИБКА: Заглавная буква
	logger.Info("Starting API server on port 8080")

	// ОШИБКА: Кириллица
	logger.Error("Ошибка при инициализации БД")

	// ОШИБКА: Спецсимволы @
	logger.Warn("Invalid email @format in request")

	// ОШИБКА: Конфиденциальные данные
	logger.Info("user token is incorrect")

	// ПРАВИЛЬНО
	logger.Info("server initialized successfully")
	logger.Error("failed to connect to database")
	logger.Warn("request timeout exceeded")

	server := &http.Server{Addr: ":8080"}
	logger.Error(server.ListenAndServe().Error())
}
