package usecase

import (
	"fmt"
	"live-code-farmer-ii/apperror"
	"live-code-farmer-ii/model"
	"live-code-farmer-ii/repo"
	"time"
)

type FertilizerUsecase interface {
	Get(int) (*model.FertilizerModel, error)
	List() (*[]model.FertilizerModel, error)
	Create(*model.FertilizerModel) error
	Update(*model.FertilizerModel) error
	Delete(int) error
}

type fertilizerUsecaseImpl struct {
	frzRepo repo.FertilizerRepo
}

func NewFertilizerUsecase(frzRepo repo.FertilizerRepo) FertilizerUsecase {
	return &fertilizerUsecaseImpl{
		frzRepo: frzRepo,
	}
}

func (frzpUsecase *fertilizerUsecaseImpl) Get(id int) (*model.FertilizerModel, error) {
	frz, err := frzpUsecase.frzRepo.GetById(id)
	if err != nil {
		return nil, fmt.Errorf("frzpUsecase.frzRepo.GetById() : %w", err)
	}
	if frz == nil {
		return nil, apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: fmt.Sprintf("data plant dengan id :%d tidak ada", id),
		}
	}
	return frz, nil
}

func (frzpUsecase *fertilizerUsecaseImpl) List() (*[]model.FertilizerModel, error) {
	frzs, err := frzpUsecase.frzRepo.List()
	if err != nil {
		return nil, fmt.Errorf("frzpUsecase.frzRepo.List() : %w", err)
	}
	if len(*frzs) == 0 {
		return nil, apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: "data plants tidak ada",
		}
	}
	return frzs, nil
}

func (frzpUsecase *fertilizerUsecaseImpl) Create(frz *model.FertilizerModel) error {
	isNameExist, err := frzpUsecase.frzRepo.GetByName(frz.Name)
	if err != nil {
		return fmt.Errorf("frzpUsecase.frzRepo.GetByName() : %w", err)
	}
	if isNameExist != nil {
		return apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: fmt.Sprintf("data fertilizer dengan nama :%v sudah ada", frz.Name),
		}
	}
	frz.CreateAt = time.Now()
	frz.UpdateBy = "-"
	return frzpUsecase.frzRepo.Create(frz)
}

func (frzpUsecase *fertilizerUsecaseImpl) Update(frz *model.FertilizerModel) error {
	isIdExist, err := frzpUsecase.frzRepo.GetById(frz.Id)
	if err != nil {
		return fmt.Errorf("frzpUsecase.frzRepo.GetById() : %w", err)
	}
	if isIdExist == nil {
		return apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: fmt.Sprintf("data fertilizer dengan id :%v belum ada", frz.Id),
		}
	}

	isNameExist, err := frzpUsecase.frzRepo.GetByName(frz.Name)
	if err != nil {
		return fmt.Errorf("frzpUsecase.frzRepo.GetByName() : %w", err)
	}
	if isNameExist != nil {
		return apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: fmt.Sprintf("data fertilizer dengan nama :%v sudah ada", frz.Name),
		}
	}
	frz.UpdateAt = time.Now()
	return frzpUsecase.frzRepo.Update(frz)
}

func (frzpUsecase *fertilizerUsecaseImpl) Delete(id int) error {
	frz, err := frzpUsecase.frzRepo.GetById(id)
	if err != nil {
		return fmt.Errorf("frzpUsecase.frzRepo.GetById() : %w", err)
	}
	if frz == nil {
		return apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: fmt.Sprintf("data plant dengan id :%d tidak ada", id),
		}
	}
	return frzpUsecase.frzRepo.Delete(id)
}
