package db

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitDB(t *testing.T) {
	// Установка переменной окружения для тестовой базы данных
	os.Setenv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/ozon_testing_app?sslmode=disable")

	err := InitDB()
	assert.NoError(t, err)
	assert.NotNil(t, DB)
}
