package usecase

import "live-code-farmer-ii/repo"

type TransactionUsecase interface {
}

type transactionUsecaseImpl struct {
	trxRepo repo.TransactionRepo
}

func NewTransactionUsecase(trxRepo repo.TransactionRepo) TransactionUsecase {
	return &transactionUsecaseImpl{
		trxRepo: trxRepo,
	}
}
