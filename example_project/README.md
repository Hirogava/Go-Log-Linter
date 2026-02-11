# Example Project - Go-Log-Linter Demo

Пример реального проекта с логиро­ванием для демонстрации Go-Log-Linter.

## Структура

```
example_project/
├── cmd/
│   └── api/
│       └── main.go           # Примеры с ОШИБКАМИ логирования
├── internal/
│   └── services/
│       ├── database.go       # Примеры с ОШИБКАМИ логирования
│       └── auth.go           # Примеры с ПРАВИЛЬНЫМ логированием
├── go.mod
└── README.md
```

## Найденные ошибки

### main.go
- Line 13: `"Starting API server..."` - заглавная буква
- Line 16: `"Ошибка при инициализации БД"` - кириллица
- Line 19: `"Invalid email @format..."` - спецсимвол `@`
- Line 22: `"user token is incorrect"` - конфиденциальные данные `token`

### database.go
- Line 18: `"Connecting to PostgreSQL..."` - заглавная буква
- Line 21: `"connection #string:..."` - спецсимвол `#`
- Line 24: `"database password validation..."` - конфиденциальные данные `password`
- Line 33: `"executing query @%s"` - спецсимвол `@`
- Line 40: `"Закрытие соединения с БД"` - кириллица

### auth.go
- Все примеры правильные!

## Проверка с Go-Log-Linter

### Способ 1: CLI инструмент

```bash
# Из корневой папки Go-Log-Linter
make build

# Проверить пример проекта
./bin/loglint ./example_project/cmd/...
./bin/loglint ./example_project/internal/...

# Проверить конкретный файл
./bin/loglint ./example_project/cmd/api/main.go
```

**Результат:**
```
example_project/cmd/api/main.go:13:2: log message must start with lowercase letter
example_project/cmd/api/main.go:16:2: log message must be in English
example_project/cmd/api/main.go:19:2: log message contains special characters
example_project/cmd/api/main.go:22:2: log message contains sensitive keyword: token

example_project/internal/services/database.go:18:2: log message must start with lowercase letter
example_project/internal/services/database.go:21:2: log message contains special characters
example_project/internal/services/database.go:24:2: log message contains sensitive keyword: password
example_project/internal/services/database.go:33:2: log message contains special characters
example_project/internal/services/database.go:40:2: log message must be in English
```

### Способ 2: Встроенные исправления (SuggestedFix)

```bash
# GO_DEBUG=stderr ./bin/loglint ./example_project/cmd/api/main.go
# будет показывать рекомендуемые исправления
```

## Как исправить ошибки

### Before (неправильно)
```go
logger.Info("Starting API server on port 8080")
logger.Error("Database connection failed")
logger.Warn("Invalid email @format in request")
logger.Info("user token is incorrect")
```

### After (правильно)
```go
logger.Info("starting api server on port 8080")
logger.Error("database connection failed")
logger.Warn("invalid email format in request")
logger.Info("user authentication required")
```

## Запомнить

1. **Нижний регистр**: Все сообщения должны начинаться со строчной буквы
2. **Английский язык**: Только английский, без кириллицы
3. **Спецсимволы**: Разрешены только `.`, `-`, `_`, `:`
4. **Конфиденциальность**: Избегайте слов: `password`, `token`, `api_key`, `secret`

## Примеры правильного логирования

```go
// ПРАВИЛЬНО
logger.Info("server started successfully")
logger.Error("database connection timeout")
logger.Debug("user authentication failed")
logger.Warn("cache invalidation detected")
logger.Info("process-timeout_value: 30s")
logger.Info("config.db.host: localhost")

// НЕПРАВИЛЬНО
logger.Info("Server started successfully")          // Заглавная буква
logger.Error("Ошибка подключения к БД")            // Кириллица
logger.Warn("Invalid email @format")                // Спецсимвол @
logger.Info("user password is incorrect")           // Конфиденциальные данные
logger.Error("connection #failed #retry")           // Спецсимволы #
```
