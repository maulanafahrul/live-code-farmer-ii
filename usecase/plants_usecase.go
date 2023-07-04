package usecase

import (
	"fmt"
	"live-code-farmer-ii/apperror"
	"live-code-farmer-ii/model"
	"live-code-farmer-ii/repo"
	"time"
)

type PlantsUsecase interface {
	Get(int) (*model.PlantModel, error)
	List() (*[]model.PlantModel, error)
	Create(*model.PlantModel) error
	Update(*model.PlantModel) error
	Delete(int) error
}

type plantsUsecaseImpl struct {
	plnsRepo repo.PlantsRepo
}

func NewPlantsUsecase(plnsRepo repo.PlantsRepo) PlantsUsecase {
	return &plantsUsecaseImpl{
		plnsRepo: plnsRepo,
	}
}

func (plnsUsecase *plantsUsecaseImpl) Get(id int) (*model.PlantModel, error) {
	pln, err := plnsUsecase.plnsRepo.GetById(id)
	if err != nil {
		return nil, fmt.Errorf("plnsUsecase.plnsRepo.GetById() : %w", err)
	}
	if pln == nil {
		return nil, apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: fmt.Sprintf("data plant dengan id :%d tidak ada", id),
		}
	}
	return pln, nil
}
func (plnsUsecase *plantsUsecaseImpl) List() (*[]model.PlantModel, error) {
	plns, err := plnsUsecase.plnsRepo.List()
	if err != nil {
		return nil, fmt.Errorf("plnsUsecase.plnsRepo.List() : %w", err)
	}
	if len(*plns) == 0 {
		return nil, apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: "data plants tidak ada",
		}
	}
	return plns, nil
}

func (plnsUsecase *plantsUsecaseImpl) Create(pln *model.PlantModel) error {
	isNameExist, err := plnsUsecase.plnsRepo.GetByName(pln.Name)
	if err != nil {
		return fmt.Errorf("plnsUsecase.plnsRepo.GetByName() : %w", err)
	}
	if isNameExist != nil {
		return apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: fmt.Sprintf("data plant dengan nama :%v sudah ada", pln.Name),
		}
	}
	pln.CreateAt = time.Now()
	pln.UpdateBy = "-"
	return plnsUsecase.plnsRepo.Create(pln)
}

func (plnsUsecase *plantsUsecaseImpl) Update(pln *model.PlantModel) error {
	isIdExist, err := plnsUsecase.plnsRepo.GetById(pln.Id)
	if err != nil {
		return fmt.Errorf("plnsUsecase.plnsRepo.GetById() : %w", err)
	}
	if isIdExist == nil {
		return apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: fmt.Sprintf("data plant dengan id :%v belum ada", pln.Id),
		}
	}

	isNameExist, err := plnsUsecase.plnsRepo.GetByName(pln.Name)
	if err != nil {
		return fmt.Errorf("plnsUsecase.plnsRepo.GetByName() : %w", err)
	}
	if isNameExist != nil {
		return apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: fmt.Sprintf("data plant dengan nama :%v sudah ada", pln.Name),
		}
	}
	pln.UpdateAt = time.Now()
	return plnsUsecase.plnsRepo.Update(pln)

}

func (plnsUsecase *plantsUsecaseImpl) Delete(id int) error {
	pln, err := plnsUsecase.plnsRepo.GetById(id)
	if err != nil {
		return fmt.Errorf("plnsUsecase.plnsRepo.GetById() : %w", err)
	}
	if pln == nil {
		return apperror.AppError{
			ErrorCode:    400,
			ErrorMassage: fmt.Sprintf("data plant dengan id :%d tidak ada", id),
		}
	}
	return plnsUsecase.plnsRepo.Delete(id)
}
