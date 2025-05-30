# In-Memory Key-Value Database

Простая in-memory key-value база данных, реализованная на Go.

## Возможности

- Поддержка базовых операций: SET, GET, DEL
- Многослойная архитектура (compute и storage слои)
- Логирование операций
- Тестовое покрытие
- Консольный интерфейс

## Установка

```bash
git clone https://github.com/your-username/in-memory-key-value.git
cd in-memory-key-value
go mod download
```

## Использование

Запуск приложения:
```bash
go run cmd/main.go
```

Примеры команд:
```
> SET weather_2_pm cold_moscow_weather
OK
> GET weather_2_pm
cold_moscow_weather
> DEL weather_2_pm
OK
```

## Структура проекта

```
.
├── cmd/
│   └── main.go           # Точка входа в приложение
├── internal/
│   ├── compute/
│   │   ├── parser.go     # Парсер запросов
│   │   └── parser_test.go
│   └── storage/
│       ├── engine.go     # In-memory движок
│       └── engine_test.go
├── pkg/
│   ├── compute/
│   │   └── interface.go  # Интерфейс compute слоя
│   └── storage/
│       └── interface.go  # Интерфейс storage слоя
└── go.mod
```

## Тестирование

Запуск тестов:
```bash
go test ./...
```

## Лицензия

MIT 