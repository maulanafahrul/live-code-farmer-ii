package handler

import (
	"errors"
	"fmt"
	"live-code-farmer-ii/apperror"
	"live-code-farmer-ii/model"
	"live-code-farmer-ii/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PlantsHandler interface {
}

type plantsHandlerImpl struct {
	plnsUsecase usecase.PlantsUsecase
}

func (plnsHandler *plantsHandlerImpl) GetPlantByIdHandler(ctx *gin.Context) {
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
	pln, err := plnsHandler.plnsUsecase.Get(id)
	if err != nil {
		appError := apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("plnsHandler.plnsUsecase.Get() : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
			return
		} else {

			fmt.Printf("plnsHandler.plnsUsecase.Get() : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "Terjadi kesalahan ketika mengambil data farmer",
			})
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    pln,
	})
}

func (plnsHandler *plantsHandlerImpl) GetAllPlantsHandler(ctx *gin.Context) {
	plns, err := plnsHandler.plnsUsecase.List()
	if err != nil {
		appError := apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("plnsHandler.plnsUsecase.List(): %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
			return
		} else {
			fmt.Printf("plnsHandler.plnsUsecase.List() : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "Terjadi kesalahan ketika mengambil data plants",
			})
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    plns,
	})
}

func (plnsHandler *plantsHandlerImpl) AddPlantsHandler(ctx *gin.Context) {
	payload := &model.PlantModel{}
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}
	// validate
	if payload.Name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Nama tidak boleh kosong",
		})
		return
	}
	if len(payload.Name) > 20 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Nama tidak boleh lebih dari 20",
		})
		return
	}
	if payload.CreateBy == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "createby tidak boleh kosong",
		})
		return
	}
	if len(payload.CreateBy) > 20 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "createby tidak boleh lebih dari 20",
		})
		return
	}
	err := plnsHandler.plnsUsecase.Create(payload)
	if err != nil {
		appError := apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("plnsHandler.plnsUsecase.Create(): %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
			return
		} else {
			fmt.Printf("plnsHandler.plnsUsecase.Create() : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "Terjadi kesalahan ketika menyimpan data plant",
			})
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (plnsHandler *plantsHandlerImpl) UpdatePlantsHandler(ctx *gin.Context) {
	payload := &model.PlantModel{}
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}
	// validate
	if payload.Id <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Id tidak boleh kosong",
		})
		return
	}
	if payload.Name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Nama tidak boleh kosong",
		})
		return
	}
	if len(payload.Name) > 20 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Nama tidak boleh lebih dari 20",
		})
		return
	}
	if payload.UpdateBy == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "UpdateBy tidak boleh kosong",
		})
		return
	}
	if len(payload.UpdateBy) > 20 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "UpdateBy tidak boleh lebih dari 20",
		})
		return
	}
	err := plnsHandler.plnsUsecase.Update(payload)
	if err != nil {
		appError := apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("plnsHandler.plnsUsecase.Update(): %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
			return
		} else {
			fmt.Printf("plnsHandler.plnsUsecase.Update() : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "Terjadi kesalahan ketika mengupdate data plant",
			})
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
func (plnsHandler *plantsHandlerImpl) DeletePlantsHandler(ctx *gin.Context) {
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
	err = plnsHandler.plnsUsecase.Delete(id)
	if err != nil {
		appError := apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("plnsHandler.plnsUsecase.Delete() : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
			return
		} else {

			fmt.Printf("plnsHandler.plnsUsecase.Delete() : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "Terjadi kesalahan ketika Menghapus data plant",
			})
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
func NewPlantsHandler(srv *gin.Engine, plnsUsecase usecase.PlantsUsecase) PlantsHandler {
	plnsHandler := &plantsHandlerImpl{
		plnsUsecase: plnsUsecase,
	}
	srv.GET("/plant/:id", plnsHandler.GetPlantByIdHandler)
	srv.GET("/plants", plnsHandler.GetAllPlantsHandler)
	srv.POST("/plant", plnsHandler.AddPlantsHandler)
	srv.PUT("/plant", plnsHandler.UpdatePlantsHandler)
	srv.DELETE("/plant/:id", plnsHandler.DeletePlantsHandler)
	return plnsHandler
}
