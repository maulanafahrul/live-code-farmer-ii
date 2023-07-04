package repo

import "database/sql"

type TransactionRepo interface {
}

type transactionRepoImpl struct {
	db *sql.DB
}

func NewTransactionRepo(db *sql.DB) TransactionRepo {
	return &transactionRepoImpl{
		db: db,
	}
}
