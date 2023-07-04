package handler

import (
	"live-code-farmer-ii/usecase"

	"github.com/gin-gonic/gin"
)

type PlantsHandler interface {
}

type plantsHandlerImpl struct {
	plnsUsecase usecase.PlantsUsecase
}

func NewPlantsHandler(srv *gin.Engine, plnsUsecase usecase.PlantsUsecase) PlantsHandler {
	plnsHandler := &plantsHandlerImpl{
		plnsUsecase: plnsUsecase,
	}
	return plnsHandler
}
