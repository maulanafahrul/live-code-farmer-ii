package manager

import (
	"live-code-farmer-ii/repo"
	"sync"
)

type RepoManager interface {
	GetTransactionRepo() repo.TransactionRepo
	GetFarmersRepo() repo.FarmersRepo
	GetPlantsRepo() repo.PlantsRepo
	GetFertilizerRepo() repo.FertilizerRepo
	GetFertilizerPricesRepo() repo.FertilizerPricesRepo
}

type repoManager struct {
	infraManager InfraManager
	frmsRepo     repo.FarmersRepo
	plnsRepo     repo.PlantsRepo
	frzRepo      repo.FertilizerRepo
	frzpRepo     repo.FertilizerPricesRepo
	trxRepo      repo.TransactionRepo
}

var onceLoadFarmersRepo sync.Once
var onceLoadPlantsRepo sync.Once
var onceLoadFertilizerRepo sync.Once
var onceLoadFertilizerPricesRepo sync.Once
var onceLoadTransactionRepo sync.Once

func (rm *repoManager) GetTransactionRepo() repo.TransactionRepo {
	onceLoadTransactionRepo.Do(func() {
		rm.trxRepo = repo.NewTransactionRepo(rm.infraManager.GetDB())
	})
	return rm.trxRepo
}
func (rm *repoManager) GetFarmersRepo() repo.FarmersRepo {
	onceLoadFarmersRepo.Do(func() {
		rm.frmsRepo = repo.NewFarmersRepo(rm.infraManager.GetDB())
	})
	return rm.frmsRepo
}
func (rm *repoManager) GetPlantsRepo() repo.PlantsRepo {
	onceLoadPlantsRepo.Do(func() {
		rm.plnsRepo = repo.NewPlantsRepo(rm.infraManager.GetDB())
	})
	return rm.plnsRepo
}
func (rm *repoManager) GetFertilizerRepo() repo.FertilizerRepo {
	onceLoadFertilizerRepo.Do(func() {
		rm.frzRepo = repo.NewFertilizerRepo(rm.infraManager.GetDB())
	})
	return rm.frzRepo
}
func (rm *repoManager) GetFertilizerPricesRepo() repo.FertilizerPricesRepo {
	onceLoadFertilizerPricesRepo.Do(func() {
		rm.frzpRepo = repo.NewFertilizerPricesRepo(rm.infraManager.GetDB())
	})
	return rm.frzpRepo
}

func NewRepoManager(infraManager InfraManager) RepoManager {
	return &repoManager{
		infraManager: infraManager,
	}
}
