-- name: GetUser :one
SELECT * FROM user
WHERE id = ? LIMIT 1;

-- name: ListUsers :many
SELECT * FROM user
ORDER BY name;

-- name: CreateUser :one
INSERT INTO user (
  name,
  bio
) VALUES (
  ?, ?
)
RETURNING *;

-- name: UpdateUser :exec
UPDATE user SET
name = ?,
bio = ?
WHERE id = ?
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM user
WHERE id = ?;

-- name: GetSubFlow :one
SELECT * FROM sub_flow
WHERE id = ? LIMIT 1;

-- name: ListSubFlows :many
SELECT * FROM sub_flow
ORDER BY name;

-- name: CreateFlow :one
INSERT INTO flow (
  name,
  description
) VALUES (
  ?, ?
)
RETURNING *;

-- name: UpdateFlow :exec
UPDATE flow SET
name = ?,
description = ?
WHERE id = ?
RETURNING *;

-- name: DeleteFlow :exec
DELETE FROM flow
WHERE id = ?;

-- name: CreateSubFlow :one
INSERT INTO sub_flow (
  name,
  description
) VALUES (
  ?, ?
)
RETURNING *;

-- name: UpdateSubFlow :exec
UPDATE sub_flow SET
name = ?,
description = ?
WHERE id = ?
RETURNING *;

-- name: DeleteSubFlow :exec
DELETE FROM sub_flow
WHERE id = ?;

-- name: GetNode :one
SELECT * FROM node
WHERE id = ? LIMIT 1;

-- name: ListNodes :many
SELECT * FROM node
WHERE sub_flow_id = ?
ORDER BY create_time;

-- name: CreateNode :one
INSERT INTO node (
  id,
  sub_flow_id,
  type,
  parent,
  position,
  styles,
  width,
  height,
  hidden,
  description
) VALUES (
  ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
)
RETURNING *;

-- name: UpdateNode :exec
UPDATE node SET
type = ?,
parent = ?,
position = ?,
styles = ?,
width = ?,
height = ?,
hidden = ?,
description = ?
WHERE id = ?
RETURNING *;

-- name: DeleteNode :exec
DELETE FROM node
WHERE id = ?;

-- name: GetEdge :one
SELECT * FROM edge
WHERE id = ? LIMIT 1;

-- name: ListEdges :many
SELECT * FROM edge
WHERE sub_flow_id = ?
ORDER BY create_time;

-- name: CreateEdge :one
INSERT INTO edge (
  id,
  sub_flow_id,
  source,
  target,
  type,
  label,
  hidden,
  marker_end,
  points
) VALUES (
  ?, ?, ?, ?, ?, ?, ?, ?, ?
)
RETURNING *;

-- name: UpdateEdge :exec
UPDATE edge SET
source = ?,
target = ?,
type = ?,
label = ?,
hidden = ?,
marker_end = ?,
points = ?
WHERE id = ?
RETURNING *;

-- name: DeleteEdge :exec
DELETE FROM edge
WHERE id = ?;
