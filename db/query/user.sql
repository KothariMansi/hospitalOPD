-- name: CreateUser :execresult
INSERT INTO User (username, full_name, hashed_password, state, city, gender, password_changed_at, age)
VALUES (?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetUser :one
SELECT * FROM User WHERE id = ?;

-- name: ListUsers :many
SELECT * FROM User ORDER BY id LIMIT ? OFFSET ?;

-- name: UpdateUser :exec
UPDATE User SET username = ?, full_name = ?, hashed_password = ?, state = ?, city = ?, gender = ?, password_changed_at = ?, age = ?
WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM User WHERE id = ?;

-- name: CountUsers :one
SELECT COUNT(*) FROM User;

-- name: SearchUsersByName :many
SELECT * FROM User
WHERE full_name LIKE ?
ORDER BY id
LIMIT ? OFFSET ?;

-- name: ListUsersByGenderAndCity :many
SELECT * FROM User
WHERE gender = ? AND city = ?
ORDER BY id
LIMIT ? OFFSET ?;
