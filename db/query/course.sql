-- name: CreateCourse :one
INSERT INTO course (
    course_code,
    department,
    number_of_student,
    course_title,
    course_unit,
    venue,
    start_time,
    end_time
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;

-- name: GetCourse :one
SELECT * FROM course
WHERE course_code = 1;

-- name: ListCourse :many
SELECT * FROM course 
ORDER BY course_code
LIMIT $1
OFFSET $2;

-- name: UpdateCourse :one
UPDATE course
SET department = $2, number_of_student = $3, course_title = $4, course_unit= $5, venue = $6, start_time = $7, end_time = $8
WHERE course_code = $1
RETURNING *;

-- name: DeleteCourse :exec
DELETE FROM course
WHERE course_code = $1;