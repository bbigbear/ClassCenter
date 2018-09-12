package models

import (
	"time"
)

type Jbjh struct {
	Id          int64
	Founder     string
	Title       string
	Explain     string
	Master      string
	Participant string
	Status      string
	FilePath    string
	CreateTime  time.Time
}

type Part struct {
	Participant []string
}
