[Back to README](../README.md) · [Архитектура →](architecture.md)

# Начало работы

Установка, требования, первый запуск MyService.

## Требования

- **Go 1.23+** — [установить](https://golang.org/doc/install)
- **PostgreSQL 12+** — [установить](https://www.postgresql.org/download/)
- **Git** — для клонирования репозитория
- **oapi-codegen** — установится автоматически через `go install`
- **golang-migrate** (опционально) — для управления миграциями

### Windows: рекомендуемый способ установки

```powershell
# Go
choco install golang
# или через scoop
scoop install go

# PostgreSQL (выберите один)
choco install postgresql
# или скачайте с https://www.postgresql.org/download/windows/

# Git
choco install git
```

## Установка

1. **Клонируйте репозиторий:**
   ```bash
   git clone https://github.com/yourusername/MyService.git
   cd MyService
   ```

2. **Загрузите зависимости:**
   ```bash
   go mod download
   ```

3. **Установите oapi-codegen:**
   ```bash
   go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
   ```

4. **Создайте файл `.env`:**
   ```bash
   cp .env.example .env
   ```
   Отредактируйте `.env`, установив параметры подключения к PostgreSQL:
   ```
   DATABASE_URL=postgres://postgres:password@localhost:5432/mydb?sslmode=disable
   ENV=development
   LOG_LEVEL=debug
   PORT=8081
   ```

## Первый запуск

### Windows (PowerShell)

```powershell
# Скачайте и соберите приложение
.\scripts\build.ps1

# Запустите приложение с логами в logs/app.log
.\scripts\run.ps1
```

Приложение запустится на `http://localhost:8081` (по умолчанию).

### Linux/macOS

```bash
# Скачайте и соберите приложение
./scripts/build.ps1  # или используйте make build

# Запустите приложение
go run ./cmd/app/main.go
```

## Проверка установки

После запуска приложение должно напечатать логи типа:

```
2026-06-09T10:00:00.000+0300    INFO    logger initialized (ENV=development)
2026-06-09T10:00:00.001+0300    INFO    initializing users table if not exists
```

Протестируйте API:

```bash
curl -X POST http://localhost:8081/user \
  -H "Content-Type: application/json" \
  -d '{"name":"Test User"}'
```

Ожидаемый ответ:
```json
{"id":1,"name":"Test User","created_at":"2026-06-09T10:00:00Z"}
```

## Миграции БД

Если вы хотите использовать `golang-migrate`:

1. **Установите утилиту:**
   ```bash
   go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
   ```

2. **Запустите миграции:**
   ```powershell
   .\scripts\migrate.ps1
   ```

Текущая миграция создаёт таблицу `users`. Дополнительные миграции добавляйте в `migrations/` с названиями вида `000N_description.{up,down}.sql`.

## Работа с логами

Логи выводятся в терминал и опционально записываются в `logs/app.log` (используйте `.\scripts\run.ps1` для записи в файл).

**Уровни логирования** (управляйте через `LOG_LEVEL` в `.env`):
- `debug` — все сообщения
- `info` — информационные и выше
- `warn` — предупреждения и выше
- `error` — только ошибки

## Следующие шаги

- 📖 [Архитектура](architecture.md) — узнайте, как организован код
- 🔧 [Конфигурация](configuration.md) — все доступные переменные окружения
- 📡 [API Reference](api-reference.md) — список всех endpoints
- 💻 [Разработка](development.md) — как добавлять новые функции

## See Also

- [API Reference](api-reference.md) — список endpoints
- [Конфигурация](configuration.md) — переменные окружения
- [Разработка](development.md) — для разработчиков

