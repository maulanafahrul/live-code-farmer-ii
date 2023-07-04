package handler

import (
	"live-code-farmer-ii/usecase"

	"github.com/gin-gonic/gin"
)

type FarmersHandler interface {
}

type farmersHandlerImpl struct {
	frmsUsecase usecase.FarmersUsecase
}

func NewFarmersHandler(srv *gin.Engine, frmsUsecase usecase.FarmersUsecase) FarmersHandler {
	frmsHandler := &farmersHandlerImpl{
		frmsUsecase: frmsUsecase,
	}
	return frmsHandler
}
