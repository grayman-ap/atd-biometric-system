package db

import (
	"context"
	"testing"
	"time"

	"github.com/grayman-ap/student_attendance/util"
	"github.com/stretchr/testify/require"
)

func createRandomDepartment(t *testing.T) Department {
	arg := CreateDepartmentParams{
		DepartmentID: util.Random(8),
		School:       util.RandomDepartment(),
	}

	dept, err := testQueries.CreateDepartment(context.Background(), arg)
	require.NoError(t, err)

	require.Equal(t, arg.School, dept.School)
	require.NotEmpty(t, dept.DepartmentID)
	require.NotEmpty(t, dept.CreatedAt)

	return dept
}

func TestCreateDepartment(t *testing.T) {
	createRandomDepartment(t)
}

func TestGetDepartment(t *testing.T) {
	dept1 := createRandomDepartment(t)

	dept2, err := testQueries.GetDepartment(context.Background(), dept1.DepartmentID)
	require.NoError(t, err)
	require.NotEmpty(t, dept2)

	require.Equal(t, dept1.School, dept2.School)

	require.Equal(t, dept1.DepartmentID, dept2.DepartmentID)
	require.WithinDuration(t, dept1.CreatedAt, dept2.CreatedAt, time.Second)
}

func TestListDepartment(t *testing.T) {
	n := int(10)

	for i := 0; i < n; i++ {
		createRandomDepartment(t)
	}

	arg := ListDepartmentsParams{
		Limit:  5,
		Offset: 5,
	}
	depts, err := testQueries.ListDepartments(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, depts, 5)

	for _, dept := range depts {
		require.NotEmpty(t, dept)
	}

}

func TestUpdateDepartment(t *testing.T) {
	dept1 := createRandomDepartment(t)

	arg := UpdateDepartmentParams{
		DepartmentID: dept1.DepartmentID,
		School:       "MMS",
	}
	updateDept, err := testQueries.UpdateDepartment(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updateDept)

	require.NotEqual(t, dept1.School, updateDept.School)
	require.Equal(t, dept1.DepartmentID, updateDept.DepartmentID)
	require.WithinDuration(t, dept1.CreatedAt, updateDept.CreatedAt, time.Second)
}

func TestDeleteDepartment(t *testing.T) {
	dept1 := createRandomDepartment(t)

	err := testQueries.DeleteStudent(context.Background(), dept1.DepartmentID)
	require.NoError(t, err)

	dept2, err := testQueries.GetStudent(context.Background(), dept1.DepartmentID)
	require.Error(t, err)
	require.Empty(t, dept2)
}
