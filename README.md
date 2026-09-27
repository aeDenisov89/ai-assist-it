# Todo API Service

REST API для управления пользователями и задачами, написанный на Go.
Проект демонстрирует профессиональное оформление REST API: документация
Swagger, структурированный README, валидация данных и понятная структура кода.

## Возможности

- CRUD для пользователей (`users`)
- CRUD для задач (`todos`)
- Валидация входных данных
- Автоматически сгенерированная Swagger-документация
- Swagger UI для интерактивного тестирования API

## Требования

- Go 1.21 или выше
- Утилита [`swag`](https://github.com/swaggo/swag) для генерации документации
  
## Установка и запуск

1. Клонируйте репозиторий:

   ```bash
   git clone https://gitverse.ru/netology/ai-assist-it.git
   cd ai-assist-it
   ```

2. Установите зависимости:

   ```bash
   go mod download
   ```

3. Установите `swag` CLI (если ещё не установлен):

   ```bash
   go install github.com/swaggo/swag/cmd/swag@latest
   ```

4. Сгенерируйте Swagger-документацию:

   ```bash
   swag init -g main.go -o docs
   ```

5. Запустите сервер:

   ```bash
   go run main.go
   ```

6. Откройте Swagger UI в браузере:

   http://localhost:8080/swagger/

## Зависимости

| Пакет | Назначение |
|-------|-----------|
| `github.com/go-chi/chi/v5` | HTTP-роутер |
| `github.com/swaggo/http-swagger/v2` | Swagger UI для chi |
| `github.com/swaggo/swag` | Генерация Swagger-документации |

## Описание API

### Группы endpoints

| Группа | Описание |
|--------|----------|
| `users` | Управление пользователями |
| `todos` | Управление задачами |

### Список endpoints

**Users**

| Метод  | Путь                  | Описание                 |
|--------|-----------------------|--------------------------|
| GET    | `/api/v1/users`       | Список пользователей     |
| POST   | `/api/v1/users`       | Создать пользователя     |
| GET    | `/api/v1/users/{id}`  | Получить пользователя    |
| PUT    | `/api/v1/users/{id}`  | Обновить пользователя    |
| DELETE | `/api/v1/users/{id}`  | Удалить пользователя     |

**Todos**

| Метод  | Путь                  | Описание                 |
|--------|-----------------------|--------------------------|
| GET    | `/api/v1/todos`       | Список задач             |
| POST   | `/api/v1/todos`       | Создать задачу           |
| GET    | `/api/v1/todos/{id}`  | Получить задачу          |
| PUT    | `/api/v1/todos/{id}`  | Обновить задачу          |
| DELETE | `/api/v1/todos/{id}`  | Удалить задачу           |

## Примеры запросов и ответов

Все примеры предполагают, что сервер запущен на `http://localhost:8080`.

### Список пользователей

```bash
curl -i http://localhost:8080/api/v1/users
```

```
HTTP/1.1 200 OK
Content-Type: application/json

{"success":true,"data":null}
```

### Создание пользователя

`user.json`:

```json
{"email":"ivan@example.com","name":"Ivan"}
```

```bash
curl -i -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d @user.json
```

```
HTTP/1.1 201 Created
Content-Type: application/json

{"success":true,"data":{"id":2,"email":"ivan@example.com","name":"Ivan","created_at":"2026-09-27T11:40:35.5897667+03:00"}}
```

### Получение пользователя по ID

```bash
curl -i http://localhost:8080/api/v1/users/2
```

```
HTTP/1.1 200 OK
Content-Type: application/json

{"success":true,"data":{"id":2,"email":"ivan@example.com","name":"Ivan","created_at":"2026-09-27T11:40:35.5897667+03:00"}}
```

### Обновление пользователя

`user_update.json`:

```json
{"name":"Ivan Petrov"}
```

```bash
curl -i -X PUT http://localhost:8080/api/v1/users/2 \
  -H "Content-Type: application/json" \
  -d @user_update.json
```

```
HTTP/1.1 200 OK
Content-Type: application/json

{"success":true,"data":{"id":2,"email":"ivan@example.com","name":"Ivan Petrov","created_at":"2026-09-27T11:40:35.5897667+03:00"}}
```

### Удаление пользователя

```bash
curl -i -X DELETE http://localhost:8080/api/v1/users/2
```

```
HTTP/1.1 204 No Content
```

### Создание задачи

`todo.json`:

```json
{"user_id":1,"title":"Buy milk","description":"2 liters","done":false}
```

```bash
curl -i -X POST http://localhost:8080/api/v1/todos \
  -H "Content-Type: application/json" \
  -d @todo.json
```

```
HTTP/1.1 201 Created
Content-Type: application/json

{"success":true,"data":{"id":1,"user_id":1,"title":"Buy milk","description":"2 liters","done":false,"created_at":"2026-09-27T11:43:12.7544587+03:00","updated_at":"2026-09-27T11:43:12.7544587+03:00"}}
```

### Список задач

```bash
curl -i http://localhost:8080/api/v1/todos
```

```
HTTP/1.1 200 OK
Content-Type: application/json

{"success":true,"data":[{"id":1,"user_id":1,"title":"Buy milk","description":"2 liters","done":false,"created_at":"2026-09-27T11:43:12.7544587+03:00","updated_at":"2026-09-27T11:43:12.7544587+03:00"}]}
```

### Обработка ошибок

Запрос несуществующего пользователя:

```bash
curl -i http://localhost:8080/api/v1/users/999
```

```
HTTP/1.1 404 Not Found
Content-Type: text/plain; charset=utf-8

Пользователь не найден
```

## Структура проекта

```
ai-assist-it/
├── main.go                     # Точка входа, роутинг, общие аннотации Swagger
├── go.mod                      # Модуль и зависимости
├── go.sum                      # Контрольные суммы зависимостей
├── README.md                   # Документация проекта
├── docs/                       # Сгенерированная Swagger-документация
│   ├── docs.go                 # Go-пакет со встроенным Swagger JSON
│   ├── swagger.json            # OpenAPI-спецификация (JSON)
│   └── swagger.yaml            # OpenAPI-спецификация (YAML)
└── internal/
    ├── handlers/               # HTTP-обработчики
    │   ├── todo.go             # Обработчики задач
    │   └── user.go             # Обработчики пользователей
    └── models/                 # Модели данных
        ├── todo.go             # Модели задач
        └── user.go             # Модели пользователей
```

### Чек-лист структуры

- [x] Точка входа `main.go` в корне проекта
- [x] Разделение на `handlers` и `models` в `internal/`
- [x] Модульная система Go (`go.mod`, `go.sum`)
- [x] Сгенерированная Swagger-документация в `docs/`
- [x] README с описанием, установкой и примерами

## Лицензия

Проект распространяется под лицензией MIT.
Подробности — в файле [LICENSE](LICENSE) (при наличии) или на
[https://opensource.org/licenses/MIT](https://opensource.org/licenses/MIT).

## Контакты

- **Автор:** Андрей
- **Email:** support@example.com
- **Репозиторий:** https://gitverse.ru/netology/ai-assist-it