-- name: CreateDoctorSpeciality :execresult
INSERT INTO DoctorSpeciality (speciality_id, docter_id)
VALUES (?, ?);

-- name: GetDoctorSpeciality :one
SELECT * FROM DoctorSpeciality WHERE id = ?;

-- name: ListDoctorSpecialities :many
SELECT * FROM DoctorSpeciality ORDER BY id LIMIT ? OFFSET ?;

-- name: UpdateDoctorSpeciality :exec
UPDATE DoctorSpeciality SET speciality_id = ?, docter_id = ? WHERE id = ?;

-- name: DeleteDoctorSpeciality :exec
DELETE FROM DoctorSpeciality WHERE id = ?;
