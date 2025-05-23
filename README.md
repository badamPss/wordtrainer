# Word Trainer

Бэкенд-сервис для изучения иностранных слов с помощью карточек. Предоставляет REST API для создания категорий слов, добавления карточек и отслеживания прогресса обучения. Может быть использован как основа для создания веб-приложения или мобильного приложения.

## Функциональность

- REST API для:
  - Регистрации и авторизации пользователей
  - Создания категорий для группировки слов
  - Добавления карточек со словами и их переводами
- Swagger документация API
- Готовность к интеграции с фронтендом

## Технологии

- Go 1.21
- PostgreSQL
- Docker & Docker Compose
- Echo Framework
- GORM
- JWT для аутентификации

## Требования

- Docker
- Docker Compose
- Make (опционально)

## Установка и запуск

1. Клонируйте репозиторий:
```bash
git clone https://github.com/yourusername/wordtrainer.git
cd wordtrainer
```

2. Запустите проект:
```bash
make up
```

Или без Make:
```bash
docker-compose up --build -d
```

## Использование

### API Endpoints

#### Регистрация
```bash
curl --location 'http://localhost:8080/register' \
--header 'Content-Type: application/json' \
--data '{
    "username": "user1",
    "password_hash": "password123"
}'
```

#### Авторизация
```bash
curl --location 'http://localhost:8080/login' \
--header 'Content-Type: application/json' \
--data '{
    "username": "user1",
    "password_hash": "password123"
}'
```

#### Создание категории
```bash
curl --location 'http://localhost:8080/categories' \
--header 'Authorization: Bearer YOUR_TOKEN' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Еда"
}'
```

#### Получение списка категорий
```bash
curl --location 'http://localhost:8080/categories' \
--header 'Authorization: Bearer YOUR_TOKEN'
```

#### Создание карточки
```bash
curl --location 'http://localhost:8080/cards' \
--header 'Authorization: Bearer YOUR_TOKEN' \
--header 'Content-Type: application/json' \
--data '{
    "category_id": 1,
    "word": "apple",
    "translation": "яблоко"
}'
```

#### Получение списка карточек
```bash
curl --location 'http://localhost:8080/cards' \
--header 'Authorization: Bearer YOUR_TOKEN'
```

#### Получение карточки по ID
```bash
curl --location 'http://localhost:8080/cards/1' \
--header 'Authorization: Bearer YOUR_TOKEN'
```

### Swagger UI

Документация API доступна по адресу: http://localhost:8080/swagger/index.html

Полная документация API в формате OpenAPI (Swagger) находится в директории `docs/`:
- `docs/swagger.json` - документация в формате JSON
- `docs/swagger.yaml` - документация в формате YAML

## Управление проектом

- `make up` - поднять проект
- `make down` - остановить проект
- `make rebuild` - пересоздать проект (удалить все данные и поднять заново)
- `make logs` - посмотреть логи
- `make clean` - очистить все контейнеры и образы
- `make build` - собрать проект
- `make migrate` - применить миграции
- `make migrate-down` - откатить миграции
- `make test` - запустить тесты
- `make lint` - запустить линтер
- `make swagger` - сгенерировать Swagger документацию

## Структура проекта

```
.
├── cmd/
│   └── main.go           # Точка входа приложения
├── internal/
│   ├── config/          # Конфигурация приложения
│   ├── db/              # Работа с базой данных
│   ├── handlers/        # HTTP обработчики
│   ├── middleware/      # Middleware компоненты
│   ├── migrations/      # Миграции базы данных
│   ├── models/          # Модели данных
│   └── utils/           # Вспомогательные функции
├── docs/                # Swagger документация
│   ├── docs.go         # Сгенерированный код Swagger
│   ├── swagger.json    # OpenAPI спецификация (JSON)
│   └── swagger.yaml    # OpenAPI спецификация (YAML)
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
└── Makefile
```

## Лицензия

MIT