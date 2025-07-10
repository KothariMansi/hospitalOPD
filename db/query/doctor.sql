-- name: CreateDoctor :execresult
INSERT INTO Doctor (
  Name, username, password, hospital_id, resident_address,
  isMorning, isEvening, isNight, checkup_time_id, isOnLeave
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetDoctor :one
SELECT * FROM Doctor WHERE id = ?;

-- name: ListDoctors :many
SELECT * FROM Doctor ORDER BY id;

-- name: UpdateDoctor :exec
UPDATE Doctor SET
  Name = ?, username = ?, password = ?, hospital_id = ?, resident_address = ?,
  isMorning = ?, isEvening = ?, isNight = ?, checkup_time_id = ?, isOnLeave = ?
WHERE id = ?;

-- name: DeleteDoctor :exec
DELETE FROM Doctor WHERE id = ?;
