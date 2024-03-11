package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
		Queries: New(db),
	}
}


func (store *Store) ExecTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %v, rb error: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

type AddTeacherStudentsParams struct {
	Teacher string `json:"teacher" binding:"required"` 
	Students []string `json:"students" binding:"required"`
}

type ReturnTeacherStudentResult struct {
	Students []string `json:"students" binding:"required"`
}

func (store *Store) AddTeacherStudents(ctx context.Context, arg AddTeacherStudentsParams)  error {
	return store.ExecTx(ctx, func(q *Queries) error {
		var err error
		for _, student := range arg.Students {
			err := q.AddTeacherStudent(ctx, AddTeacherStudentParams{
				Teacher: arg.Teacher,
				Student: student,
			})
			if err != nil {
				return err
			} 
		}
		return err
	})
}

func (store *Store) GetTeachersStudents(ctx context.Context, teacher []string) (ReturnTeacherStudentResult, error) {
	var result ReturnTeacherStudentResult
	err := store.ExecTx(ctx, func(q *Queries) error {
		var err error
		students, err := q.GetTeachersStudents(ctx, teacher)
		if err != nil {
			return err
		}
		result.Students = students
		return nil
	})
	return result, err
}

func (store *Store) AddSuspension(ctx context.Context, student string) error {
	return store.ExecTx(ctx, func(q *Queries) error {
		return q.AddSuspension(ctx, student)
	})
}

func (store *Store) GetSuspendedStudents(ctx context.Context) ([]string, error) {
	var students []string
	err := store.ExecTx(ctx, func(q *Queries) error {
		var err error
		students, err = q.GetSuspendedStudents(ctx)
		if err != nil {
			return err
		}
		return nil
	})
	return students, err
}