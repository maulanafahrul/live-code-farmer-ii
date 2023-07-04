package usecase

import (
	"fmt"
	"live-code-farmer-ii/apperror"
	"live-code-farmer-ii/model"
	"live-code-farmer-ii/repo"
)

type FarmersUsecase interface {
	Get(int) (*model.FarmersModel, error)
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
