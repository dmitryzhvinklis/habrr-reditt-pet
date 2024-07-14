package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	// Создадим временный файл .env
	file, err := os.Create(".env")
	assert.NoError(t, err)
	defer os.Remove(".env")

	_, err = file.WriteString("DATABASE_URL=postgres://user:password@localhost:5432/dbname\n")
	assert.NoError(t, err)
	file.Close()

	// Загружаем конфигурацию
	LoadConfig()

	// Проверяем, что переменная окружения загружена
	databaseURL := os.Getenv("DATABASE_URL")
	assert.Equal(t, "postgres://user:password@localhost:5432/dbname", databaseURL)
}
