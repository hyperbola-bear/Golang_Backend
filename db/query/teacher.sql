-- name: AddTeacherStudent :exec
INSERT INTO teachers (
    teacher,
    student 
) VALUES (
    $1, $2
);

-- name: GetTeachersStudents :many

SELECT student
FROM teachers
WHERE teacher = ANY($1::text[])
GROUP BY student
HAVING COUNT(DISTINCT teacher) = (
    SELECT COUNT(DISTINCT t) FROM UNNEST($1::text[]) AS t
);

-- name: GetTeacherStudents :many
SELECT student
FROM teachers
WHERE teacher = $1;


