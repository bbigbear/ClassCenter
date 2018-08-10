package models

import (
	"time"
)

type Xzz struct {
	Id          int64
	Name        string
	GroupLeader string
	GroupMember string
	Remark      string
	CreateTime  time.Time
}
