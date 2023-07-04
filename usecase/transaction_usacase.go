package usecase

import (
	"fmt"
	"live-code-farmer-ii/apperror"
	"live-code-farmer-ii/model"
	"live-code-farmer-ii/repo"
	"time"
)

type TransactionUsecase interface {
	List() (*[]model.BillModel, error)
	Create(trx *model.BillModel) error
}

type transactionUsecaseImpl struct {
	trxRepo    repo.TransactionRepo
	frmplnRepo repo.FarmersPlantsRepo
	frzpRepo   repo.FertilizerPricesRepo
	frzRepo    repo.FertilizerRepo
}

func NewTransactionUsecase(trxRepo repo.TransactionRepo, frmplnRepo repo.FarmersPlantsRepo, frzpRepo repo.FertilizerPricesRepo, frzRepo repo.FertilizerRepo) TransactionUsecase {
	return &transactionUsecaseImpl{
		trxRepo:    trxRepo,
		frmplnRepo: frmplnRepo,
		frzpRepo:   frzpRepo,
		frzRepo:    frzRepo,
	}
}

func (trxUsecase *transactionUsecaseImpl) List() (*[]model.BillModel, error) {
	trx, err := trxUsecase.trxRepo.List()
	if err != nil {
		return nil, fmt.Errorf("frmsUsecase.frmsRepo.List() : %w", err)
	}
	if len(*trx) == 0 {
		return nil, apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: "data transaction tidak ada",
		}
	}
	return trx, nil
}

func (trxUsecase *transactionUsecaseImpl) Create(trx *model.BillModel) error {
	isFarmerIdExist, err := trxUsecase.frmplnRepo.GetFarmerId(trx.FarmerId)
	if err != nil {
		return fmt.Errorf("trxUsecase.frmplnRepo.GetFarmerId : %w", err)
	}
	if isFarmerIdExist == nil {
		return apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: fmt.Sprintf("data farmer dengan id :%v belum ada", trx.Id),
		}
	}
	details := []model.BillDetailModel{}
	for _, det := range trx.ArrDetail {
		isFertilizerPriceExist, err := trxUsecase.frzpRepo.GetByFertilizerId(det.FertilizerPriceId)
		if err != nil {
			return fmt.Errorf("trxUsecase.frmplnRepo.GetFarmerId : %w", err)
		}
		if isFertilizerPriceExist == nil {
			return apperror.AppError{
				ErrorCode:    400,
				ErrorMassage: fmt.Sprintf("data fertilizer price dengan id :%v belum ada", det.FertilizerPriceId),
			}
		}
		detail := model.BillDetailModel{
			BillId:            det.BillId,
			FertilizerPriceId: det.FertilizerPriceId,
			Qty:               det.Qty,
		}
		fertilizerId, err := trxUsecase.frzpRepo.GetById(det.FertilizerPriceId)
		if err != nil {
			return fmt.Errorf("error on trxUsecase.frzpRepo.GetById : %w", err)
		}
		err = trxUsecase.frzRepo.ReduceQty(fertilizerId.FertilizerId, det.Qty)
		if err != nil {
			return fmt.Errorf("trxUsecase.frzRepo.ReduceQty : %w", err)
		}
		details = append(details, detail)
	}
	trx.ArrDetail = details
	trx.Date = time.Now()
	trx.CreateAt = time.Now()
	trx.UpdateBy = "-"
	return trxUsecase.trxRepo.Create(trx)

}
