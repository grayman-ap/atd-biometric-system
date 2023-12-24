package db

import (
	"context"
	"testing"
	"time"

	"github.com/grayman-ap/student_attendance/util"
	"github.com/stretchr/testify/require"
)

func createRandomCourse(t *testing.T) Course {
	dept := createRandomDepartment(t)
	arg := CreateCourseParams{
		CourseCode:      util.RandomCourseCode(),
		Department:      dept.DepartmentID,
		NumberOfStudent: util.RandomInt(1, 200),
		CourseTitle:     util.RandomEmail(),
		CourseUnit:      "2 Unit",
		Venue:           "LT2",
		StartTime:       time.Now(),
	}

	course, err := testQueries.CreateCourse(context.Background(), arg)
	require.NoError(t, err)

	require.Equal(t, arg.Department, course.Department)
	require.Equal(t, arg.NumberOfStudent, course.NumberOfStudent)
	require.Equal(t, arg.CourseTitle, course.CourseTitle)
	require.Equal(t, arg.CourseUnit, course.CourseUnit)
	require.Equal(t, arg.Venue, course.Venue)
	require.NotEmpty(t, course.CourseCode)
	require.NotEmpty(t, course.StartTime)
	require.NotEmpty(t, course.CreatedAt)

	return course
}

func TestCreateCourse(t *testing.T) {
	createRandomCourse(t)
}

func TestGetCourse(t *testing.T) {
	course1 := createRandomCourse(t)

	course2, err := testQueries.GetCourse(context.Background(), course1.CourseCode)
	require.NoError(t, err)
	require.NotEmpty(t, course2)

	require.Equal(t, course1.Department, course2.Department)
	require.Equal(t, course1.NumberOfStudent, course2.NumberOfStudent)
	require.Equal(t, course1.CourseTitle, course2.CourseTitle)
	require.Equal(t, course1.CourseUnit, course2.CourseUnit)
	require.Equal(t, course1.Venue, course2.Venue)
	require.NotEmpty(t, course1.CourseCode)
	require.NotEmpty(t, course1.StartTime)
	require.WithinDuration(t, course1.CreatedAt, course2.CreatedAt, time.Second)
}

func TestListCourse(t *testing.T) {
	n := int(4)

	for i := 0; i < n; i++ {
		createRandomCourse(t)
	}

	arg := ListCourseParams{
		Limit:  4,
		Offset: 4,
	}
	courses, err := testQueries.ListCourse(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, courses, 4)

	for _, course := range courses {
		require.NotEmpty(t, course)
	}

}

func TestUpdateCourse(t *testing.T) {
	course1 := createRandomCourse(t)
	dept := createRandomDepartment(t)
	arg := UpdateCourseParams{
		CourseCode:      course1.CourseCode,
		Department:      dept.DepartmentID,
		NumberOfStudent: util.RandomInt(1, 200),
		CourseTitle:     util.RandomEmail(),
		CourseUnit:      "3 Unit",
		Venue:           "LT4",
		StartTime:       time.Now(),
		EndTime:         time.Now(),
	}
	updatedCourse, err := testQueries.UpdateCourse(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedCourse)

	require.NotEqual(t, course1.Department, updatedCourse.Department)
	require.NotEqual(t, course1.NumberOfStudent, updatedCourse.NumberOfStudent)
	require.NotEqual(t, course1.CourseTitle, updatedCourse.CourseTitle)
	require.NotEqual(t, course1.CourseUnit, updatedCourse.CourseUnit)
	require.NotEqual(t, course1.Venue, updatedCourse.Venue)
	// require.Equal(t, course1.CourseCode, updatedCourse.CourseCode)
	require.NotEqual(t, course1.StartTime, updatedCourse.StartTime)
	require.NotEqual(t, course1.EndTime, updatedCourse.EndTime)
	require.WithinDuration(t, course1.CreatedAt, updatedCourse.CreatedAt, time.Second)
}

func TestDeleteCourse(t *testing.T) {
	course1 := createRandomCourse(t)

	err := testQueries.DeleteTutor(context.Background(), course1.CourseCode)
	require.NoError(t, err)

	course2, err := testQueries.GetTutor(context.Background(), course1.CourseCode)
	require.Error(t, err)
	require.Empty(t, course2)
}
