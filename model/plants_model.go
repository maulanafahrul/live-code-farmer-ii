package model

import "time"

type PlantModel struct {
	Id       int
	Name     string
	CreateAt time.Time
	UpdateAt time.Time
	CreateBy string
	UpdateBy string
}
