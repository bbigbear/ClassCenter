package models

import (
	"time"
)

type Case struct {
	Id         int64
	CaseName   string
	Year       string
	CourseName string
	Path       string
	Status     string
	Evaluate   string
	CreateTime time.Time
}

type TeachGroup struct {
	Id          int64
	Name        string
	GroupLeader string
	GroupMember string
}
