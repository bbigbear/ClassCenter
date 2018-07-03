package models

import (
	"time"
)

type XkSetting struct {
	Id              int64
	PepoleMax       int
	ClassPepopleMax int
	WomenMax        int
	ManMax          int
}

type Xkrw struct {
	Id            int64
	TaskId        string
	Name          string
	Term          string
	StartTime     time.Time
	EndTime       time.Time
	TeacherEnable string
	StudentEnable string
}
