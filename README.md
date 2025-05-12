# Word Trainer

Бэкенд-сервис для изучения иностранных слов с помощью карточек. Предоставляет REST API для создания категорий слов, добавления карточек и отслеживания прогресса обучения. Может быть использован как основа для создания веб-приложения или мобильного приложения.

## Функциональность

- REST API для:
  - Регистрации и авторизации пользователей
  - Создания категорий для группировки слов
  - Добавления карточек со словами и их переводами
  - Отслеживания попыток изучения слов
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

### Swagger UI

Документация API доступна по адресу: http://localhost:8080/swagger/index.html

## Управление проектом

- `make up` - поднять проект
- `make down` - остановить проект
- `make rebuild` - пересоздать проект (удалить все данные и поднять заново)
- `make logs` - посмотреть логи
- `make clean` - очистить все контейнеры и образы
- `make build` - собрать проект

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
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
└── Makefile
```