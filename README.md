# Go-Log-Linter

[![CI](https://github.com/Hirogava/Go-Log-Linter/workflows/CI/badge.svg)](https://github.com/Hirogava/Go-Log-Linter/actions)

Мощный анализатор кода для проверки соглашений о логировании в Go. Поддерживает оба логгера: `log/slog` и `go.uber.org/zap` с интерфейсом `Sugar()`.

## Возможности

**Проверка нижнего регистра** - Сообщения лога должны начинаться со строчной буквы
**Проверка английского языка** - Сообщения лога должны быть на английском (без кириллицы)
**Проверка спецсимволов** - Обнаружение запрещённых спецсимволов
**Проверка чувствительных данных** - Предотвращает логирование конфиденциальной информации
**SuggestedFix** - Автоматическое исправление нарушений нижнего регистра
**Интеграция с golangci-lint** - Работает как кастомный плагин линтера
**Поддержка множества логгеров** - `log/slog`, `go.uber.org/zap` и обёртки

## Установка

```bash
go get github.com/Hirogava/Go-Log-Linter
```

## Быстрый старт

### Вариант 1: CLI инструмент (все ОС)

```bash
# Собрать линтер
make build

# Проверить проект из его директории
cd your_project
../bin/loglint ./...
cd ..

# Или проверить конкретный пакет
cd your_project
../bin/loglint ./cmd/...
../bin/loglint ./internal/...
cd ..
```

**Windows (PowerShell):**
```powershell
# Собрать
$env:GOOS="windows"; $env:GOARCH="amd64"; go build -o .\bin\loglint.exe .\cmd\loglint

# Проверить
cd your_project
..\bin\loglint.exe ./...
cd ..
```

**Linux/macOS:**
```bash
# Собрать
GOOS=linux GOARCH=amd64 go build -o ./bin/loglint ./cmd/loglint

# Проверить
cd your_project
../bin/loglint ./...
cd ..
```

### Вариант 2: golangci-lint плагин (Linux/macOS)

```bash
# Собрать плагин
make build-plugin

# Добавить в .golangci.yml (см. конфигурацию ниже)

# Запустить
golangci-lint run
```

### Вариант 3: Программный API

```go
package main

import (
    "github.com/Hirogava/Go-Log-Linter/analyzer"
    "golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
    singlechecker.Main(analyzer.Analyzer)
}
```

### Как отдельный инструмент

```bash
# Собрать CLI линтер
make build

# Использовать
./bin/loglint ./cmd/...
./bin/loglint ./analyzer/...
```

## Конфигурация

### Параметры конфигурации

Линтер поддерживает следующие конфигурационные флаги:

| Флаг | Тип | По умолчанию | Описание |
|------|-----|--------------|---------|
| `allow-exclamation` | bool | false | Разрешить восклицательные знаки в конце логов |
| `allowed-special-chars` | string | "" | Дополнительные разрешённые спецсимволы (например: "!@#") |
| `custom-sensitive` | strings | [] | Дополнительные чувствительные ключевые слова |
| `strict` | bool | false | Строгий режим с дополнительными проверками |
| `ignore-numbers` | bool | false | Игнорировать числа при проверке спецсимволов |

### CLI использование

```bash
# С флагами (при использовании как отдельный инструмент)
go run ./cmd/loglint -allow-exclamation -allowed-special-chars="!@#" ./...

# Или через golangci-lint
golangci-lint run --linter-settings=loglint.allow-exclamation=true ./...
```

### Конфигурация через .golangci.yml

```yaml
linters-settings:
  custom:
    loglint:
      path: ./loglint.so
      description: "Анализатор логов для проверки соглашений о логировании"

linters:
  enable:
    - loglint
```

Для передачи дополнительных параметров:

```yaml
linters:
  enable:
    - loglint

issues:
  exclude-rules:
    - linters:
        - loglint
      text: "special characters"  # Исключить проверку спецсимволов для определённых файлов
```

### Примеры конфигурации

**Разрешить восклицательные знаки:**
```yaml
# .golangci.yml
linters-settings:
  custom:
    loglint:
      path: ./loglint.so
```
*Флаги передаются через CLI: `golangci-lint run --linter-settings 'loglint.allow-exclamation=true'`*

**Добавить собственные чувствительные ключевые слова:**
Отредактируйте конфигурацию в `analyzer/config.go` или передайте через программный API:

```go
import "github.com/Hirogava/Go-Log-Linter/analyzer"

config := &analyzer.Config{
    CustomSensitive: []string{"customer_id", "credit_card"},
}
```

## Правила

### 1. Правило нижнего регистра
Сообщения логов должны начинаться со строчной буквы.

**Неправильно:**
```go
logger.Info("Database connection failed")
logger.Error("API request timeout")
```

**Правильно:**
```go
logger.Info("database connection established")
logger.Error("api request timeout")
```

**Исправление:** Автоматическое - первая буква будет переведена в нижний регистр.

### 2. Правило английского языка
Сообщения логов должны быть на английском языке (без кириллицы).

**Неправильно:**
```go
logger.Error("Ошибка при обработке запроса")
logger.Warn("Пользователь не найден")
```

**Правильно:**
```go
logger.Error("error processing request")
logger.Warn("user not found in database")
```

### 3. Правило спецсимволов

Разрешены только следующие спецсимволы: `.`, `-`, `_`, `:` + дополнительные из конфигурации

**Неправильно:**
```go
logger.Info("invalid @character in message")
logger.Error("error #tag not allowed")
logger.Warn("price $100 invalid")
```

**Правильно:**
```go
logger.Info("user operation completed successfully")
logger.Error("process-timeout: retry in 30s")
logger.Warn("db.connection_pool.size=100")
```

**Конфигурация**: Добавьте дополнительные символы через `allowed-special-chars="!@#"` флаг

### 4. Правило чувствительных данных
Предотвращает логирование конфиденциальной информации вроде паролей, токенов, ключей API.

**Чувствительные ключевые слова:**
- `password`
- `token`
- `api_key`
- `secret`

**Неправильно:**
```go
logger.Info("user password is incorrect")
logger.Error("api_key validation failed")
```

**Правильно:**
```go
logger.Info("user authentication required")
logger.Error("authorization validation failed")
```

## Поддерживаемые логгеры

- Пакет `log/slog`
- `go.uber.org/zap` (структурированное логирование и Sugar API)

### Примеры

```go
import (
    "log/slog"
    "go.uber.org/zap"
)

// slog - структурированное логирование
logger := slog.Default()
logger.Info("database connection established")

// zap - структурированное логирование
logger, _ := zap.NewProduction()
logger.Info("database connection established")

// zap sugar - форматированное логирование
sugar := logger.Sugar()
sugar.Info("database connection established")
```

## Структура проекта

```
Go-Log-Linter/
├── go.mod                    # Определение модуля
├── go.sum                    # Хеши зависимостей
├── Makefile                  # Команды сборки
├── README.md                 # Этот файл
├── .golangci.yml             # Конфигурация golangci-lint
│
├── cmd/
│   └── loglint/
│       └── main.go           # CLI инструмент для прямого использования
│
├── analyzer/
│   ├── analyzer.go           # Главный анализатор с интеграцией конфигурации
│   ├── config.go             # Структура Config и управление конфигурацией
│   ├── config_test.go        # Тесты конфигурации
│   ├── logger_detector.go    # Детектор slog/zap логгеров
│   ├── logger_detector_test.go
│   ├── utils.go              # Утилиты (парсинг строк и т.д.)
│   ├── analyzer_test.go      # Тесты анализатора
│   ├── testdata_test.go      # Интеграционные тесты
│   └── rules/
│       ├── rules.go          # Реализация всех 4 правил + Checker
│       └── rules_test.go     # Unit тесты правил (27+ тест-кейсов)
│
├── plugin/
│   └── main.go               # Обёртка плагина для golangci-lint
│
└── testdata/                 # Интеграционные test fixtures
    ├── lowercase/
    │   ├── bad.go            # Нарушение нижнего регистра
    │   └── good.go           # Правильный нижний регистр
    ├── english/
    │   ├── bad.go            # Кириллица в логах
    │   └── good.go           # Только английский
    ├── special/
    │   ├── bad.go            # Запрещённые спецсимволы
    │   └── good.go           # Разрешённые спецсимволы
    ├── sensitive/
    │   ├── bad.go            # Конфиденциальные данные
    │   └── good.go           # Безопасное логирование
    ├── zap/
    │   ├── bad.go            # Нарушения при использовании Zap
    │   └── good.go           # Правильный Zap
    └── slog/
        ├── bad.go            # Нарушения при использовании slog
        └── good.go           # Правильный slog
```

### Ключевые компоненты

- **Config**: Управление конфигурацией линтера (5 параметров)
- **RuleConfig Interface**: Интерфейс для передачи конфигурации правилам
- **Checker**: Фабрика для создания правил с учётом конфигурации
- **Rules**: 4 основных правила проверки логов
- **Logger Detector**: Поддержка slog и zap логгеров

## Разработка

### Программный API

Используйте линтер в своём коде:

```go
package main

import (
    "github.com/Hirogava/Go-Log-Linter/analyzer"
    "golang.org/x/tools/go/analysis"
)

func main() {
    // Создать конфигурацию
    config := analyzer.DefaultConfig()
    config.AllowExclamation = true
    config.AllowedSpecialChars = "!@#"
    config.CustomSensitive = []string{"credit_card", "ssn"}
    
    // Использовать анализатор с конфигурацией
    a := analyzer.Analyzer
    // ... использовать анализатор
}
```

### Архитектура

1. **Config система**: Централизованная конфигурация со значениями по умолчанию
2. **RuleConfig интерфейс**: Позволяет правилам получать доступ к конфигурации
3. **Checker фабрика**: Создаёт правила с учётом конфигурации
4. **Logger Detector**: Автоматически обнаруживает slog и zap вызовы
5. **Analyzer**: Интегрирует всё вместе через golang.org/x/tools/go/analysis

### Сборка

**Использование Makefile (Linux/macOS):**
```bash
make build          # Собрать CLI инструмент для Linux
make build-plugin   # Собрать плагин для golangci-lint
make build-windows  # Собрать для Windows
make test           # Запустить тесты
make clean          # Очистить артефакты сборки
```

**Windows (PowerShell):**
```powershell
make build-windows  # или вручную:
$env:GOOS="windows"; $env:GOARCH="amd64"; go clean ./...; go build -o .\bin\loglint.exe .\cmd\loglint
```

**Linux:**
```bash
make build-linux    # или вручную:
GOOS=linux GOARCH=amd64 go build -o ./bin/loglint ./cmd/loglint
```

**macOS:**
```bash
make build-macos    # или вручную:
GOOS=darwin GOARCH=amd64 go build -o ./bin/loglint ./cmd/loglint
```

### Тестирование

Директория testdata содержит тест-кейсы для каждого правила. Каждая категория включает:
- `bad.go` - Примеры, которые должны вызвать ошибки анализатора
- `good.go` - Примеры, которые должны пройти успешно

Все тесты описаны через комментарии `// want` которые указывают ожидаемые ошибки.

### Расширение функциональности

Чтобы добавить новое правило:

1. Создайте структуру в `analyzer/rules/rules.go`:
```go
type CustomRule struct {
    config RuleConfig
}

func (r *CustomRule) Check(msg string) error {
    // Ваша логика проверки
    return nil
}

func (r *CustomRule) GetSuggestedFix(msg string) string {
    // Простые исправления
    return msg
}
```

2. Добавьте правило в `Checker.rules` список в `NewChecker()`

3. Добавьте тесты в `rules_test.go`

4. Добавьте тестовые примеры в `testdata/custom_rule/bad.go` и `good.go`

## Требования

- **Go**: 1.24 или выше
- **Для плагина golangci-lint**: Linux или macOS (Windows не поддерживает `-buildmode=plugin`)
- **golangci-lint**: v1.50+ (для интеграции плагина)

## Статус проекта

- 4 основных правила реализованы
- Система конфигурации реализована (5 параметров)
- Поддержка slog и zap логгеров
- Интеграция с golangci-lint
- 150+ тест-кейсов
- Документация (русский язык)
- SuggestedFix поддержка

## Версия

v1.0.0 - Добавлена система конфигурации и расширенная функциональность

## Автор

Hirogava