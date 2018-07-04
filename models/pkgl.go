package models

type PkSetting struct {
	Id                 int64
	TogetherMax        int
	CultureMax         int
	ProfessionalMax    int
	AmNotPm            string
	TogetherFirst      string
	CultureAmFirst      string
	ProfessionalPmFirst  string
	TeacherMax         int
	ClassCourseWeekMax int
	EachClassDayMax    int
}
