package usecase

import (
	"fmt"
	"live-code-farmer-ii/apperror"
	"live-code-farmer-ii/model"
	"live-code-farmer-ii/repo"
	"time"
)

type FertilizerPricesUsecase interface {
	Get(int) (*model.FertilizerPricesModel, error)
	List() (*[]model.FertilizerPricesModel, error)
	Create(*model.FertilizerPricesModel) error
	Update(*model.FertilizerPricesModel) error
	Delete(int) error
}

type fertilizerPricesUsecaseImpl struct {
	frzpRepo repo.FertilizerPricesRepo
	frzRepo  repo.FertilizerRepo
}

func NewFertilizerPricesUsecase(frzpRepo repo.FertilizerPricesRepo, frzRepo repo.FertilizerRepo) FertilizerPricesUsecase {
	return &fertilizerPricesUsecaseImpl{
		frzpRepo: frzpRepo,
		frzRepo:  frzRepo,
	}
}

func (frzpUsecase *fertilizerPricesUsecaseImpl) Get(id int) (*model.FertilizerPricesModel, error) {
	frz, err := frzpUsecase.frzpRepo.GetById(id)
	if err != nil {
		return nil, fmt.Errorf("frzpUsecase.frzpRepo.GetById() : %w", err)
	}
	if frz == nil {
		return nil, apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: fmt.Sprintf("data fertilizer price dengan id :%d tidak ada", id),
		}
	}
	return frz, nil
}

func (frzpUsecase *fertilizerPricesUsecaseImpl) List() (*[]model.FertilizerPricesModel, error) {
	frzps, err := frzpUsecase.frzpRepo.List()
	if err != nil {
		return nil, fmt.Errorf("frzUsecase.frzRepo.List() : %w", err)
	}
	if len(*frzps) == 0 {
		return nil, apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: "data fertilizers tidak ada",
		}
	}
	return frzps, nil
}

func (frzpUsecase *fertilizerPricesUsecaseImpl) Create(frzp *model.FertilizerPricesModel) error {
	fertilizer, err := frzpUsecase.frzRepo.GetById(frzp.FertilizerId)
	if err != nil {
		return fmt.Errorf("frzpUsecase.frzRepo.GetById() : %w", err)
	}
	if fertilizer == nil {
		return apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: fmt.Sprintf("data fertilizer dengan id :%d tidak ada", frzp.FertilizerId),
		}
	}
	frzp.CreateAt = time.Now()
	frzp.UpdateBy = "-"
	return frzpUsecase.frzpRepo.Create(frzp)
}

func (frzpUsecase *fertilizerPricesUsecaseImpl) Update(frzp *model.FertilizerPricesModel) error {
	isIdExist, err := frzpUsecase.frzpRepo.GetById(frzp.Id)
	if err != nil {
		return fmt.Errorf("frzpUsecase.frzpRepo.GetById() : %w", err)
	}
	if isIdExist == nil {
		return apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: fmt.Sprintf("data fertilizer price dengan id :%v belum ada", frzp.Id),
		}
	}
	fertilizer, err := frzpUsecase.frzRepo.GetById(frzp.FertilizerId)
	if err != nil {
		return fmt.Errorf("frzpUsecase.frzRepo.GetById() : %w", err)
	}
	if fertilizer == nil {
		return apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: fmt.Sprintf("data fertilizer dengan id :%d tidak ada", frzp.FertilizerId),
		}
	}
	frzp.UpdateAt = time.Now()
	return frzpUsecase.frzpRepo.Update(frzp)
}

func (frzpUsecase *fertilizerPricesUsecaseImpl) Delete(id int) error {
	frzp, err := frzpUsecase.frzpRepo.GetById(id)
	if err != nil {
		return fmt.Errorf("frzpUsecase.frzpRepo() : %w", err)
	}
	if frzp == nil {
		return apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: fmt.Sprintf("data fertilizer price dengan id :%d tidak ada", id),
		}
	}
	return frzpUsecase.frzpRepo.Delete(id)
}
