package handler

import (
	"live-code-farmer-ii/usecase"

	"github.com/gin-gonic/gin"
)

type TransactionHandler interface {
}

type transactionHandlerImpl struct {
	trxUsecase usecase.TransactionUsecase
}

func NewTransactionHandler(srv *gin.Engine, trxUsecase usecase.TransactionUsecase) TransactionHandler {
	trxHandler := &transactionHandlerImpl{
		trxUsecase: trxUsecase,
	}
	return trxHandler
}
