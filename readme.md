![photo_2024-07-10_22-58-10](https://github.com/dmitryzhvinklis/habrr-reditt-pet/assets/161613076/ec25a224-a5ae-4c83-ae72-55bb05d923e7)

Сгенерировать пост - 

mutation {
  createPost(user_id: 1, content: "Новый пост", allow_comments: true) {
    id
    user_id
    content
    allow_comments
  }
}

Сгенерировать комментарий к посту - 

mutation {
  createComment(post_id: 1, user_id: 1, parent_comment_id: 1, content: "Отличный пост") {
    id
    post_id
    user_id
    parent_comment_id
    content
    created_at
  }
}

