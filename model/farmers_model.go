package model

import "time"

type Farmers struct {
	Id          int
	Name        string
	Address     string
	PhoneNumber string
	CreateAt    time.Time
	UpdateAt    time.Time
	CreateBy    string
	UpdateBy    string
}
