package models

//style 0 选修1 必修2 教师3教室
type Kbsz struct {
	Id        int64
	Name      string
	StartTime string
	EndTime   string
	Mon       string
	Tue       string
	Wed       string
	Thu       string
	Fri       string
	Sat       string
	Sun       string
	Style     int
	TeacherId string
	ClassId   string
}

type KbszSetting struct {
	Id       int64
	AmNum    int
	PmNum    int
	NightNum int
}
