[← API Reference](api-reference.md) · [Back to README](../README.md) · [Разработка →](development.md)

# Конфигурация

Переменные окружения, файлы конфигурации и настройки.

## Переменные окружения

Все переменные загружаются из файла `.env` (используется godotenv).

### DATABASE_URL (обязательна)

**Описание:** Строка подключения к PostgreSQL

**Формат:**
```
postgres://[username]:[password]@[host]:[port]/[database][?sslmode=disable]
```

**Примеры:**
```
postgres://postgres:password@localhost:5432/mydb?sslmode=disable
postgres://user:pass@db.example.com:5432/prod_db?sslmode=require
```

**По умолчанию:** не задана (обязательна)

**Примечание:** В development используйте `sslmode=disable`. В production используйте `sslmode=require` или `sslmode=verify-full`.

---

### ENV (опционально)

**Описание:** Окружение для логирования и конфигурации

**Допустимые значения:**
- `development` — развёрнутые логи, debug-сообщения (default)
- `production` — JSON-логи, только важные сообщения

**По умолчанию:** `development`

**Примечание:** Влияет на формат логов через zap-конфиг.

---

### LOG_LEVEL (опционально)

**Описание:** Уровень логирования

**Допустимые значения:**
- `debug` — все сообщения (verbose)
- `info` — информационные и выше (default в dev)
- `warn` — предупреждения и выше
- `error` — только ошибки

**По умолчанию:** `info` (production), `debug` (development)

**Примечание:** Устанавливает минимальный уровень логирования.

---

### PORT (опционально)

**Описание:** Порт, на котором слушает HTTP сервер

**Допустимые значения:** 1024-65535

**По умолчанию:** `8081`

**Примечание:** Убедитесь, что порт не занят другим приложением.

---

## Файл .env

### Пример конфигурации (development)

```bash
# Database connection
DATABASE_URL=postgres://postgres:password@localhost:5432/mydb?sslmode=disable

# Environment and logging
ENV=development
LOG_LEVEL=debug

# Server
PORT=8081
```

### Пример конфигурации (production)

```bash
# Database connection (с SSL)
DATABASE_URL=postgres://app_user:secure_password@db.prod.example.com:5432/mydb?sslmode=require

# Environment and logging
ENV=production
LOG_LEVEL=info

# Server
PORT=8000
```

### Создание .env файла

1. Скопируйте `.env.example`:
   ```bash
   cp .env.example .env
   ```

2. Отредактируйте `.env`, установив ваши значения

3. **Важно:** не коммитьте `.env` в git (он в `.gitignore`)

---

## PostgreSQL конфигурация

### Создание БД и пользователя

```sql
-- Создать БД
CREATE DATABASE mydb;

-- Создать пользователя (опционально)
CREATE USER app_user WITH PASSWORD 'secure_password';

-- Дать права
GRANT ALL PRIVILEGES ON DATABASE mydb TO app_user;
```

### Проверка подключения

```bash
# Локально
psql -U postgres -d mydb -h localhost

# С параметрами
psql -U app_user -d mydb -h db.prod.example.com -p 5432
```

---

## Загрузка конфигурации

При старте приложение:

1. Загружает `.env` файл (godotenv.Load())
2. Инициализирует логгер с параметрами `ENV` и `LOG_LEVEL`
3. Подключается к БД используя `DATABASE_URL`
4. Запускает сервер на порту из `PORT`

### Порядок приоритета

1. Переменные окружения системы (highest priority)
2. Значения из `.env` файла
3. Значения по умолчанию в коде (lowest priority)

---

## Пример: изменение уровня логирования

Без перезапуска приложения можно изменить уровень логирования через переменную окружения:

**Linux/macOS:**
```bash
export LOG_LEVEL=debug
go run ./cmd/app
```

**Windows (PowerShell):**
```powershell
$env:LOG_LEVEL = "debug"
go run ./cmd/app
```

При следующем запуске приложение прочитает новое значение.

---

## Проверка текущей конфигурации

При старте приложения логирует:

```
2026-06-09T10:00:00.000+0300    INFO    logger initialized (ENV=development)
```

Это подтверждает, что конфигурация загружена. Проверьте логи для отладки.

---

## See Also

- [Начало работы](getting-started.md) — установка и первый запуск
- [Разработка](development.md) — информация о логировании

