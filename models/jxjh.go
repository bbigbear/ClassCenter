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
	Status     string
	Core       string
}

type Jxrw_teacher_allot struct {
	Id          int64
	PlanId      string
	CourseId    string
	TeacherId   string
	TeacherName string
	Campus      string
	Classroom   string
	WeekTime    int
	Status      string
}
