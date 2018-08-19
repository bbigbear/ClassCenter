package models

import (
	"time"
)

type Case struct {
	Id          int64
	TeacherId   string
	TeacherName string
	CaseName    string
	Year        string
	CourseName  string
	Path        string
	Status      string
	Evaluate    string
	CreateTime  time.Time
}

type TeachGroup struct {
	Id          int64
	TeacherId   string
	TeacherName string
	Name        string
	GroupLeader string
	GroupMember string
}
