-- name: GetUser :one
SELECT * FROM user
WHERE id = ? LIMIT 1;

-- name: GetUsers :many
SELECT * FROM user
ORDER BY name;

-- name: CreateAuthor :one
INSERT INTO user (
  name, bio
) VALUES (
  ?, ?
)
RETURNING *;

-- name: UpdateAuthor :exec
UPDATE user SET
name = ?,
bio = ?
WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM user
WHERE id = ?;
