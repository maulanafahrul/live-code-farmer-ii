package model

import "time"

type BillModel struct {
	Id       int
	FarmerId int
	Date     time.Time
	CreateAt time.Time
	UpdateAt time.Time
	CreateBy string
	UpdateBy string

	arrDetail []BillDetailModel
}
