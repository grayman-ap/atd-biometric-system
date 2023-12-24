-- name: CreateAttendance :one
INSERT INTO attendance(
    student,
    course_code,
    mark_student,
    last_attendance
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetAttendance :one 
SELECT * FROM attendance
WHERE id = 1 LIMIT $1;

-- name: UpdateAttendance :one
UPDATE attendance
SET student = $2, course_code = $3, mark_student = $4, last_attendance = $5
WHERE id = $1
RETURNING *;

-- name :ListAttendance :many
SELECT * FROM attendance
ORDER BY id
LIMIT $1
OFFSET $2;

-- name :DeleteAttendance :exec 
DELETE FROM attendance
WHERE id = $1;