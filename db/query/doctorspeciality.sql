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

-- name: ListSpecialitiesByDoctorID :many
SELECT s.id, s.speciality_name
FROM DoctorSpeciality ds
JOIN Speciality s ON ds.speciality_id = s.id
WHERE ds.docter_id = ?
ORDER BY s.speciality_name;

-- name: ListDoctorsBySpecialityID :many
SELECT d.*
FROM DoctorSpeciality ds
JOIN Doctor d ON ds.docter_id = d.id
WHERE ds.speciality_id = ?
ORDER BY d.name
LIMIT ? OFFSET ?;

