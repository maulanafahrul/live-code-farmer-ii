package handler

import (
	"errors"
	"fmt"
	"live-code-farmer-ii/apperror"
	"live-code-farmer-ii/model"
	"live-code-farmer-ii/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionHandler interface {
}

type transactionHandlerImpl struct {
	trxUsecase usecase.TransactionUsecase
}

func (trxHandler *transactionHandlerImpl) GetAllTransactionHandler(ctx *gin.Context) {
	trx, err := trxHandler.trxUsecase.List()
	if err != nil {
		appError := apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("trxHandler.trxUsecase.List(): %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
			return
		} else {
			fmt.Printf("trxHandler.trxUsecase.List() : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "Terjadi kesalahan ketika mengambil data transaction",
			})
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    trx,
	})
}

func (trxHandler *transactionHandlerImpl) AddTransactionHandler(ctx *gin.Context) {
	trxBill := &model.BillModel{}
	if err := ctx.ShouldBindJSON(&trxBill); err != nil {
		fmt.Printf("error parse : %v ", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}
	// validate
	if trxBill.FarmerId < 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Farmer Id tidak boleh kosong",
		})
		return
	}
	if trxBill.CreateBy == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "CreateBy tidak boleh kosong",
		})
		return
	}
	for _, det := range trxBill.ArrDetail {
		if det.FertilizerPriceId <= 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success":      false,
				"errorMessage": "FertilizerPriceId tidak boleh kosong atau minus",
			})
			return
		}
		if det.Qty <= 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success":      false,
				"errorMessage": "Qty tidak boleh kosong atau minus",
			})
			return
		}
	}
	err := trxHandler.trxUsecase.Create(trxBill)
	if err != nil {
		appError := apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("trxHandler.trxUsecase.Create() : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("trxHandler.trxUsecase.Create(): %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "Terjadi kesalahan ketika menyimpan data transaction",
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func NewTransactionHandler(srv *gin.Engine, trxUsecase usecase.TransactionUsecase) TransactionHandler {
	trxHandler := &transactionHandlerImpl{
		trxUsecase: trxUsecase,
	}
	srv.GET("/transactions", trxHandler.GetAllTransactionHandler)
	srv.POST("/transaction", trxHandler.AddTransactionHandler)
	return trxHandler
}
