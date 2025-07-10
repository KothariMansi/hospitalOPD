-- name: CreateCheckUpTime :execresult
INSERT INTO CheckUpTime (morning, evening, night)
VALUES (?, ?, ?);

-- name: GetCheckUpTime :one
SELECT * FROM CheckUpTime WHERE id = ?;

-- name: ListCheckUpTimes :many
SELECT * FROM CheckUpTime ORDER BY id;

-- name: UpdateCheckUpTime :exec
UPDATE CheckUpTime SET morning = ?, evening = ?, night = ? WHERE id = ?;

-- name: DeleteCheckUpTime :exec
DELETE FROM CheckUpTime WHERE id = ?;
