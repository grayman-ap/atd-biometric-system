package db

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/grayman-ap/student_attendance/util"
	"github.com/stretchr/testify/require"
)

func createRandomTutor(t *testing.T) Tutor {
	dept := createRandomDepartment(t)
	arg := CreateTutorParams{
		StaffID:    util.Random(8),
		FirstName:  util.Random(5),
		LastName:   util.Random(5),
		Email:      util.RandomEmail(),
		Department: dept.DepartmentID,
	}

	tutor, err := testQueries.CreateTutor(context.Background(), arg)
	require.NoError(t, err)

	require.Equal(t, arg.StaffID, tutor.StaffID)
	require.Equal(t, arg.FirstName, tutor.FirstName)
	require.Equal(t, arg.LastName, tutor.LastName)
	require.Equal(t, arg.Email, tutor.Email)
	require.NotEmpty(t, tutor.CreatedAt)

	return tutor
}

func TestCreateTutor(t *testing.T) {
	createRandomTutor(t)
}

func TestGetTutor(t *testing.T) {
	tutor1 := createRandomTutor(t)

	tutor2, err := testQueries.GetTutor(context.Background(), tutor1.StaffID)
	require.NoError(t, err)
	require.NotEmpty(t, tutor2)

	require.Equal(t, tutor1.StaffID, tutor2.StaffID)
	require.Equal(t, tutor1.FirstName, tutor2.FirstName)
	require.Equal(t, tutor1.LastName, tutor2.LastName)
	require.WithinDuration(t, tutor1.CreatedAt, tutor2.CreatedAt, time.Second)
}

func TestListTutor(t *testing.T) {
	n := int(10)

	for i := 0; i < n; i++ {
		createRandomTutor(t)
	}

	arg := ListTutorsParams{
		Limit:  5,
		Offset: 5,
	}
	tutors, err := testQueries.ListTutors(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, tutors, 5)

	for _, tutor := range tutors {
		require.NotEmpty(t, tutor)
		fmt.Print("hello")
	}

}

func TestUpdateTutor(t *testing.T) {
	tutor1 := createRandomTutor(t)
	dept := createRandomDepartment(t)
	arg := UpdateTutorParams{
		StaffID:    tutor1.StaffID,
		FirstName:  "Peter",
		LastName:   "Adeshina",
		Email:      "peteradeshina3@futminna.com",
		Department: dept.DepartmentID,
	}
	updatedStudent, err := testQueries.UpdateTutor(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedStudent)

	require.NotEqual(t, tutor1.FirstName, updatedStudent.FirstName)
	require.NotEqual(t, tutor1.LastName, updatedStudent.LastName)
	require.NotEqual(t, tutor1.Email, updatedStudent.Email)
	require.Equal(t, tutor1.StaffID, updatedStudent.StaffID)
	require.WithinDuration(t, tutor1.CreatedAt, updatedStudent.CreatedAt, time.Second)
}

func TestDeleteTutor(t *testing.T) {
	tutor1 := createRandomTutor(t)

	err := testQueries.DeleteTutor(context.Background(), tutor1.StaffID)
	require.NoError(t, err)

	tutor2, err := testQueries.GetTutor(context.Background(), tutor1.StaffID)
	require.Error(t, err)
	require.Empty(t, tutor2)
}
