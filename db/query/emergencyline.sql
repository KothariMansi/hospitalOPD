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

-- name: CountEmergencyLines :one
SELECT COUNT(*) FROM EmergencyLine;

-- name: ListEmergencyWithDetails :many
SELECT 
  e.id, e.reg_time, e.token_number, e.isChecked, e.checked_time,
  c.name AS client_name,
  d.name AS doctor_name
FROM EmergencyLine e
JOIN Client c ON e.client_id = c.id
JOIN Doctor d ON e.doctor_id = d.id
ORDER BY e.reg_time DESC
LIMIT ? OFFSET ?;

-- name: SearchEmergencyByDate :many
SELECT 
  e.id, e.reg_time, e.token_number, e.isChecked, e.checked_time,
  c.name AS client_name,
  d.name AS doctor_name
FROM EmergencyLine e
JOIN Client c ON e.client_id = c.id
JOIN Doctor d ON e.doctor_id = d.id
WHERE e.reg_time BETWEEN ? AND ?
ORDER BY e.reg_time DESC
LIMIT ? OFFSET ?;
