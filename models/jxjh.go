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
