package usecase

import "live-code-farmer-ii/repo"

type FarmersUsecase interface {
}

type farmersUsecaseImpl struct {
	frmsRepo repo.FarmersRepo
}

func NewFarmersUsecase(frmsRepo repo.FarmersRepo) FarmersUsecase {
	return &farmersUsecaseImpl{
		frmsRepo: frmsRepo,
	}
}
