package repo

import "database/sql"

type FarmersRepo interface {
}

type farmersRepoImpl struct {
	db *sql.DB
}

func NewFarmersRepo(db *sql.DB) FarmersRepo {
	return &farmersRepoImpl{
		db: db,
	}
}
