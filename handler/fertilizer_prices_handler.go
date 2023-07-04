package handler

import (
	"live-code-farmer-ii/usecase"

	"github.com/gin-gonic/gin"
)

type FertilizerPricesHandler interface {
}

type fertilizerPricesHandlerImpl struct {
	frzpUsecase usecase.FertilizerPricesUsecase
}

func NewFertilizerPricesHandler(srv *gin.Engine, frzpUsecase usecase.FertilizerPricesUsecase) FertilizerPricesHandler {
	frzpHandler := &fertilizerPricesHandlerImpl{
		frzpUsecase: frzpUsecase,
	}
	return frzpHandler
}
