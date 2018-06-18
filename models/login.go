package models

type Login struct {
	Id       int64
	ReadName string `json:"real_name"`
	Tel      string `json:"tel"`
	UserId   string `json:"user_id"`
	Username string `json:"tusername"`
}
