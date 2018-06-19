package models

import (
	"time"
)

type Jxkc struct {
	Id                 int64
	ElectiveCategory   string
	CourseCategory     string
	CourseId           string
	CourseName         string
	CourseEnName       string
	CourseAlias        string
	CourseCost         string
	OpeningInstitution string
	OpeningCoefficient string
	ExaminationForm    string
	TeachingMethods    string
	TotalHours         float64
	TheoreticalHours   float64
	PracticeHours      float64
	OtherHours         float64
	Bibliography       string
	Credit             float64
	CourseIntroduction string
	CourseNumber       string
	NoCourseNumber     string
	CreateTime         time.Time
	Status             string
}
