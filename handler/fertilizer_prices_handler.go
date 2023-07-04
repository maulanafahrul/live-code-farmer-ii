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

type FertilizerPricesHandler interface {
}

type fertilizerPricesHandlerImpl struct {
	frzpUsecase usecase.FertilizerPricesUsecase
}

func (frzpHandler *fertilizerPricesHandlerImpl) GetFertilizerPriceByIdHandler(ctx *gin.Context) {
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
	frzp, err := frzpHandler.frzpUsecase.Get(id)
	if err != nil {
		appError := apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("frzpHandler.frzpUsecase.Get() : %v ", err.Error())
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
		"data":    frzp,
	})
}

func (frzpHandler *fertilizerPricesHandlerImpl) GetAllFertilizerPriceHandler(ctx *gin.Context) {
	frzps, err := frzpHandler.frzpUsecase.List()
	if err != nil {
		appError := apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("frzpHandler.frzpUsecase.List(): %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
			return
		} else {
			fmt.Printf("frzpHandler.frzpUsecase.List() : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "Terjadi kesalahan ketika mengambil data fertilizer price",
			})
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    frzps,
	})
}

func (frzpHandler *fertilizerPricesHandlerImpl) AddFertilizerPriceHandler(ctx *gin.Context) {
	payload := &model.FertilizerPricesModel{}
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}
	// validate
	if payload.FertilizerId <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Fertilizer id tidak boleh kosong atau minus",
		})
		return
	}
	if payload.Price <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Price tidak boleh kosong atau minus",
		})
		return
	}
	if payload.CreateBy == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "CreateBy tidak boleh kosong",
		})
		return
	}
	if len(payload.CreateBy) > 20 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "CreateBy tidak boleh lebih dari 20",
		})
		return
	}
	err := frzpHandler.frzpUsecase.Create(payload)
	if err != nil {
		appError := apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("frzpHandler.frzpUsecase.Create(): %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
			return
		} else {
			fmt.Printf("frzpHandler.frzpUsecase.Create() : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "Terjadi kesalahan ketika menyimpan data fertilizer price",
			})
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (frzpHandler *fertilizerPricesHandlerImpl) UpdateFertilizerPriceHandler(ctx *gin.Context) {
	payload := &model.FertilizerPricesModel{}
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
			"errorMessage": "Id tidak boleh kosong atau minus",
		})
		return
	}
	if payload.Price <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Price tidak boleh kosong atau minus",
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
	err := frzpHandler.frzpUsecase.Update(payload)
	if err != nil {
		appError := apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("frzpHandler.frzpUsecase.Update(): %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
			return
		} else {
			fmt.Printf("frzpHandler.frzpUsecase.Update() : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "Terjadi kesalahan ketika mengupdate data fertilizer price",
			})
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (frzpHandler *fertilizerPricesHandlerImpl) DeleteFertilizerPriceHandler(ctx *gin.Context) {
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
	err = frzpHandler.frzpUsecase.Delete(id)
	if err != nil {
		appError := apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("frzpHandler.frzpUsecase.Delete() : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
			return
		} else {

			fmt.Printf("frzpHandler.frzpUsecase.Delete() : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "Terjadi kesalahan ketika Menghapus data fertilizer price",
			})
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func NewFertilizerPricesHandler(srv *gin.Engine, frzpUsecase usecase.FertilizerPricesUsecase) FertilizerPricesHandler {
	frzpHandler := &fertilizerPricesHandlerImpl{
		frzpUsecase: frzpUsecase,
	}
	srv.GET("/fertilizerprice/:id", frzpHandler.GetFertilizerPriceByIdHandler)
	srv.GET("/fertilizerprices", frzpHandler.GetAllFertilizerPriceHandler)
	srv.POST("/fertilizerprice", frzpHandler.AddFertilizerPriceHandler)
	srv.PUT("/fertilizerprice", frzpHandler.UpdateFertilizerPriceHandler)
	srv.DELETE("/fertilizerprice/:id", frzpHandler.DeleteFertilizerPriceHandler)
	return frzpHandler
}
