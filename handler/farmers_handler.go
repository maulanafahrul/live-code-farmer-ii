package handler

import (
	"errors"
	"fmt"
	"live-code-farmer-ii/apperror"
	"live-code-farmer-ii/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FarmersHandler interface {
}

type farmersHandlerImpl struct {
	frmsUsecase usecase.FarmersUsecase
}

func (frmsHandler *farmersHandlerImpl) GetFarmerByIdHandler(ctx *gin.Context) {
	idText := ctx.Param("id")
	if idText == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Id tidak boleh kosong",
		})
		return
	}
	id, err := strconv.Atoi(idText)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Id harus angka",
		})
		return
	}
	frm, err := frmsHandler.frmsUsecase.Get(id)
	if err != nil {
		appError := apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("frmsHandler.frmsUsecase.Get() : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
			return
		} else {

			fmt.Printf("frmsHandler.frmsUsecase.Get() : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "Terjadi kesalahan ketika mengambil data service",
			})
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    frm,
	})
}

func (frmsHandler *farmersHandlerImpl) GetAllFarmerHandler(ctx *gin.Context) {
	fmrs, err := frmsHandler.frmsUsecase.List()
	if err != nil {
		appError := apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("frmsHandler.frmsUsecase.List(): %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
			return
		} else {

			fmt.Printf("frmsHandler.frmsUsecase.List() : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "Terjadi kesalahan ketika mengambil data service",
			})
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    fmrs,
	})
}

func NewFarmersHandler(srv *gin.Engine, frmsUsecase usecase.FarmersUsecase) FarmersHandler {
	frmsHandler := &farmersHandlerImpl{
		frmsUsecase: frmsUsecase,
	}
	srv.GET("/farmer/:id", frmsHandler.GetFarmerByIdHandler)
	srv.GET("/farmers", frmsHandler.GetAllFarmerHandler)
	return frmsHandler
}
