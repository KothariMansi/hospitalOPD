-- name: CreateClient :execresult
INSERT INTO Client (name, state, city, age)
VALUES (?, ?, ?, ?);

-- name: GetClient :one
SELECT * FROM Client WHERE id = ?;

-- name: ListClients :many
SELECT * FROM Client ORDER BY id LIMIT ? OFFSET ?;

-- name: UpdateClient :exec
UPDATE Client SET name = ?, state = ?, city = ?, age = ? WHERE id = ?;

-- name: DeleteClient :exec
DELETE FROM Client WHERE id = ?;
