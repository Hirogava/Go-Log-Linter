package services

import (
	"context"
	"log/slog"
)

type AuthService struct {
	logger *slog.Logger
}

func NewAuthService(logger *slog.Logger) *AuthService {
	return &AuthService{logger: logger}
}

func (as *AuthService) Authenticate(ctx context.Context, username string) error {
	// ПРАВИЛЬНО: все в нижнем регистре, английский, без спецсимволов
	as.logger.Info("authentication attempt for user: " + username)
	as.logger.Debug("checking credentials in database")

	// ПРАВИЛЬНО: используются разрешённые спецсимволы
	as.logger.Info("user-session_id: abc123-def456")
	as.logger.Info("role: admin, permissions: read.write")

	// ПРАВИЛЬНО: безопасное логирование без конфиденциальных данных
	as.logger.Info("authentication successful")
	as.logger.Error("authentication failed: invalid credentials")

	return nil
}

func (as *AuthService) RefreshToken(ctx context.Context, userID string) error {
	// ПРАВИЛЬНО
	as.logger.Debug("token refresh requested for user-id: " + userID)
	as.logger.Info("issuing new access token")
	as.logger.Info("token expiration set to 1 hour")

	return nil
}

func (as *AuthService) Logout(ctx context.Context, userID string) error {
	// ПРАВИЛЬНО
	as.logger.Info("user logout: " + userID)
	as.logger.Debug("invalidating all user sessions")
	as.logger.Info("logout completed successfully")

	return nil
}
