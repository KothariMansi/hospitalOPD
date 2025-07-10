-- name: CreateUser :execresult
INSERT INTO User (name, password, state, city, gender, age)
VALUES (?, ?, ?, ?, ?, ?);

-- name: GetUser :one
SELECT * FROM User WHERE id = ?;

-- name: ListUsers :many
SELECT * FROM User ORDER BY id LIMIT ? OFFSET ?;

-- name: UpdateUser :exec
UPDATE User SET name = ?, password = ?, state = ?, city = ?, gender = ?, age = ?
WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM User WHERE id = ?;
