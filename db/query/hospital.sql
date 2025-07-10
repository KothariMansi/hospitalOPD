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

-- name: CountHospitals :one
SELECT COUNT(*) FROM Hospital;

-- name: SearchHospitalsByName :many
SELECT * FROM Hospital
WHERE name LIKE ?
ORDER BY id
LIMIT ? OFFSET ?;

-- name: ListHospitalsByLocation :many
SELECT * FROM Hospital
WHERE state = ? AND city = ?
ORDER BY id
LIMIT ? OFFSET ?;
