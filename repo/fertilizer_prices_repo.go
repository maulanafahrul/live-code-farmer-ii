package repo

import "database/sql"

type FertilizerPricesRepo interface {
}

type fertilizerPricesRepoImpl struct {
	db *sql.DB
}

func NewFertilizerPricesRepo(db *sql.DB) FertilizerPricesRepo {
	return &fertilizerPricesRepoImpl{
		db: db,
	}
}
