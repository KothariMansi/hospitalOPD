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
