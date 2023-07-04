package usecase

import "live-code-farmer-ii/repo"

type FertilizerPricesUsecase interface {
}

type fertilizerPricesUsecaseImpl struct {
	frzpRepo repo.FertilizerPricesRepo
}

func NewFertilizerPricesUsecase(frzpRepo repo.FertilizerPricesRepo) FertilizerPricesUsecase {
	return &fertilizerPricesUsecaseImpl{
		frzpRepo: frzpRepo,
	}
}
