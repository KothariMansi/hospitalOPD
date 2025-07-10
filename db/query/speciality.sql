-- name: CreateSpeciality :execresult
INSERT INTO Speciality (speciality_name)
VALUES (?);

-- name: GetSpeciality :one
SELECT * FROM Speciality WHERE id = ?;

-- name: ListSpecialities :many
SELECT * FROM Speciality ORDER BY id LIMIT ? OFFSET ?;

-- name: UpdateSpeciality :exec
UPDATE Speciality SET speciality_name = ? WHERE id = ?;

-- name: DeleteSpeciality :exec
DELETE FROM Speciality WHERE id = ?;

-- name: CountSpecialities :one
SELECT COUNT(*) FROM Speciality;

-- name: SearchSpecialitiesByName :many
SELECT * FROM Speciality
WHERE speciality_name LIKE ?
ORDER BY speciality_name
LIMIT ? OFFSET ?;
