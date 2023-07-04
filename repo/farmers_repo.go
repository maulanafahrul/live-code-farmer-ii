package repo

import (
	"database/sql"
	"fmt"
	"live-code-farmer-ii/model"
)

type FarmersRepo interface {
	GetById(int) (*model.FarmersModel, error)
	GetByName(string) (*model.FarmersModel, error)
	List() (*[]model.FarmersModel, error)
	Create(*model.FarmersModel) error
	Update(*model.FarmersModel) error
	Delete(int) error
}

type farmersRepoImpl struct {
	db *sql.DB
}

func NewFarmersRepo(db *sql.DB) FarmersRepo {
	return &farmersRepoImpl{
		db: db,
	}
}

func (frmsRepo *farmersRepoImpl) GetById(id int) (*model.FarmersModel, error) {
	qry := "SELECT id, name, address, phone_number, create_at, update_at, create_by, update_by FROM ms_farmers WHERE id = $1"

	frm := &model.FarmersModel{}
	err := frmsRepo.db.QueryRow(qry, id).Scan(&frm.Id, &frm.Name, &frm.Address, &frm.PhoneNumber, &frm.CreateAt, &frm.UpdateAt, &frm.CreateBy, &frm.UpdateBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on farmersRepoImpl.GetById() : %w", err)
	}
	return frm, nil
}
func (frmsRepo *farmersRepoImpl) GetByName(name string) (*model.FarmersModel, error) {
	qry := "SELECT id, name, address, phone_number, create_at, update_at, create_by, update_by FROM ms_farmers WHERE name = $1"

	frm := &model.FarmersModel{}
	err := frmsRepo.db.QueryRow(qry, name).Scan(&frm.Id, &frm.Name, &frm.Address, &frm.PhoneNumber, &frm.CreateAt, &frm.UpdateAt, &frm.CreateBy, &frm.UpdateBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on farmersRepoImpl.GetByName() : %w", err)
	}
	return frm, nil
}

func (frmsRepo *farmersRepoImpl) List() (*[]model.FarmersModel, error) {
	qry := "SELECT id, name, address, phone_number, create_at, update_at, create_by, update_by FROM ms_farmers"
	rows, err := frmsRepo.db.Query(qry)
	if err != nil {
		return nil, fmt.Errorf("error on farmersRepoImpl.List() : %w", err)
	}
	defer rows.Close()
	var farmers []model.FarmersModel
	for rows.Next() {
		frm := &model.FarmersModel{}
		rows.Scan(&frm.Id, &frm.Name, &frm.Address, &frm.PhoneNumber, &frm.CreateAt, &frm.UpdateAt, &frm.CreateBy, &frm.UpdateBy)
		farmers = append(farmers, *frm)
	}
	return &farmers, nil
}

func (frmsRepo *farmersRepoImpl) Create(frm *model.FarmersModel) error {
	tx, err := frmsRepo.db.Begin()
	if err != nil {
		return fmt.Errorf("AddTransaction() Begin : %w", err)
	}
	qry := "INSERT INTO ms_farmers(name, address, phone_number, create_at, update_at, create_by, update_by) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id"
	err = tx.QueryRow(qry, frm.Name, frm.Address, frm.PhoneNumber, frm.CreateAt, frm.UpdateAt, frm.CreateBy, frm.UpdateBy).Scan(&frm.Id)
	if err != nil {
		return fmt.Errorf("error on farmersRepoImpl.Create() table farmer: %w", err)
	}
	qry = "INSERT INTO farmers_plants(farmer_id) VALUES ($1)"
	_, err = tx.Exec(qry, frm.Id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("AddFarmerId() : %w", err)
	}

	tx.Commit()
	return nil
}
func (frmsRepo *farmersRepoImpl) Update(frm *model.FarmersModel) error {
	qry := "UPDATE ms_farmers SET name=$1, address=$2, phone_number=$3, update_at=$4, update_by=$5 WHERE id=$6"
	_, err := frmsRepo.db.Exec(qry, frm.Name, frm.Address, frm.PhoneNumber, frm.UpdateAt, frm.UpdateBy, frm.Id)
	if err != nil {
		return fmt.Errorf("error on farmersRepoImpl.Update() : %w", err)
	}
	return nil
}

func (frmsRepo *farmersRepoImpl) Delete(id int) error {
	qry := "DELETE FROM ms_farmers WHERE id = $1"
	_, err := frmsRepo.db.Exec(qry, id)
	if err != nil {
		return fmt.Errorf("error on farmersRepoImpl.Delete() : %w", err)
	}
	return nil
}
