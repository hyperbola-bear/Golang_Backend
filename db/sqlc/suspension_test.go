package db

import (
	"context"
	"testing"

	"example.com/golang_backend/util"
	"github.com/stretchr/testify/require"
)

func CreateSuspendedStudent(t *testing.T) string {
	student := util.RandomEmail()
	require.NotEmpty(t, student)
	err := testQueries.AddSuspension(context.Background(), student)
	require.NoError(t, err)
	return student
}

func TestAddSuspension(t *testing.T) {
	CreateSuspendedStudent(t)
}

func TestGetSuspendedStudents(t *testing.T) {
	student := CreateSuspendedStudent(t)

	students, err := testQueries.GetSuspendedStudents(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, students)
	require.Contains(t, students, student)
}