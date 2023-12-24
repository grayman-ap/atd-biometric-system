package db

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/grayman-ap/student_attendance/util"
	"github.com/stretchr/testify/require"
)

func createRandomStudent(t *testing.T) Student {
	dept := createRandomDepartment(t)
	arg := CreateStudentParams{
		StudentID:  util.Random(8),
		FirstName:  util.Random(5),
		LastName:   util.Random(5),
		Email:      util.RandomEmail(),
		Department: dept.DepartmentID,
	}

	student, err := testQueries.CreateStudent(context.Background(), arg)
	require.NoError(t, err)

	require.Equal(t, arg.StudentID, student.StudentID)
	require.Equal(t, arg.FirstName, student.FirstName)
	require.Equal(t, arg.LastName, student.LastName)
	require.Equal(t, arg.Department, student.Department)
	require.Equal(t, arg.Email, student.Email)
	require.NotEmpty(t, student.CreatedAt)

	return student
}

func TestCreateStudent(t *testing.T) {
	createRandomStudent(t)
}

func TestGetStudent(t *testing.T) {
	student1 := createRandomStudent(t)

	student2, err := testQueries.GetStudent(context.Background(), student1.StudentID)
	require.NoError(t, err)
	require.NotEmpty(t, student2)

	require.Equal(t, student1.StudentID, student2.StudentID)
	require.Equal(t, student1.FirstName, student2.FirstName)
	require.Equal(t, student1.Department, student2.Department)
	require.Equal(t, student1.LastName, student2.LastName)
	require.WithinDuration(t, student1.CreatedAt, student2.CreatedAt, time.Second)
}

func TestListStudent(t *testing.T) {
	n := int(10)

	for i := 0; i < n; i++ {
		createRandomStudent(t)
	}

	arg := ListStudentsParams{
		Limit:  5,
		Offset: 5,
	}
	students, err := testQueries.ListStudents(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, students, 5)

	for _, student := range students {
		require.NotEmpty(t, student)
		fmt.Print("hello")
	}

}

func TestUpdateStudent(t *testing.T) {
	student1 := createRandomStudent(t)
	dept := createRandomDepartment(t)
	arg := UpdateStudentParams{
		StudentID:  student1.StudentID,
		FirstName:  "Peter",
		LastName:   "Adeshina",
		Email:      "peteradeshina3@futminna.com",
		Department: dept.DepartmentID,
	}
	updatedStudent, err := testQueries.UpdateStudent(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedStudent)

	require.NotEqual(t, student1.FirstName, updatedStudent.FirstName)
	require.NotEqual(t, student1.LastName, updatedStudent.LastName)
	require.NotEqual(t, student1.Department, updatedStudent.Department)
	require.NotEqual(t, student1.Email, updatedStudent.Email)
	require.Equal(t, student1.StudentID, updatedStudent.StudentID)
	require.WithinDuration(t, student1.CreatedAt, updatedStudent.CreatedAt, time.Second)
}

func TestDeleteStudent(t *testing.T) {
	student1 := createRandomStudent(t)

	err := testQueries.DeleteStudent(context.Background(), student1.StudentID)
	require.NoError(t, err)

	student2, err := testQueries.GetStudent(context.Background(), student1.StudentID)
	require.Error(t, err)
	require.Empty(t, student2)
}
