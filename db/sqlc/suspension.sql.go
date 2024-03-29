// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: suspension.sql

package db

import (
	"context"
)

const addSuspension = `-- name: AddSuspension :exec
INSERT INTO suspensions (
    student 
) VALUES (
    $1 
)
`

func (q *Queries) AddSuspension(ctx context.Context, student string) error {
	_, err := q.db.ExecContext(ctx, addSuspension, student)
	return err
}

const getSuspendedStudents = `-- name: GetSuspendedStudents :many
SELECT student FROM suspensions
`

func (q *Queries) GetSuspendedStudents(ctx context.Context) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getSuspendedStudents)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var student string
		if err := rows.Scan(&student); err != nil {
			return nil, err
		}
		items = append(items, student)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
