package main

import (
	"log"
	"net/http"

	"github.com/graphql-go/handler"
	"github.com/joho/godotenv"

	"ozon-test/internal/db"
	"ozon-test/internal/graph"
)

func main() {
	// Загрузка переменных окружения из файла .env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Инициализация базы данных
	db.InitDB()

	// Создание GraphQL схемы
	schema := graph.NewSchema()

	// Настройка GraphQL сервера
	graphqlHandler := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true, // Включаем GraphiQL интерфейс для удобства разработки
	})
	// Endpoint для GraphQL API
	http.Handle("/graphql", graphqlHandler)

	// Запуск сервера
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
