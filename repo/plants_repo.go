package repo

import (
	"database/sql"
	"fmt"
	"live-code-farmer-ii/model"
)

type PlantsRepo interface {
	GetById(int) (*model.PlantModel, error)
	GetByName(string) (*model.PlantModel, error)
	List() (*[]model.PlantModel, error)
	Create(*model.PlantModel) error
	Update(*model.PlantModel) error
	Delete(int) error
}

type plantsRepoImpl struct {
	db *sql.DB
}

func NewPlantsRepo(db *sql.DB) PlantsRepo {
	return &plantsRepoImpl{
		db: db,
	}
}

func (plnsRepo *plantsRepoImpl) GetById(id int) (*model.PlantModel, error) {
	qry := "SELECT id, name, create_at, update_at, create_by, update_by FROM ms_plants WHERE id = $1"

	pln := &model.PlantModel{}
	err := plnsRepo.db.QueryRow(qry, id).Scan(&pln.Id, &pln.Name, &pln.CreateAt, &pln.UpdateAt, &pln.CreateBy, &pln.UpdateBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on plantsRepoImpl.GetById() : %w", err)
	}
	return pln, nil
}

func (plnsRepo *plantsRepoImpl) GetByName(name string) (*model.PlantModel, error) {
	qry := "SELECT id, name, create_at, update_at, create_by, update_by FROM ms_plants WHERE name = $1"

	pln := &model.PlantModel{}
	err := plnsRepo.db.QueryRow(qry, name).Scan(&pln.Id, &pln.Name, &pln.CreateAt, &pln.UpdateAt, &pln.CreateBy, &pln.UpdateBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on plantsRepoImpl.GetByName() : %w", err)
	}
	return pln, nil
}

func (plnsRepo *plantsRepoImpl) List() (*[]model.PlantModel, error) {
	qry := "SELECT id, name, create_at, update_at, create_by, update_by FROM ms_plants"
	rows, err := plnsRepo.db.Query(qry)
	if err != nil {
		return nil, fmt.Errorf("error on plantsRepoImpl.List() : %w", err)
	}
	defer rows.Close()
	var plants []model.PlantModel
	for rows.Next() {
		pln := &model.PlantModel{}
		rows.Scan(&pln.Id, &pln.Name, &pln.CreateAt, &pln.UpdateAt, &pln.CreateBy, &pln.UpdateBy)
		plants = append(plants, *pln)
	}
	return &plants, nil
}

func (plnsRepo *plantsRepoImpl) Create(pln *model.PlantModel) error {
	qry := "INSERT INTO ms_plants(name, create_at, update_at, create_by, update_by) VALUES($1, $2, $3, $4, $5)"
	_, err := plnsRepo.db.Exec(qry, pln.Name, pln.CreateAt, pln.UpdateAt, pln.CreateBy, pln.UpdateBy)
	if err != nil {
		return fmt.Errorf("error on plantsRepoImpl.Create() : %w", err)
	}
	return nil
}

func (plnsRepo *plantsRepoImpl) Update(pln *model.PlantModel) error {
	qry := "UPDATE ms_plants SET name=$1, update_at=$2, update_by=$3 WHERE id=$4"
	_, err := plnsRepo.db.Exec(qry, pln.Name, pln.UpdateAt, pln.UpdateBy, pln.Id)
	if err != nil {
		return fmt.Errorf("error on plantsRepoImpl.Update() : %w", err)
	}
	return nil
}

func (plnsRepo *plantsRepoImpl) Delete(id int) error {
	qry := "DELETE FROM ms_plants WHERE id = $1"
	_, err := plnsRepo.db.Exec(qry, id)
	if err != nil {
		return fmt.Errorf("error on farmersRepoImpl.Delete() : %w", err)
	}
	return nil
}
