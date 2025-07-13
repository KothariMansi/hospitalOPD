-- name: CreateDoctor :execresult
INSERT INTO Doctor (
  name, username, password, hospital_id,
  resident_address, checkup_time_id, is_on_leave
) VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: GetDoctor :one
SELECT * FROM Doctor WHERE id = ?;

-- name: ListDoctors :many
SELECT * FROM Doctor ORDER BY hospital_id
LIMIT ? OFFSET ?;

-- name: UpdateDoctor :exec
UPDATE Doctor SET
  name = ?,
  username = ?,
  password = ?,
  hospital_id = ?,
  resident_address = ?,
  checkup_time_id = ?,
  is_on_leave = ?
WHERE id = ?;

-- name: DeleteDoctor :exec
DELETE FROM Doctor WHERE id = ?;

-- name: CountDoctors :one
SELECT COUNT(*) FROM Doctor;

-- name: SearchDoctorsByName :many
SELECT * FROM Doctor
WHERE name LIKE ?
ORDER BY id
LIMIT ? OFFSET ?;

-- name: ListDoctorsWithHospital :many
SELECT 
  d.id, d.name, d.username, d.password, d.hospital_id,
  h.name AS hospital_name,
  d.resident_address, d.checkup_time_id, d.is_on_leave,
  d.created_at, d.updated_at
FROM Doctor d
JOIN Hospital h ON d.hospital_id = h.id
ORDER BY d.id
LIMIT ? OFFSET ?;

-- name: ListDoctorsByHospital :many
SELECT * FROM Doctor
WHERE hospital_id = ?
ORDER BY id
LIMIT ? OFFSET ?;

-- name: ListDoctorsOnLeave :many
SELECT * FROM Doctor
WHERE is_on_leave = TRUE
ORDER BY id
LIMIT ? OFFSET ?;

-- name: GetDoctorByUsername :one
SELECT * FROM Doctor
WHERE username = ?;
