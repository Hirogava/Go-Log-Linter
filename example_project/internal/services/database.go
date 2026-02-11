package services

import (
	"context"
	"log/slog"
)

type DatabaseService struct {
	logger *slog.Logger
}

func NewDatabaseService(logger *slog.Logger) *DatabaseService {
	return &DatabaseService{logger: logger}
}

func (ds *DatabaseService) Connect(ctx context.Context) error {
	// ОШИБКА: Заглавная буква
	ds.logger.Info("Connecting to PostgreSQL database")

	// ОШИБКА: Спецсимволы #
	ds.logger.Info("connection #string: postgres://user:pass@localhost")

	// ОШИБКА: Конфиденциальные данные
	ds.logger.Debug("database password validation complete")

	// ПРАВИЛЬНО
	ds.logger.Info("database connection established")
	ds.logger.Info("connection pool size: 10")

	return nil
}

func (ds *DatabaseService) Query(sql string) error {
	// ОШИБКА: Спецсимволы @
	ds.logger.Info("executing query @%s", sql)

	// ПРАВИЛЬНО
	ds.logger.Error("query execution failed: timeout")
	ds.logger.Info("query results: 42 rows")

	return nil
}

func (ds *DatabaseService) Close() error {
	// ОШИБКА: Кириллица
	ds.logger.Info("Закрытие соединения с БД")

	// ПРАВИЛЬНО
	ds.logger.Info("database connection closed")

	return nil
}
