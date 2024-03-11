-- name: AddSuspension :exec
INSERT INTO suspensions (
    student 
) VALUES (
    $1 
);

-- name: GetSuspendedStudents :many
SELECT * FROM suspensions;