-- name: CreateDepartment :one
INSERT INTO department(
    department_id, 
    school
) VALUES(
    $1, $2
) RETURNING *;

-- name: GetDepartment :one
SELECT * FROM department 
WHERE department_id = $1;

-- name: ListDepartments :many
SELECT * FROM department
ORDER BY department_id
LIMIT $1
OFFSET $2;

-- name: UpdateDepartment :one
UPDATE department
SET school = $2
WHERE department_id = $1
RETURNING *;

-- name: DeleteDepartment :exec
DELETE FROM department
WHERE department_id = $1;