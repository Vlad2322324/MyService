[← Архитектура](architecture.md) · [Back to README](../README.md) · [Конфигурация →](configuration.md)

# API Reference

Описание всех endpoints, запросов, ответов и кодов ошибок.

## Base URL

```
http://localhost:8081/
```

## Аутентификация

Текущая версия не требует аутентификации. В future версиях планируется добавить OAuth2 или JWT.

## Endpoints

### Создать пользователя

**POST** `/user`

Создаёт нового пользователя.

**Request:**
```json
{
  "name": "Alice Johnson"
}
```

**Параметры:**
- `name` (string, required) — имя пользователя (1-100 символов)

**Response (200 OK):**
```json
{
  "id": 1,
  "name": "Alice Johnson",
  "created_at": "2026-06-09T10:00:00Z"
}
```

**Errors:**
- `400 Bad Request` — пустое имя или тело запроса
- `500 Internal Server Error` — ошибка БД

**Examples:**
```bash
curl -X POST http://localhost:8081/user \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice Johnson"}'
```

---

### Получить пользователя

**GET** `/user/{id}`

Получает пользователя по ID.

**Параметры пути:**
- `id` (integer, required) — ID пользователя (≥1)

**Response (200 OK):**
```json
{
  "id": 1,
  "name": "Alice Johnson",
  "created_at": "2026-06-09T10:00:00Z"
}
```

**Errors:**
- `404 Not Found` — пользователь не существует
- `500 Internal Server Error` — ошибка БД

**Examples:**
```bash
curl http://localhost:8081/user/1
```

---

### Обновить пользователя

**PUT** `/user/{id}`

Обновляет имя пользователя.

**Параметры пути:**
- `id` (integer, required) — ID пользователя (≥1)

**Request:**
```json
{
  "name": "Alice Smith"
}
```

**Параметры:**
- `name` (string, required) — новое имя (1-100 символов)

**Response (200 OK):**
```json
{
  "id": 1,
  "name": "Alice Smith",
  "created_at": "2026-06-09T10:00:00Z"
}
```

**Errors:**
- `400 Bad Request` — пустое имя или тело запроса
- `404 Not Found` — пользователь не существует
- `500 Internal Server Error` — ошибка БД

**Examples:**
```bash
curl -X PUT http://localhost:8081/user/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice Smith"}'
```

---

### Удалить пользователя

**DELETE** `/user/{id}`

Удаляет пользователя.

**Параметры пути:**
- `id` (integer, required) — ID пользователя (≥1)

**Response (204 No Content):**
```
(пустой body)
```

**Errors:**
- `404 Not Found` — пользователь не существует
- `500 Internal Server Error` — ошибка БД

**Examples:**
```bash
curl -X DELETE http://localhost:8081/user/1
```

---

## Коды ответов

| Код | Описание |
|-----|---------|
| 200 | OK — успешный запрос |
| 204 | No Content — успешное удаление |
| 400 | Bad Request — ошибка валидации входных данных |
| 404 | Not Found — ресурс не найден |
| 500 | Internal Server Error — ошибка сервера |

## Формат ошибок

При ошибках возвращается JSON:

```json
{
  "code": 400,
  "message": "Empty User Name"
}
```

**Поля:**
- `code` (integer) — HTTP статус код
- `message` (string) — описание ошибки

## Headers

### Request

- `Content-Type: application/json` — обязателен для POST/PUT
- `X-Request-ID` (optional) — custom request ID (если не передан, сервер генерирует)

### Response

- `Content-Type: application/json` — всегда присутствует
- `X-Request-ID` — echo от request, используется для трассировки

## Примеры сценариев

### Сценарий 1: Создание и получение пользователя

```bash
# 1. Создать пользователя
curl -X POST http://localhost:8081/user \
  -H "Content-Type: application/json" \
  -d '{"name":"Bob"}'
# Ответ: {"id":1,"name":"Bob","created_at":"2026-06-09T10:00:00Z"}

# 2. Получить пользователя (используем id=1)
curl http://localhost:8081/user/1
# Ответ: {"id":1,"name":"Bob","created_at":"2026-06-09T10:00:00Z"}
```

### Сценарий 2: Обновление и удаление

```bash
# 1. Обновить
curl -X PUT http://localhost:8081/user/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Bob Updated"}'
# Ответ: {"id":1,"name":"Bob Updated","created_at":"2026-06-09T10:00:00Z"}

# 2. Удалить
curl -X DELETE http://localhost:8081/user/1
# Ответ: 204 No Content
```

## See Also

- [Архитектура](architecture.md) — структура кода
- [Конфигурация](configuration.md) — переменные окружения
- [Разработка](development.md) — для разработчиков

