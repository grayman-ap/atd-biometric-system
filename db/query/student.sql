-- name: CreateStudent :one
INSERT INTO student (
    student_id,
    first_name,
    last_name,
    email
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetStudent :one
SELECT * FROM student
WHERE student_id = $1;

-- name: ListStudents :many
SELECT * FROM student 
ORDER BY student_id
LIMIT $1
OFFSET $2;

-- name: UpdateStudent :one
UPDATE student
SET first_name = $2, last_name = $3, email = $4
WHERE student_id = $1
RETURNING *;

-- name: DeleteStudent :exec
DELETE FROM student
WHERE student_id = $1;