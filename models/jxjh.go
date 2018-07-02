package models

import (
	"time"
)

type Jxjh struct {
	Id               int64
	PlanId           string
	Major            string
	PlanGrade        string
	PlanClass        string
	TotalCredits     float64
	SettingTime      time.Time
	SchoolSystem     float64
	TrainTarget      string
	ApplyTime        time.Time
	ApplyDescription string
	Status           string
}

type Jxrw struct {
	Id         int64
	CourseId   string
	CourseName string
	Term       string
	PlanId     string
	State      int
	Core       int
}
