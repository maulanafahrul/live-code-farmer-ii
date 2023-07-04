package repo

import (
	"database/sql"
	"fmt"
	"live-code-farmer-ii/model"
)

type FarmersPlantsRepo interface {
	GetFarmerId(int) (*model.FarmersPlaintsModel, error)
}

type farmersPlantsRepo struct {
	db *sql.DB
}

func NewFarmersPlantsRepo(db *sql.DB) FarmersPlantsRepo {
	return &farmersPlantsRepo{
		db: db,
	}
}

func (frmplnRepo *farmersPlantsRepo) GetFarmerId(id int) (*model.FarmersPlaintsModel, error) {
	qry := "SELECT farmer_id FROM farmers_plants WHERE farmer_id = $1"

	frmpln := &model.FarmersPlaintsModel{}
	err := frmplnRepo.db.QueryRow(qry, id).Scan(&frmpln.FarmerId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on farmersRepoImpl.GetById() : %w", err)
	}
	return frmpln, nil
}
