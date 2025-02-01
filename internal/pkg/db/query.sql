-- name: GetUser :one
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = ? LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY name;

-- name: CreateUser :one
INSERT INTO users (
  email,
  password,
  name,
  bio,
  update_at
) VALUES (
  ?, ?, ?, ?, ?
)
RETURNING *;

-- name: UpdateUser :exec
UPDATE users SET
email = ?,
name = ?,
password = ?,
bio = ?,
update_at = ?
WHERE id = ?
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?;

-- name: GetPost :one
SELECT * FROM posts
WHERE id = ? LIMIT 1;

-- name: ListPostByUserID :many
SELECT * FROM posts
WHERE user_id = ?
ORDER BY name;

-- name: CreatePost :one
INSERT INTO posts (
  user_id,
  title,
  body,
  tags
) VALUES (
  ?, ?, ?, ?
)
RETURNING *;

-- name: UpdatePost :exec
UPDATE posts SET
title = ?,
body = ?,
tags = ?,
update_at = ?
WHERE id = ?
RETURNING *;

-- name: DeletePosts :exec
DELETE FROM posts
WHERE id = ?;
