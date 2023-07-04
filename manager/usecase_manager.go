package manager

import (
	"live-code-farmer-ii/usecase"
	"sync"
)

type UsecaseManager interface {
	GetFarmersUsecase() usecase.FarmersUsecase
	GetPlantsUsecase() usecase.PlantsUsecase
	GetFertilizerUsecase() usecase.FertilizerUsecase
	GetFertilizerPricesUsecase() usecase.FertilizerPricesUsecase
	GetTransactionUsecase() usecase.TransactionUsecase
}

type usecaseManager struct {
	repoManager RepoManager
	frmsUsecase usecase.FarmersUsecase
	plnsUsecase usecase.PlantsUsecase
	frzUsecase  usecase.FertilizerUsecase
	frzpUsecase usecase.FertilizerPricesUsecase
	trxUsecase  usecase.TransactionUsecase
}

var onceLoadFarmersUsecase sync.Once
var onceLoadPlantsUsecase sync.Once
var onceLoadFertilizerUsecase sync.Once
var onceLoadFertilizerPricesUsecase sync.Once
var onceLoadTransactionUsecase sync.Once

func (um *usecaseManager) GetFarmersUsecase() usecase.FarmersUsecase {
	onceLoadFarmersUsecase.Do(func() {
		um.frmsUsecase = usecase.NewFarmersUsecase(um.repoManager.GetFarmersRepo())
	})
	return um.frmsUsecase
}
func (um *usecaseManager) GetPlantsUsecase() usecase.PlantsUsecase {
	onceLoadPlantsUsecase.Do(func() {
		um.plnsUsecase = usecase.NewPlantsUsecase(um.repoManager.GetPlantsRepo())
	})
	return um.plnsUsecase
}
func (um *usecaseManager) GetFertilizerUsecase() usecase.FertilizerUsecase {
	onceLoadFertilizerUsecase.Do(func() {
		um.frzUsecase = usecase.NewFertilizerUsecase(um.repoManager.GetFertilizerRepo())
	})
	return um.frzUsecase
}
func (um *usecaseManager) GetFertilizerPricesUsecase() usecase.FertilizerPricesUsecase {
	onceLoadFertilizerPricesUsecase.Do(func() {
		um.frzpUsecase = usecase.NewFertilizerPricesUsecase(um.repoManager.GetFertilizerPricesRepo(), um.repoManager.GetFertilizerRepo())
	})
	return um.frzpUsecase
}
func (um *usecaseManager) GetTransactionUsecase() usecase.TransactionUsecase {
	onceLoadTransactionUsecase.Do(func() {
		um.trxUsecase = usecase.NewTransactionUsecase(um.repoManager.GetTransactionRepo(), um.repoManager.GetFarmersPlantsRepo(), um.repoManager.GetFertilizerPricesRepo(), um.repoManager.GetFertilizerRepo())
	})
	return um.trxUsecase
}

func NewUsecaseManager(repoManager RepoManager) UsecaseManager {
	return &usecaseManager{
		repoManager: repoManager,
	}
}
