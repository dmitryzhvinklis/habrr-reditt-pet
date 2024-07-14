GO := go
GOCMD := $(GO)
GOBUILD := $(GOCMD) build
GOTEST := $(GOCMD) test
GOCLEAN := $(GOCMD) clean
GOFMT := $(GOCMD) fmt
GOINSTALL := $(GOCMD) install
BINARY_NAME := app

.DEFAULT_GOAL := help

# Помощь
help:
	@echo "Available targets:"
	@echo "  build          : Build the application"
	@echo "  clean          : Clean the build files"
	@echo "  fmt            : Format the source code"
	@echo "  install        : Install dependencies"
	@echo "  run            : Run the application"
	@echo "  test           : Run tests"
	@echo "  test-cover     : Run tests and show coverage"
	@echo "  cover-html     : Generate HTML report of coverage"
	@echo "  up             : Bring up the container"
	@echo "  init_test_db   : Initialize the test database"

# Сборка приложения
build: test
	$(GOBUILD) -o $(BINARY_NAME) -v ./cmd/...

# Очистка собранных файлов
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

# Форматирование кода
fmt:
	$(GOFMT) ./...

# Установка зависимостей
install:
	$(GOINSTALL) ./...

# Запуск приложения
run: test
	$(GOBUILD) -o $(BINARY_NAME) -v ./cmd/...
	./$(BINARY_NAME)

# Создание тестовой БД
init_test_db:
	@echo "Initializing test database..."
	@psql -U postgres -tc "SELECT 1 FROM pg_database WHERE datname = 'ozon_testing_app'" | grep -q 1 || psql -U postgres -c "CREATE DATABASE ozon_testing_app;"
	@psql -U postgres -d ozon_testing_app -f init_db.sql

# Запуск тестов с инициализацией тестовой БД
test: init_test_db
	@echo "Running tests..."
	$(GOTEST) -v ./...

# Запуск тестов с покрытием
test-cover: init_test_db
	@echo "Running tests with coverage..."
	$(GOTEST) -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -func=coverage.out

# Генерация HTML отчета о покрытии
cover-html:
	$(GOCMD) tool cover -html=coverage.out -o coverage.html

# Поднятие контейнера
up:
	docker-compose up -d

.PHONY: help build clean fmt install run test test-cover cover-html up init_test_db
