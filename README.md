# MyService

Короткое описание

MyService — простой REST-сервис на Go для работы с сущностью "пользователь" (Users). Проект использует Echo для HTTP-сервера, pgx для подключения к PostgreSQL и сгенерированные хендлеры из OpenAPI спецификации в `openapi/openapi.yaml`.

Чеклист (что в этом README):

- Описание проекта
- Требования и зависимости
- Быстрый старт (сборка и запуск)
- Конфигурация окружения
- Работа с базой данных / миграции
- Генерация серверных хендлеров из OpenAPI
- Примеры запросов
- Структура проекта

Требования

- Go 1.20+ (или версия, указанная в go.mod)
- make (Makefile) — для удобных команд (на Windows можно использовать mingw/msys, WSL или просто выполнять команды вручную в PowerShell)
- PostgreSQL
- oapi-codegen — требуется для генерации API (Makefile вызывает `oapi-codegen`)

Установка oapi-codegen (если ещё не установлен)

PowerShell / Windows (scoop/choco) пример:

```powershell
# через scoop
scoop install oapi-codegen

# через go (если не установлен бинарник):
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
```

Быстрый старт

1) Создайте файл с переменными окружения `.env` в корне проекта (или установите переменную `DATABASE_URL`). Пример содержимого:

```
DATABASE_URL=postgres://postgres:password@localhost:5432/mydb?sslmode=disable
```

2) Собрать и запустить приложение (Makefile автоматически вызовет `oapi-codegen`):

PowerShell:

```powershell
# Сборка
make build

# Запуск
make run
```

Альтернатива: запустить напрямую через go

```powershell
# сгенерировать код (если нужно)
make codegen

# собрать и запустить без Makefile
go build -o bin/myapp.exe ./cmd/app/main.go
.
bin\myapp.exe

# или
go run ./cmd/app/main.go
```

Конфигурация базы данных

Приложение ожидает переменную окружения `DATABASE_URL` в формате PostgreSQL URI, например:

postgres://user:password@localhost:5432/dbname?sslmode=disable

Если БД пуста, можно создать таблицу вручную (SQL):

```sql
CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

В текущем коде есть метод `InitUserTable` в `internal/repositories`, который выполняет создание таблицы — его можно вызвать из `main` или из отдельного инструмента миграции.

Генерация OpenAPI-кода

Спецификация находится в `openapi/openapi.yaml`. Makefile содержит цель `codegen`, которая использует `openapi/oapi-codegen.yaml` + `oapi-codegen` для генерации серверной обвязки в `internal/api/generated`.

Примеры запросов

После запуска сервис слушает порт 8080 (в `cmd/app/main.go` указано `:8080`). Примеры (PowerShell / curl):

Создать пользователя:

```powershell
curl -X POST http://localhost:8080/user -H "Content-Type: application/json" -d '{"name":"John Doe"}'
```

Получить пользователя по id:

```powershell
curl http://localhost:8080/user/1
```

Обновить пользователя:

```powershell
curl -X PUT http://localhost:8080/user/1 -H "Content-Type: application/json" -d '{"name":"Updated Name"}'
```

Удалить пользователя:

```powershell
curl -X DELETE http://localhost:8080/user/1
```

Структура проекта (важные директории)

- `cmd/app` — точка входа `main.go`.
- `internal/api/generated` — сгенерированные хендлеры и типы (oapi-codegen).
- `internal/database` — подключение к БД (`postgres.go`).
- `internal/repositories` — слой доступа к данным (`user_repository.go`).
- `internal/services` — бизнес-логика (`user_service.go`).
- `internal/handlers` — реализация HTTP-хендлеров поверх сгенерированного API.
- `openapi` — спецификация OpenAPI и конфиг для oapi-codegen.
- `migrations` — место для миграций (пока может быть пустым).
- `pkg/logger` — пакет логирования.
- `Makefile` — удобные команды для сборки, запуска, генерации кода и миграций.
- `.env` — файл для переменных окружения (не в репозитории, должен быть создан локально).



