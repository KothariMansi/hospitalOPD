-- name: CreateEmergencyLine :execresult
INSERT INTO EmergencyLine (reg_time, token_number, client_id, doctor_id, isChecked, checked_time)
VALUES (?, ?, ?, ?, ?, ?);

-- name: GetEmergencyLine :one
SELECT * FROM EmergencyLine WHERE id = ?;

-- name: ListEmergencyLines :many
SELECT * FROM EmergencyLine ORDER BY reg_time DESC LIMIT ? OFFSET ?;

-- name: UpdateEmergencyLine :exec
UPDATE EmergencyLine SET
  reg_time = ?, token_number = ?, client_id = ?, doctor_id = ?, isChecked = ?, checked_time = ?
WHERE id = ?;

-- name: DeleteEmergencyLine :exec
DELETE FROM EmergencyLine WHERE id = ?;
