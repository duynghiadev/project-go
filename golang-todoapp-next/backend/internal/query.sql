-- name: GetTodos :one
SELECT * FROM todos
WHERE id = $1 LIMIT 1;

-- name: ListAllTodos :many
SELECT * FROM todos
ORDER BY id;

-- name: CreateTodos :one
INSERT INTO todos (
  title, content, createdAt
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateTodos :exec
UPDATE todos
  set title = $2,
  content = $3
WHERE id = $1;

-- name: DeleteTodos :exec
DELETE FROM todos
WHERE id = $1;