// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"
)

type Querier interface {
	CreateStudent(ctx context.Context, arg CreateStudentParams) (Student, error)
	DeleteStudent(ctx context.Context, studentID string) error
	GetStudent(ctx context.Context, limit int32) (Student, error)
	ListStudents(ctx context.Context, arg ListStudentsParams) ([]Student, error)
	UpdateStudent(ctx context.Context, arg UpdateStudentParams) (Student, error)
}

var _ Querier = (*Queries)(nil)
