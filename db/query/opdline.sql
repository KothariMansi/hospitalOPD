-- name: CreateOPDLine :execresult
INSERT INTO OPDLine (reg_time, token_number, client_id, doctor_id, isChecked, checked_time)
VALUES (?, ?, ?, ?, ?, ?);

-- name: GetOPDLine :one
SELECT * FROM OPDLine WHERE id = ?;

-- name: ListOPDLines :many
SELECT * FROM OPDLine ORDER BY reg_time DESC LIMIT ? OFFSET ?;

-- name: UpdateOPDLine :exec
UPDATE OPDLine SET
  reg_time = ?, token_number = ?, client_id = ?, doctor_id = ?, isChecked = ?, checked_time = ?
WHERE id = ?;

-- name: DeleteOPDLine :exec
DELETE FROM OPDLine WHERE id = ?;

-- name: ListOPDWithDetails :many
SELECT 
  o.id, o.reg_time, o.token_number, o.isChecked, o.checked_time,
  c.name AS client_name,
  d.name AS doctor_name
FROM OPDLine o
JOIN Client c ON o.client_id = c.id
JOIN Doctor d ON o.doctor_id = d.id
ORDER BY o.reg_time DESC
LIMIT ? OFFSET ?;

