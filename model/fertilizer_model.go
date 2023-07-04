package model

import "time"

type FertilizerModel struct {
	Id       int
	Name     string
	Stock    int
	CreateAt time.Time
	UpdateAt time.Time
	CreateBy string
	UpdateBy string
}
