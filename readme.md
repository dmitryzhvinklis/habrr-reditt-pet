![photo_2024-07-10_22-58-10](https://github.com/dmitryzhvinklis/habrr-reditt-pet/assets/161613076/ec25a224-a5ae-4c83-ae72-55bb05d923e7)

Сборка приложения:

```
make build
```
Это соберет ваше приложение и создаст исполняемый файл app.

Очистка собранных файлов: 
```
make clean
```
Удаляет собранные файлы и исполняемый файл.

Форматирование кода:

```
make fmt
```
Применяет форматирование кода с помощью go fmt.

Установка зависимостей:
```
make install
```
Устанавливает зависимости вашего проекта.

Запуск приложения:
```
make run
```
Собирает и запускает приложение.

Запуск тестов:
```
make test
```
Запускает тесты вашего проекта.

Поднять контейнеры:
```
make up
```
Эта команда запустит все сервисы, описанные в файле docker-compose.yml.

Сгенерировать пост - 

```
mutation {
  createPost(user_id: 1, content: "Новый пост", allow_comments: true) {
    id
    user_id
    content
    allow_comments
  }
}
```
Сгенерировать комментарий к посту - 
```
mutation {
  createComment(post_id: 1, user_id: 1, parent_comment_id: 1, content: "Отличный пост") {
    ID
    PostID
    UserID
    ParentCommentID
    Content
    CreatedAt
  }
}

```