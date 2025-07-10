-- name: CreateClient :exec
INSERT INTO Client (
    name, state, city, age
) VALUES (
    ?, ?, ?, ?
);

-- name: GetLastInsertedClient :one
SELECT * FROM Client WHERE id = LAST_INSERT_ID();

-- name: GetClient :one
SELECT * FROM Client
WHERE id = ? LIMIT 1;

-- name: ListClient :many
SELECT * FROM Client
ORDER BY id
LIMIT ? 
OFFSET ?;

-- name: UpdateClient :exec
UPDATE Client
SET name = ?,
state = ?,
city = ?,
age = ?
WHERE id = ?;

-- name: DeleteClient :exec
DELETE FROM Client WHERE id = ?;