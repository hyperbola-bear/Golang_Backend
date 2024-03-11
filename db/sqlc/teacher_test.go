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

func TestAddTeacherStudentDuplicate(t *testing.T) {
	teacher := CreateRandomTeacherStudent(t)
	err := testQueries.AddTeacherStudent(context.Background(), AddTeacherStudentParams{
		Teacher: teacher.Teacher,
		Student: teacher.Student,
	})
	require.Error(t, err)
}
// cannot test because there is no required tag in the struct
// func TestAddTeacherStudentInvalidTeacher(t *testing.T) {
// 	arg := AddTeacherStudentParams{
// 		Teacher: "",
// 		Student: util.RandomEmail(),
// 	}
// 	err := testQueries.AddTeacherStudent(context.Background(), arg)
// 	require.Error(t, err)
// }


func TestGetTeacherStudents(t *testing.T) {
	teacher := CreateRandomTeacherStudent(t)
	students, err := testQueries.GetTeacherStudents(context.Background(), teacher.Teacher)
	require.NoError(t, err)
	require.NotEmpty(t, students)
	require.Contains(t, students, teacher.Student)
}

func TestGetTeacherStudentsNoStudents(t *testing.T) {
	teacher := util.RandomEmail()
	students, err := testQueries.GetTeacherStudents(context.Background(), teacher)
	require.NoError(t, err)
	require.Empty(t, students)
}