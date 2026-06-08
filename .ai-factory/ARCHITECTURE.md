# ARCHITECTURE — MyService

Краткое описание архитектуры проекта MyService.

Стек:
- Go (модули)
- Echo (HTTP framework)
- pgx (Postgres driver)
- OpenAPI (spec: `openapi/openapi.yaml`) + oapi-codegen

Структура компонентов:
- `cmd/app` — точка входа, init сервисов, регистрация хендлеров, middleware.
- `internal/database` — логика подключения к БД.
- `internal/repositories` — слой доступа к данным (UserRepo).
- `internal/services` — бизнес-логика (UserService).
- `internal/handlers` — реализации хендлеров, связывающие OpenAPI generated -> сервисы.
- `internal/api/generated` — сгенерированный код из OpenAPI (oapi-codegen).
- `migrations/` — SQL миграции (golang-migrate).
- `pkg/logger` — централизованный логгер (zap wrapper).
- `internal/middleware` — middleware (request-id и др.).

Принципы:
- Генерация кода из OpenAPI должна быть воспроизводимой (CI запускает codegen).
- Миграции — source of truth для схемы БД.
- Логи — структурированные (JSON в prod), с уровнем и request_id.

