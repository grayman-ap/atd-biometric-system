package db

import (
	"context"
	"testing"

	"github.com/grayman-ap/student_attendance/util"
	"github.com/stretchr/testify/require"
)

func createRandomStudent(t *testing.T) Student {
	arg := CreateStudentParams{
		StudentID: util.Random(8),
		FirstName: util.Random(5),
		LastName:  util.Random(5),
		Email:     util.RandomEmail(),
	}

	student, err := testQueries.CreateStudent(context.Background(), arg)
	require.NoError(t, err)

	require.Equal(t, arg.StudentID, student.StudentID)
	require.Equal(t, arg.FirstName, student.FirstName)
	require.Equal(t, arg.LastName, student.LastName)
	require.Equal(t, arg.Email, student.Email)
	require.NotEmpty(t, student.ID)
	require.NotEmpty(t, student.CreatedAt)

	return student
}

func TestCreateStudent(t *testing.T) {
	createRandomStudent(t)
}
