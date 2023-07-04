package model

import "time"

type FertilizerPricesModel struct {
	Id           int
	FertilizerId int
	Price        int
	IsActive     bool
	CreateAt     time.Time
	UpdateAt     time.Time
	CreateBy     string
	UpdateBy     string
}
