package model

import "time"

type BillDetailModel struct {
	Id                int
	BillId            int
	FertilizerPriceId int
	Qty               int
	CreateAt          time.Time
	UpdateAt          time.Time
	CreateBy          string
	UpdateBy          string
}
