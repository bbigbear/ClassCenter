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
	CreateTime  time.Time
}
