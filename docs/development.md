[← Конфигурация](configuration.md) · [Back to README](../README.md)

# Разработка

Информация для разработчиков: логирование, миграции, middleware, тестирование и расширение.

## Локальная разработка

### Запуск с отладкой

```bash
# Development режим с debug логами
export LOG_LEVEL=debug
export ENV=development
go run ./cmd/app
```

### Горячая перезагрузка (optional)

Используйте `air` для автоматической перезагрузки при изменении кода:

```bash
go install github.com/cosmtrek/air@latest
air
```

---

## Логирование

Проект использует **zap** для структурированного логирования.

### Использование в коде

```go
import "myservice/pkg/logger"

// Info уровень
logger.Log.Infof("user created: id=%d name=%s", user.ID, user.Name)

// Error уровень
logger.Log.Errorf("failed to fetch user: %v", err)

// Debug (видна только с LOG_LEVEL=debug)
logger.Log.Debugf("query result: %v", rows)
```

### Уровни логирования

| Уровень | Использование |
|---------|--------------|
| `debug` | Детальная информация для отладки (значения переменных, промежуточные шаги) |
| `info` | Информационные события (запуск, создание, успех) |
| `warn` | Предупреждения (deprecated, замещено) |
| `error` | Ошибки (операция не выполнена, но сервис работает) |

### Форматы логов

**Development (текстовый):**
```
2026-06-09T10:00:00.000+0300    INFO    myservice/user_service.go:22    CreateUser called with name=Alice
```

**Production (JSON):**
```json
{"level":"info","ts":1686306000.000,"caller":"myservice/user_service.go:22","msg":"CreateUser called with name=Alice"}
```

---

## Миграции БД

### Структура миграций

Миграции находятся в `migrations/`. Каждая миграция состоит из пары файлов:

```
migrations/
├── 0001_create_users.up.sql      # Применить
└── 0001_create_users.down.sql    # Откатить
```

### Добавление новой миграции

1. Создайте файлы с номером миграции (incrementing):
   ```bash
   # Далее создавайте:
   migrations/0002_add_email_to_users.up.sql
   migrations/0002_add_email_to_users.down.sql
   ```

2. `0002_add_email_to_users.up.sql`:
   ```sql
   ALTER TABLE users ADD COLUMN email VARCHAR(100);
   ```

3. `0002_add_email_to_users.down.sql`:
   ```sql
   ALTER TABLE users DROP COLUMN email;
   ```

4. Запустите миграции:
   ```bash
   .\scripts\migrate.ps1
   ```

### Правила для миграций

- **Idempotent:** миграции должны быть безопасны при повторном запуске
- **Реверсивные:** каждому `.up.sql` должен соответствовать `.down.sql`
- **Транзакции:** используйте явные транзакции если нужна атомарность
- **Тестирование:** тестируйте как `.up` так и `.down`

---

## Middleware

### Request ID Middleware

Находится в `internal/middleware/request_id.go`.

Функция:
- Генерирует или получает `X-Request-ID` из заголовка запроса
- Сохраняет в контексте (`c.Set("request_id", reqID)`)
- Возвращает в ответе (`X-Request-ID` заголовок)

Используется для трассировки запросов в логах.

### Добавление нового middleware

```go
// internal/middleware/custom.go
package middleware

import "github.com/labstack/echo/v4"

func CustomMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        // Что-то до handler
        err := next(c)
        // Что-то после handler
        return err
    }
}
```

Регистрация в `cmd/app/main.go`:

```go
e := echo.New()
e.Use(middleware.CustomMiddleware)
```

---

## Тестирование

### Unit-тесты

Пример теста для сервиса:

```go
// internal/services/user_service_test.go
package services

import (
    "context"
    "testing"
)

func TestCreateUser(t *testing.T) {
    // Arrange
    repo := &mockUserRepository{}
    service := NewUserService(repo)
    
    // Act
    user, err := service.CreateUser(context.Background(), "Alice")
    
    // Assert
    if err != nil {
        t.Fatalf("expected no error, got %v", err)
    }
    if user.Name != "Alice" {
        t.Fatalf("expected name=Alice, got %s", user.Name)
    }
}
```

### Запуск тестов

```bash
# Все тесты
go test ./...

# С coverage
go test -cover ./...

# Verbose
go test -v ./...
```

### Интеграционные тесты

Для интеграционных тестов используйте тестовую БД:

```bash
# Запустить PostgreSQL в Docker для тестов
docker run -d \
  -e POSTGRES_PASSWORD=test \
  -e POSTGRES_DB=test_db \
  -p 5432:5432 \
  postgres:15
```

---

## Добавление новой функции

### 1. Обновите OpenAPI спецификацию

Отредактируйте `openapi/openapi.yaml` и добавьте новый endpoint:

```yaml
paths:
  /user/{id}/profile:
    get:
      operationId: GetUserProfile
      parameters:
        - name: id
          in: path
          required: true
          schema:
            type: integer
      responses:
        '200':
          description: User profile
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/UserProfile'
```

### 2. Генерируйте код

```bash
.\scripts\codegen.ps1
```

Это обновит `internal/api/generated/api.gen.go` с новым методом.

### 3. Реализуйте handler

```go
// internal/handlers/user_handler/user_handler.go
func (h *UserHandler) GetUserProfile(
    ctx context.Context,
    request api.GetUserProfileRequestObject,
) (api.GetUserProfileResponseObject, error) {
    profile, err := h.us.GetUserProfile(ctx, request.Id)
    if err != nil {
        logger.Log.Errorf("GetUserProfile error: %v", err)
        return api.GetUserProfile500JSONResponse{...}, nil
    }
    return api.GetUserProfile200JSONResponse{...}, nil
}
```

### 4. Реализуйте сервис

```go
// internal/services/user_service.go
func (us *UserService) GetUserProfile(ctx context.Context, id int) (..., error) {
    // Бизнес-логика
    return ...
}
```

### 5. Реализуйте репо-метод

```go
// internal/repositories/user_repository.go
func (ur *UserRepo) GetUserProfile(ctx context.Context, id int) (..., error) {
    // SQL запрос
    return ...
}
```

### 6. Добавьте миграцию (если нужна схема)

Создайте файлы в `migrations/` и запустите `.\scripts\migrate.ps1`.

### 7. Протестируйте

```bash
go test ./... -v
curl http://localhost:8081/user/1/profile
```

---

## Code Style

- Используйте `gofmt` для форматирования: `go fmt ./...`
- Следуйте Go conventions: CamelCase для exported, lowercase для unexported
- Добавляйте комментарии для exported функций/типов
- Обрабатывайте ошибки явно (не игнорируйте `err`)

---

## Коммиты

Используйте Conventional Commits:

```
feat(user): add email field to users table
fix(api): handle nil pointers in GetUser handler
docs(readme): update installation steps
chore(deps): upgrade golang.org/x/sys to v0.44.0
```

---

## See Also

- [Архитектура](architecture.md) — структура проекта
- [API Reference](api-reference.md) — endpoints
- [Конфигурация](configuration.md) — переменные окружения

