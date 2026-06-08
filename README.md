# MyService

> REST API сервис для управления пользователями на Go

Простой, масштабируемый REST-сервис на Go с использованием Echo, PostgreSQL и OpenAPI. Проект демонстрирует лучшие практики: структурированное логирование, graceful shutdown, миграции БД и CI/CD.

## Быстрый старт

**Требования:** Go 1.23+, PostgreSQL, oapi-codegen

```bash
# Клонировать репозиторий
git clone <repo-url>
cd MyService

# Установить зависимости
go mod download

# Установить oapi-codegen
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

# Скопировать .env.example в .env и заполнить DATABASE_URL
cp .env.example .env

# Собрать и запустить
.\scripts\build.ps1
.\scripts\run.ps1
```

## Ключевые возможности

- **OpenAPI-first**: спецификация → код через oapi-codegen
- **Структурированное логирование**: zap с поддержкой уровней (debug/info/warn/error)
- **Request tracking**: X-Request-ID middleware для отслеживания запросов
- **Миграции БД**: golang-migrate + SQL-скрипты в `migrations/`
- **Graceful shutdown**: корректное завершение при SIGINT/SIGTERM
- **CI/CD**: GitHub Actions (codegen, vet, test, build)

## Примеры

```bash
# Создать пользователя
curl -X POST http://localhost:8081/user \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice"}'

# Получить пользователя
curl http://localhost:8081/user/1

# Обновить пользователя
curl -X PUT http://localhost:8081/user/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice Updated"}'

# Удалить пользователя
curl -X DELETE http://localhost:8081/user/1
```

---

## Документация

| Раздел | Описание |
|--------|---------|
| [Начало работы](docs/getting-started.md) | Установка, требования, первый запуск |
| [Архитектура](docs/architecture.md) | Структура проекта, компоненты, взаимодействие |
| [API Reference](docs/api-reference.md) | Endpoints, запросы, ответы, коды ошибок |
| [Конфигурация](docs/configuration.md) | Переменные окружения, файлы конфигурации |
| [Разработка](docs/development.md) | Логирование, миграции, middleware, тестирование |

## Стек технологий

- **Go 1.23** — язык программирования
- **Echo v4** — HTTP framework
- **pgx v5** — PostgreSQL драйвер
- **zap** — структурированное логирование
- **oapi-codegen** — генерация API из OpenAPI
- **golang-migrate** — управление миграциями

## Запуск тестов

```bash
go test ./... -v
go vet ./...
```

## Лицензия

MIT — см. [LICENSE](LICENSE)

## Больше информации

- 📖 [GitHub](https://github.com/yourusername/MyService)
- 🐛 [Issues](https://github.com/yourusername/MyService/issues)
- 💬 [Discussions](https://github.com/yourusername/MyService/discussions)



