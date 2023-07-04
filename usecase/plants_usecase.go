package usecase

import "live-code-farmer-ii/repo"

type PlantsUsecase interface {
}

type plantsUsecaseImpl struct {
	plnsRepo repo.PlantsRepo
}

func NewPlantsUsecase(plnsRepo repo.PlantsRepo) PlantsUsecase {
	return &plantsUsecaseImpl{
		plnsRepo: plnsRepo,
	}
}
