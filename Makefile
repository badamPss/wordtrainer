.PHONY: up down build clean logs test lint swagger migrate migrate-down

# Поднять проект
up:
	docker-compose up --build -d

# Остановить проект
down:
	docker-compose down

# Пересоздать проект (удалить все данные и поднять заново)
rebuild:
	docker-compose down -v
	docker-compose up --build -d

# Посмотреть логи
logs:
	docker-compose logs -f

# Очистить все контейнеры и образы
clean:
	docker-compose down -v
	docker system prune -f

# Собрать проект
build:
	docker-compose build

# Запустить тесты
test:
	go test -v ./...

# Запустить линтер
lint:
	golangci-lint run

# Сгенерировать Swagger документацию
swagger:
	swag init -g cmd/main.go

# Запустить миграции
migrate:
	docker-compose exec app ./main migrate

# Откатить миграции
migrate-down:
	docker-compose exec app ./main migrate down 