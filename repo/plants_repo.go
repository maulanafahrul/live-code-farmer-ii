package repo

import "database/sql"

type PlantsRepo interface {
}

type plantsRepoImpl struct {
	db *sql.DB
}

func NewPlantsRepo(db *sql.DB) PlantsRepo {
	return &plantsRepoImpl{
		db: db,
	}
}
