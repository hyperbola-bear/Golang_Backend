package db

import (
	"context"
	"testing"

	"example.com/golang_backend/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomTeacherStudent(t *testing.T) Teacher {
	arg := AddTeacherStudentParams{
		Teacher: util.RandomEmail(),
		Student: util.RandomEmail(),
	}
	require.NotEmpty(t, arg.Teacher)
	require.NotEmpty(t, arg.Student)
	err := testQueries.AddTeacherStudent(context.Background(), arg)
	require.NoError(t, err)
	var teacher Teacher
	teacher.Teacher = arg.Teacher
	teacher.Student = arg.Student
	return teacher
}

func TestAddTeacherStudent(t *testing.T) {
	CreateRandomTeacherStudent(t)
}


func TestGetTeacherStudents(t *testing.T) {
	teacher := CreateRandomTeacherStudent(t)
	students, err := testQueries.GetTeacherStudents(context.Background(), teacher.Teacher)
	require.NoError(t, err)
	require.NotEmpty(t, students)
	require.Contains(t, students, teacher.Student)
}

func TestGetTeachersStudent(t *testing.T) {
	teacher := CreateRandomTeacherStudent(t)
	var teacher_arr []string
	teacher_arr = append(teacher_arr, teacher.Teacher)
	require.NotEmpty(t, teacher.Teacher)
	student, err := testQueries.GetTeachersStudents(context.Background(), teacher_arr)
	require.NoError(t, err)
	require.NotEmpty(t, student)
	require.Equal(t, teacher.Student, student[0])


}