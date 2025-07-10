-- name: CreateClient :one
INSERT INTO Client(
    name, state, city, age
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;

-- name: GetClient :one
SELECT * FROM Client
WHERE id=$1 LIMIT 1;

-- name: ListClient :many
SELECT * FROM Client
ORDER BY id
LIMIT $1 
OFFSET $2;

-- name: UpdateClient :one
UPDATE Client
SET name = $2,
state = $3,
city = $4,
age = $5
WHERE id=$1
RETURNING *;

-- name: DeleteClient :exec
DELETE FROM Client WHERE id=$1;