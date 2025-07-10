-- name: CreateHospital :execresult
INSERT INTO Hospital (name, photo, state, city, address, phone)
VALUES (?, ?, ?, ?, ?, ?);

-- name: GetHospital :one
SELECT * FROM Hospital WHERE id = ?;

-- name: ListHospitals :many
SELECT * FROM Hospital ORDER BY id LIMIT ? OFFSET ?;

-- name: UpdateHospital :exec
UPDATE Hospital
SET name = ?, photo = ?, state = ?, city = ?, address = ?, phone = ?
WHERE id = ?;

-- name: DeleteHospital :exec
DELETE FROM Hospital WHERE id = ?;
