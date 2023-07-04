package handler

import (
	"live-code-farmer-ii/usecase"

	"github.com/gin-gonic/gin"
)

type FertilizerHandler interface {
}

type fertilizerHandlerImpl struct {
	frzUsecase usecase.FertilizerUsecase
}

func NewFertilizerHandler(srv *gin.Engine, frzUsecase usecase.FertilizerUsecase) FertilizerHandler {
	frzHandler := &fertilizerHandlerImpl{
		frzUsecase: frzUsecase,
	}
	return frzHandler
}
