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

type Xzztl struct {
	Id         int64
	XzzId      int64
	Title      string
	Content    string
	CreateTime time.Time
}
