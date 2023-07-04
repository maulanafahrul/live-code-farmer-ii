package repo

import "database/sql"

type FertilizerRepo interface {
}

type fertilizerRepoImpl struct {
	db *sql.DB
}

func NewFertilizerRepo(db *sql.DB) FertilizerRepo {
	return &fertilizerRepoImpl{
		db: db,
	}
}
