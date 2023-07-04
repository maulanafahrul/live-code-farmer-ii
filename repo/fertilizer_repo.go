package repo

import (
	"database/sql"
	"fmt"
	"live-code-farmer-ii/model"
)

type FertilizerRepo interface {
	GetById(int) (*model.FertilizerModel, error)
	GetByName(string) (*model.FertilizerModel, error)
	List() (*[]model.FertilizerModel, error)
	Create(*model.FertilizerModel) error
	Update(*model.FertilizerModel) error
	Delete(int) error
	ReduceQty(int, int) error
}

type fertilizerRepoImpl struct {
	db *sql.DB
}

func NewFertilizerRepo(db *sql.DB) FertilizerRepo {
	return &fertilizerRepoImpl{
		db: db,
	}
}

func (frzRepo *fertilizerRepoImpl) GetById(id int) (*model.FertilizerModel, error) {
	qry := "SELECT id, name, stock, create_at, update_at, create_by, update_by FROM ms_fertilizers WHERE id = $1"

	frz := &model.FertilizerModel{}
	err := frzRepo.db.QueryRow(qry, id).Scan(&frz.Id, &frz.Name, &frz.Stock, &frz.CreateAt, &frz.UpdateAt, &frz.CreateBy, &frz.UpdateBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on fertilizerRepoImpl.GetById() : %w", err)
	}
	return frz, nil
}

func (frzRepo *fertilizerRepoImpl) GetByName(name string) (*model.FertilizerModel, error) {
	qry := "SELECT id, name, stock, create_at, update_at, create_by, update_by FROM ms_fertilizers WHERE name = $1"

	frz := &model.FertilizerModel{}
	err := frzRepo.db.QueryRow(qry, name).Scan(&frz.Id, &frz.Name, &frz.Stock, &frz.CreateAt, &frz.UpdateAt, &frz.CreateBy, &frz.UpdateBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on fertilizerRepoImpl.GetByName() : %w", err)
	}
	return frz, nil
}

func (frzRepo *fertilizerRepoImpl) List() (*[]model.FertilizerModel, error) {
	qry := "SELECT id, name, stock, create_at, update_at, create_by, update_by FROM ms_fertilizers"
	rows, err := frzRepo.db.Query(qry)
	if err != nil {
		return nil, fmt.Errorf("error on fertilizerRepoImpl.List() : %w", err)
	}
	defer rows.Close()
	var fertilizers []model.FertilizerModel
	for rows.Next() {
		ftz := &model.FertilizerModel{}
		rows.Scan(&ftz.Id, &ftz.Name, &ftz.Stock, &ftz.CreateAt, &ftz.UpdateAt, &ftz.CreateBy, &ftz.UpdateBy)
		fertilizers = append(fertilizers, *ftz)
	}
	return &fertilizers, nil
}

func (frzRepo *fertilizerRepoImpl) Create(frz *model.FertilizerModel) error {
	qry := "INSERT INTO ms_fertilizers(name, stock,create_at, update_at, create_by, update_by) VALUES($1, $2, $3, $4, $5,$6)"
	_, err := frzRepo.db.Exec(qry, frz.Name, frz.Stock, frz.CreateAt, frz.UpdateAt, frz.CreateBy, frz.UpdateBy)
	if err != nil {
		return fmt.Errorf("error on fertilizerRepoImpl.Create() : %w", err)
	}
	return nil
}

func (frzRepo *fertilizerRepoImpl) Update(frz *model.FertilizerModel) error {
	qry := "UPDATE ms_fertilizers SET name=$1, stock=$2, update_at=$3, update_by=$4 WHERE id=$5"
	_, err := frzRepo.db.Exec(qry, frz.Name, frz.Stock, frz.UpdateAt, frz.UpdateBy, frz.Id)
	if err != nil {
		return fmt.Errorf("error on fertilizerRepoImpl.Update() : %w", err)
	}
	return nil
}

func (frzRepo *fertilizerRepoImpl) Delete(id int) error {
	qry := "DELETE FROM ms_fertilizers WHERE id = $1"
	_, err := frzRepo.db.Exec(qry, id)
	if err != nil {
		return fmt.Errorf("error on fertilizerRepoImpl.Delete() : %w", err)
	}
	return nil
}

func (frzRepo *fertilizerRepoImpl) ReduceQty(id int, qty int) error {
	qry := "UPDATE ms_fertilizers SET stock=stock - $1 WHERE id = $2"
	_, err := frzRepo.db.Exec(qry, qty, id)
	if err != nil {
		return fmt.Errorf("error on fertilizerRepoImpl.ReduceQty() : %w", err)
	}
	return nil
}
