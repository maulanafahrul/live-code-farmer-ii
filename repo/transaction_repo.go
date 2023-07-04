package repo

import (
	"database/sql"
	"fmt"
	"live-code-farmer-ii/model"
)

type TransactionRepo interface {
	List() (*[]model.BillModel, error)
	Create(*model.BillModel) error
	GetTransactionDetailsByTrxBillId(int) (*[]model.BillDetailModel, error)
}

type transactionRepoImpl struct {
	db *sql.DB
}

func NewTransactionRepo(db *sql.DB) TransactionRepo {
	return &transactionRepoImpl{
		db: db,
	}
}

func (trxRepo *transactionRepoImpl) List() (*[]model.BillModel, error) {
	qry := "SELECT id, farmer_id, date, create_at, update_at, create_by, update_by FROM tr_bills"

	rows, err := trxRepo.db.Query(qry)
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction headers: %v", err)
	}
	defer rows.Close()

	var bills []model.BillModel

	for rows.Next() {
		header := model.BillModel{}
		err := rows.Scan(&header.Id, &header.FarmerId, &header.Date, &header.CreateAt, &header.UpdateAt, &header.CreateBy, &header.UpdateBy)
		if err != nil {
			return nil, fmt.Errorf("failed to scan bill header: %v", err)
		}

		details, err := trxRepo.GetTransactionDetailsByTrxBillId(header.Id)
		if err != nil {
			return nil, fmt.Errorf("failed to get transaction details for header No %v: %v", header.Id, err)
		}
		bill := model.BillModel{
			Id:        header.Id,
			FarmerId:  header.FarmerId,
			Date:      header.Date,
			CreateAt:  header.CreateAt,
			UpdateAt:  header.UpdateAt,
			CreateBy:  header.CreateBy,
			UpdateBy:  header.UpdateBy,
			ArrDetail: *details,
		}

		bills = append(bills, bill)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate over transaction headers: %v", err)
	}

	return &bills, nil
}

func (trxRepo *transactionRepoImpl) GetTransactionDetailsByTrxBillId(billId int) (*[]model.BillDetailModel, error) {
	qry := "SELECT id, bill_id, fertilizer_price_id ,qty, create_at, update_at, create_by, update_by FROM tr_bill_details WHERE bill_id = $1"

	rows, err := trxRepo.db.Query(qry, billId)
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction details: %v", err)
	}
	defer rows.Close()

	var details []model.BillDetailModel

	for rows.Next() {
		det := model.BillDetailModel{}
		err := rows.Scan(&det.Id, &det.BillId, &det.FertilizerPriceId, &det.Qty, &det.CreateAt, &det.UpdateAt, &det.CreateBy, &det.UpdateBy)
		if err != nil {
			return nil, fmt.Errorf("failed to scan transaction detail: %v", err)
		}
		details = append(details, det)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate over transaction details: %v", err)
	}

	return &details, nil
}

func (trxRepo *transactionRepoImpl) Create(trx *model.BillModel) error {
	tx, err := trxRepo.db.Begin()
	if err != nil {
		return fmt.Errorf("AddTransaction() Begin : %w", err)
	}

	qry := "INSERT INTO tr_bills(farmer_id, date, create_at, update_at, create_by, update_by) VALUES($1, $2, $3, $4, $5, $6) RETURNING id"

	err = tx.QueryRow(qry, trx.FarmerId, trx.Date, trx.CreateAt, trx.UpdateAt, trx.CreateBy, trx.UpdateBy).Scan(&trx.Id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("AddTransaction() Header : %w", err)
	}

	qry = "INSERT INTO tr_bill_details(bill_id, fertilizer_price_id ,qty, create_at, update_at, create_by, update_by) VALUES($1, $2, $3, $4, $5, $6, $7)"
	for _, det := range trx.ArrDetail {
		_, err := tx.Exec(qry, trx.Id, det.FertilizerPriceId, det.Qty, trx.CreateAt, trx.UpdateAt, trx.CreateBy, trx.UpdateBy)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("AddTransaction() Detail : %w", err)
		}
	}

	tx.Commit()

	return nil
}
