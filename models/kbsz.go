package models

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
}

type KbszSetting struct {
	Id       int64
	AmNum    int
	PmNum    int
	NightNum int
}
