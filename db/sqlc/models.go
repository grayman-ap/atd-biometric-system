// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"time"
)

type Attendance struct {
	ID             int64     `json:"id"`
	Student        string    `json:"student"`
	CourseCode     string    `json:"course_code"`
	MarkStudent    bool      `json:"mark_student"`
	LastAttendance time.Time `json:"last_attendance"`
	CreatedAt      time.Time `json:"created_at"`
}

type Course struct {
	ID         int64     `json:"id"`
	Student    string    `json:"student"`
	Tutor      string    `json:"tutor"`
	CourseCode string    `json:"course_code"`
	Location   string    `json:"location"`
	Duration   time.Time `json:"duration"`
	CreatedAt  time.Time `json:"created_at"`
}

type Department struct {
	ID             int64     `json:"id"`
	DepartmentName string    `json:"department_name"`
	Student        string    `json:"student"`
	Tutor          string    `json:"tutor"`
	CreatedAt      time.Time `json:"created_at"`
}

type Student struct {
	ID        int64     `json:"id"`
	StudentID string    `json:"student_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type Tutor struct {
	ID        int64     `json:"id"`
	StaffID   string    `json:"staff_id"`
	CreatedAt time.Time `json:"created_at"`
}
