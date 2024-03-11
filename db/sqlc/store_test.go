package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)
func TestAddTeacherStudents(t *testing.T) {
	store := NewStore(testDB)
	fmt.Println("store: ", store)
	arg := AddTeacherStudentsParams{
		Teacher: "teacher1",
		Students: []string{"student1", "student2"},
	}
	err := store.AddTeacherStudents(context.Background(), arg)
	require.NoError(t, err)
}

func TestGetTeachersStudents(t *testing.T) {
	store := NewStore(testDB)
	arg := AddTeacherStudentsParams{
		Teacher: "teacher3",
		Students: []string{"student3", "student4"},
	}
	err := store.AddTeacherStudents(context.Background(), arg)
	require.NoError(t, err)

	tests := []struct {
		name    string
		teacher string
		want    []string
	}{
		{
			name:    "valid",
			teacher: "teacher3",
			want:    []string{"student3", "student4"},
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
			require.Equal(t, tt.want, got.Students)
		})
	}
}
