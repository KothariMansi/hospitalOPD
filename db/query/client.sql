-- name: CreateClient :execresult
INSERT INTO Client (name, state, city, number, age)
VALUES (?, ?, ?, ?, ?);

-- name: GetClient :one
SELECT * FROM Client WHERE id = ?;

-- name: ListClients :many
SELECT * FROM Client ORDER BY id LIMIT ? OFFSET ?;

-- name: UpdateClient :exec
UPDATE Client SET name = ?, state = ?, city = ?, number = ?, age = ? WHERE id = ?;

-- name: DeleteClient :exec
DELETE FROM Client WHERE id = ?;

-- name: CountClients :one
SELECT COUNT(*) FROM Client;

-- name: SearchClientsByName :many
SELECT * FROM Client
WHERE name LIKE ?
ORDER BY id
LIMIT ? OFFSET ?;

-- name: ListClientsByLocation :many
SELECT * FROM Client
WHERE city = ? AND state = ?
ORDER BY id
LIMIT ? OFFSET ?;
