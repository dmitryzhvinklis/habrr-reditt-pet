package services

import (
	"os"
	"testing"

	"ozon-test/internal/db"

	"github.com/stretchr/testify/assert"
)

func TestCreatePost(t *testing.T) {
	// Установка переменной окружения для тестовой базы данных
	os.Setenv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/ozon_testing_app?sslmode=disable")

	err := db.InitDB()
	assert.NoError(t, err)

	post, err := CreatePost(1, "Тестовый контент", true)
	assert.NoError(t, err)
	assert.NotEqual(t, 0, post.ID)
	assert.Equal(t, 1, post.UserID)
	assert.Equal(t, "Тестовый контент", post.Content)
	assert.True(t, post.AllowComments)
}
