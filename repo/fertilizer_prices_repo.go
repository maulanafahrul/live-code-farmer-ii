package repo

import (
	"database/sql"
	"fmt"
	"live-code-farmer-ii/model"
)

type FertilizerPricesRepo interface {
	GetById(int) (*model.FertilizerPricesModel, error)
	List() (*[]model.FertilizerPricesModel, error)
	Create(*model.FertilizerPricesModel) error
	Update(*model.FertilizerPricesModel) error
	Delete(int) error
}

type fertilizerPricesRepoImpl struct {
	db *sql.DB
}

func NewFertilizerPricesRepo(db *sql.DB) FertilizerPricesRepo {
	return &fertilizerPricesRepoImpl{
		db: db,
	}
}

func (frzpRepo *fertilizerPricesRepoImpl) GetById(id int) (*model.FertilizerPricesModel, error) {
	qry := "SELECT id, fertilizer_id, price, is_active, create_at, update_at, create_by, update_by FROM ms_fertilizer_prices WHERE id = $1"

	frzp := &model.FertilizerPricesModel{}
	err := frzpRepo.db.QueryRow(qry, id).Scan(&frzp.Id, &frzp.FertilizerId, &frzp.Price, &frzp.IsActive, &frzp.CreateAt, &frzp.UpdateAt, &frzp.CreateBy, &frzp.UpdateBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on fertilizerRepoImpl.GetById() : %w", err)
	}
	return frzp, nil
}

func (frzpRepo *fertilizerPricesRepoImpl) List() (*[]model.FertilizerPricesModel, error) {
	qry := "SELECT id, fertilizer_id, price, is_active, create_at, update_at, create_by, update_by FROM ms_fertilizer_prices"
	rows, err := frzpRepo.db.Query(qry)
	if err != nil {
		return nil, fmt.Errorf("error on fertilizerPricesRepoImpl.List() : %w", err)
	}
	defer rows.Close()
	var fertilizers []model.FertilizerPricesModel
	for rows.Next() {
		ftzp := &model.FertilizerPricesModel{}
		rows.Scan(&ftzp.Id, &ftzp.FertilizerId, &ftzp.Price, &ftzp.IsActive, &ftzp.CreateAt, &ftzp.UpdateAt, &ftzp.CreateBy, &ftzp.UpdateBy)
		fertilizers = append(fertilizers, *ftzp)
	}
	return &fertilizers, nil
}

func (frzpRepo *fertilizerPricesRepoImpl) Create(frzp *model.FertilizerPricesModel) error {
	qry := "INSERT INTO ms_fertilizer_prices(fertilizer_id, price, is_Active, create_at, update_at, create_by, update_by) VALUES($1, $2, $3, $4, $5, $6, $7)"
	_, err := frzpRepo.db.Exec(qry, frzp.FertilizerId, frzp.Price, frzp.IsActive, frzp.CreateAt, frzp.UpdateAt, frzp.CreateBy, frzp.UpdateBy)
	if err != nil {
		return fmt.Errorf("error on fertilizerPricesRepoImpl.Create() : %w", err)
	}
	return nil
}

func (frzpRepo *fertilizerPricesRepoImpl) Update(frzp *model.FertilizerPricesModel) error {
	qry := "UPDATE ms_fertilizer_prices SET fertilizer_id=$1, price=$2, is_active=$3, update_at=$4, update_by=$5 WHERE id=$6"
	_, err := frzpRepo.db.Exec(qry, frzp.FertilizerId, frzp.Price, frzp.IsActive, frzp.UpdateAt, frzp.UpdateBy, frzp.Id)
	if err != nil {
		return fmt.Errorf("error on fertilizerPricesRepoImpl.Update() : %w", err)
	}
	return nil
}

func (frzRepo *fertilizerPricesRepoImpl) Delete(id int) error {
	qry := "DELETE FROM ms_fertilizer_prices WHERE id = $1"
	_, err := frzRepo.db.Exec(qry, id)
	if err != nil {
		return fmt.Errorf("error on fertilizerPricesRepoImpl.Delete() : %w", err)
	}
	return nil
}
