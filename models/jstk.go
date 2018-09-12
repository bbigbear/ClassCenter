package models

import (
	"time"
)

type Jstk struct {
	Id           int64
	ProcessName  string
	DisplayOrder string
	Explain      string
	Staffing     string
	MinDay       int
}

type Tksq struct {
	Id            int64
	TkTime        time.Time
	Subject       string
	GiveClass     string
	TkNum         int
	CheckPeople   string
	TkStyle       string
	ChangeTeacher string
	ChangeSubject string
	ChangeClass   string
	ChangeTime    time.Time
	ChangeNum     int
	Status        string
}

type Staff struct {
	Id       int64
	Staffing []string
}
