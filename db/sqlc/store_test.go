package db

import (
	"context"
	"testing"

	"example.com/golang_backend/util"
	"github.com/stretchr/testify/require"
)
type TeacherTwoStudent struct {
	Teacher   string   `json:"teacher"`
	Students []string `json:"students"`
}
func CreateRandomTeacherTwoStudent(t *testing.T) TeacherTwoStudent {
	store := NewStore(testDB)
	arg := AddTeacherStudentsParams{
		Teacher: util.RandomEmail(),
		Students: []string{util.RandomEmail(), util.RandomEmail()},
	}
	require.NotEmpty(t, arg.Teacher)
	require.NotEmpty(t, arg.Students)
	err := store.AddTeacherStudents(context.Background(), arg)
	require.NoError(t, err)
	var teacherTwoStudent TeacherTwoStudent
	teacherTwoStudent.Teacher = arg.Teacher
	teacherTwoStudent.Students = arg.Students
	return teacherTwoStudent
}
func TestAddTeacherStudents(t *testing.T) {
	CreateRandomTeacherTwoStudent(t)
}

// cannot enforce required tag in the struct
// func TestAddTeacherStudentsInvalidTeacher(t *testing.T) {
// 	store := NewStore(testDB)
// 	arg := AddTeacherStudentsParams{
// 		Teacher: "",
// 		Students: []string{"student1", "student2"},
// 	}
// 	err := store.AddTeacherStudents(context.Background(), arg)
// 	require.Error(t, err)
// }

func TestGetTeachersStudents(t *testing.T) {
	store := NewStore(testDB)
	teacherTwoStudent := CreateRandomTeacherTwoStudent(t)
	tests := []struct {
		name    string
		teacher string
		want    []string
	}{
		{
			name:    "valid",
			teacher: teacherTwoStudent.Teacher,
			want:    []string{teacherTwoStudent.Students[0], teacherTwoStudent.Students[1]},
		},
		{
			name:    "invalid",
			teacher: "teacher4",
			want:    []string(nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var teacher_arr []string
			teacher_arr = append(teacher_arr, tt.teacher)
			got, err := store.GetTeachersStudents(context.Background(), teacher_arr)
			require.NoError(t, err)
			if (got.Students == nil){
				require.Equal(t, tt.want, got.Students)
				return
			}
			require.Contains(t, tt.want, got.Students[0])
			require.Contains(t, tt.want, got.Students[1])
		})
	}
}

func TestGetTeachersStudentsTwoInvalidTeachers(t *testing.T) {
	store := NewStore(testDB)
	teacher_arr := []string{"teacher5", "teacher6"}
	got, err := store.GetTeachersStudents(context.Background(), teacher_arr)
	require.NoError(t, err)
	require.Equal(t, []string(nil), got.Students)
}