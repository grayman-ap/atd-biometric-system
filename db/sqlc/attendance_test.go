package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomAttendance(t *testing.T) Attendance {
	student := createRandomStudent(t)
	course := createRandomCourse(t)
	arg := CreateAttendanceParams{
		Student:        student.StudentID,
		CourseCode:     course.CourseCode,
		MarkStudent:    true,
		LastAttendance: time.Now().Add(time.Hour),
	}

	attendance, err := testQueries.CreateAttendance(context.Background(), arg)
	require.NoError(t, err)

	require.Equal(t, arg.Student, attendance.Student)
	require.Equal(t, arg.CourseCode, attendance.CourseCode)
	require.Equal(t, arg.CourseCode, attendance.CourseCode)
	require.NotEmpty(t, attendance.LastAttendance)
	require.NotEmpty(t, course.CreatedAt)

	return attendance
}

func TestCreateAttendance(t *testing.T) {
	createRandomAttendance(t)
}

func TestGetAttendance(t *testing.T) {
	attendance1 := createRandomAttendance(t)

	attendance2, err := testQueries.GetAttendance(context.Background(), attendance1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, attendance2)

	require.Equal(t, attendance1.Student, attendance2.Student)
	require.Equal(t, attendance1.CourseCode, attendance2.CourseCode)
	require.Equal(t, attendance1.CourseCode, attendance2.CourseCode)
	require.Equal(t, attendance1.LastAttendance, attendance2.LastAttendance)
	require.WithinDuration(t, attendance1.CreatedAt, attendance2.CreatedAt, time.Second)
}

func TestListAttendance(t *testing.T) {
	n := int(4)

	for i := 0; i < n; i++ {
		createRandomAttendance(t)
	}

	arg := ListAttendanceParams{
		Limit:  4,
		Offset: 4,
	}
	attendances, err := testQueries.ListAttendance(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, attendances, 4)

	for _, attendance := range attendances {
		require.NotEmpty(t, attendance)
	}

}

func TestUpdateAttendance(t *testing.T) {
	attendance1 := createRandomAttendance(t)
	student := createRandomStudent(t)
	course := createRandomCourse(t)
	arg := UpdateAttendanceParams{
		ID:             attendance1.ID,
		CourseCode:     course.CourseCode,
		Student:        student.StudentID,
		MarkStudent:    false,
		LastAttendance: time.Now(),
	}
	updatedAttendance, err := testQueries.UpdateAttendance(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedAttendance)

	require.NotEqual(t, attendance1.Student, updatedAttendance.Student)
	require.NotEqual(t, attendance1.CourseCode, updatedAttendance.CourseCode)
	require.NotEqual(t, attendance1.CourseCode, updatedAttendance.CourseCode)
	require.NotEqual(t, attendance1.LastAttendance, updatedAttendance.LastAttendance)
	require.Equal(t, attendance1.ID, updatedAttendance.ID)
	require.WithinDuration(t, attendance1.CreatedAt, updatedAttendance.CreatedAt, time.Second)
}

func TestDeleteAttendance(t *testing.T) {
	attendance := createRandomAttendance(t)

	err := testQueries.DeleteAttendance(context.Background(), attendance.ID)
	require.NoError(t, err)

	course2, err := testQueries.GetAttendance(context.Background(), attendance.ID)
	require.Error(t, err)
	require.Empty(t, course2)
}
