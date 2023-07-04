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

type FertilizerHandler interface {
}

type fertilizerHandlerImpl struct {
	frzUsecase usecase.FertilizerUsecase
}

func (frzHandler *fertilizerHandlerImpl) GetFertilizerByIdHandler(ctx *gin.Context) {
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
	pln, err := frzHandler.frzUsecase.Get(id)
	if err != nil {
		appError := apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("frzHandler.frzUsecase.Get() : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
			return
		} else {

			fmt.Printf("frzHandler.frzUsecase.Get() : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "Terjadi kesalahan ketika mengambil data fertilizer",
			})
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    pln,
	})
}

func (frzHandler *fertilizerHandlerImpl) GetAllFertilizerHandler(ctx *gin.Context) {
	frzs, err := frzHandler.frzUsecase.List()
	if err != nil {
		appError := apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("frzHandler.frzUsecase.List(): %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
			return
		} else {
			fmt.Printf("frzHandler.frzUsecase.List() : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "Terjadi kesalahan ketika mengambil data fertilizers",
			})
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    frzs,
	})
}

func (frzHandler *fertilizerHandlerImpl) AddFertilizerHandler(ctx *gin.Context) {
	payload := &model.FertilizerModel{}
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
	if payload.Stock <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Stock tidak boleh kosong",
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
	err := frzHandler.frzUsecase.Create(payload)
	if err != nil {
		appError := apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("frzHandler.frzUsecase.Create(): %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
			return
		} else {
			fmt.Printf("frzHandler.frzUsecase.Create() : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "Terjadi kesalahan ketika menyimpan data fertilizer",
			})
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (frzHandler *fertilizerHandlerImpl) UpdateFertilizerHandler(ctx *gin.Context) {
	payload := &model.FertilizerModel{}
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
	if payload.Stock <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Stock tidak boleh kosong",
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
	err := frzHandler.frzUsecase.Update(payload)
	if err != nil {
		appError := apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("frzHandler.frzUsecase.Update(): %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
			return
		} else {
			fmt.Printf("frzHandler.frzUsecase.Update() : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "Terjadi kesalahan ketika mengupdate data dertilizer",
			})
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (frzHandler *fertilizerHandlerImpl) DeleteFertilizerHandler(ctx *gin.Context) {
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
	err = frzHandler.frzUsecase.Delete(id)
	if err != nil {
		appError := apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("frzHandler.frzUsecase.Delete() : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
			return
		} else {

			fmt.Printf("frzHandler.frzUsecase.Delete() : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "Terjadi kesalahan ketika Menghapus data fertilizer",
			})
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func NewFertilizerHandler(srv *gin.Engine, frzUsecase usecase.FertilizerUsecase) FertilizerHandler {
	frzHandler := &fertilizerHandlerImpl{
		frzUsecase: frzUsecase,
	}
	srv.GET("/fertilizer/:id", frzHandler.GetFertilizerByIdHandler)
	srv.GET("/fertilizers", frzHandler.GetAllFertilizerHandler)
	srv.POST("/fertilizer", frzHandler.AddFertilizerHandler)
	srv.PUT("/fertilizer", frzHandler.UpdateFertilizerHandler)
	srv.DELETE("/fertilizer/:id", frzHandler.DeleteFertilizerHandler)
	return frzHandler
}
