package usecase

import "live-code-farmer-ii/repo"

type FertilizerUsecase interface {
}

type fertilizerUsecaseImpl struct {
	frzRepo repo.FertilizerRepo
}

func NewFertilizerUsecase(frzRepo repo.FertilizerRepo) FertilizerUsecase {
	return &fertilizerUsecaseImpl{
		frzRepo: frzRepo,
	}
}
