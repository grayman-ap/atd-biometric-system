-- name: CreateTutor :one
INSERT INTO tutor (
    staff_id,
    first_name,
    last_name,
    email,
    department
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetTutor :one
SELECT * FROM tutor
WHERE staff_id = $1;

-- name: ListTutors :many
SELECT * FROM tutor 
ORDER BY staff_id
LIMIT $1
OFFSET $2;

-- name: UpdateTutor :one
UPDATE tutor
SET first_name = $2, last_name = $3, email = $4, department = $5
WHERE staff_id = $1
RETURNING *;

-- name: DeleteTutor :exec
DELETE FROM tutor
WHERE staff_id = $1;