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

# Сборка приложения
build:
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
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./cmd/...
	./$(BINARY_NAME)

# Запуск тестов
test:
	$(GOTEST) -v ./...
	
# Поднятие контейнера
up:
	docker-compose up -d

.PHONY: help build clean fmt install run test
