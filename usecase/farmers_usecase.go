package usecase

import (
	"fmt"
	"live-code-farmer-ii/apperror"
	"live-code-farmer-ii/model"
	"live-code-farmer-ii/repo"
	"time"
)

type FarmersUsecase interface {
	Get(int) (*model.FarmersModel, error)
	List() (*[]model.FarmersModel, error)
	Create(*model.FarmersModel) error
	Update(*model.FarmersModel) error
	Delete(int) error
}

type farmersUsecaseImpl struct {
	frmsRepo repo.FarmersRepo
}

func NewFarmersUsecase(frmsRepo repo.FarmersRepo) FarmersUsecase {
	return &farmersUsecaseImpl{
		frmsRepo: frmsRepo,
	}
}

func (frmsUsecase *farmersUsecaseImpl) Get(id int) (*model.FarmersModel, error) {
	frm, err := frmsUsecase.frmsRepo.GetById(id)
	if err != nil {
		return nil, fmt.Errorf("frmsUsecase.frmsRepo.GetById() : %w", err)
	}
	if frm == nil {
		return nil, apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: fmt.Sprintf("data farmer dengan id :%d tidak ada", id),
		}
	}
	return frm, nil
}

func (frmsUsecase *farmersUsecaseImpl) List() (*[]model.FarmersModel, error) {
	frms, err := frmsUsecase.frmsRepo.List()
	if err != nil {
		return nil, fmt.Errorf("frmsUsecase.frmsRepo.List() : %w", err)
	}
	if len(*frms) == 0 {
		return nil, apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: "data farmer tidak ada",
		}
	}
	return frms, nil
}

func (frmsUsecase *farmersUsecaseImpl) Create(frm *model.FarmersModel) error {
	isNameExist, err := frmsUsecase.frmsRepo.GetByName(frm.Name)
	if err != nil {
		return fmt.Errorf("frmsUsecase.frmsRepo.GetByName() : %w", err)
	}
	if isNameExist != nil {
		return apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: fmt.Sprintf("data farmer dengan nama :%v sudah ada", frm.Name),
		}
	}
	frm.CreateAt = time.Now()
	frm.UpdateBy = "-"
	return frmsUsecase.frmsRepo.Create(frm)
}

func (frmsUsecase *farmersUsecaseImpl) Update(frm *model.FarmersModel) error {
	isIdExist, err := frmsUsecase.frmsRepo.GetById(frm.Id)
	if err != nil {
		return fmt.Errorf("frmsUsecase.frmsRepo.GetById() : %w", err)
	}
	if isIdExist == nil {
		return apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: fmt.Sprintf("data farmer dengan id :%v belum ada", frm.Id),
		}
	}

	isNameExist, err := frmsUsecase.frmsRepo.GetByName(frm.Name)
	if err != nil {
		return fmt.Errorf("frmsUsecase.frmsRepo.GetByName() : %w", err)
	}
	if isNameExist != nil {
		return apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: fmt.Sprintf("data farmer dengan nama :%v sudah ada", frm.Name),
		}
	}
	frm.UpdateAt = time.Now()
	return frmsUsecase.frmsRepo.Update(frm)

}

func (frmsUsecase *farmersUsecaseImpl) Delete(id int) error {
	frm, err := frmsUsecase.frmsRepo.GetById(id)
	if err != nil {
		return fmt.Errorf("frmsUsecase.frmsRepo.GetById() : %w", err)
	}
	if frm == nil {
		return apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: fmt.Sprintf("data farmer dengan id :%d tidak ada", id),
		}
	}
	return frmsUsecase.frmsRepo.Delete(id)
}
