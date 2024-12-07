package store

import (
	"errors"

	"github.com/VitalyCone/account/internal/app/apiserver/dtos"
	"github.com/VitalyCone/account/internal/app/model"
)

type CompanyRepository struct {
	store *Store
}

func (r *CompanyRepository) Create(m *model.Company, creatorUsername, membersParticipantTable, moderatorsParticipantTable string) error{
	
	participants := make([]model.Participant, 0)

	tx, err := r.store.db.Begin()
	if err != nil{
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(
		"INSERT INTO companies (avatar, name, description, email, phone, inn, manager_telegram) "+
			"VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, created_at, updated_at")
	if err != nil{
		return errors.New("company: " + err.Error())
	}

	stmt.QueryRow(
		m.Avatar,
		m.Name, 
		m.Description,
		m.Email,
		m.Phone,
		m.INN,
		m.ManagerTelegram).Scan(&m.ID, &m.CreatedAt, &m.UpdatedAt)
	
	participants = append(participants, model.Participant{
		Company: *m,
		User: model.User{Username: creatorUsername}})

	//return fmt.Errorf("%+v .... %s",participants[0], membersParticipantTable)
	err1 := r.store.Participant().CreateWithTx(tx, participants[0], membersParticipantTable)
	err2 := r.store.Participant().CreateWithTx(tx, participants[0], moderatorsParticipantTable)
	if err1 != nil || err2 != nil {
		return errors.New("participant: " + errors.Join(err1, err2).Error())
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	stmt.Close()

	// users, err := r.store.User().FindUsersByParticipants(participants)
	// if err != nil{
	// 	return errors.New("users: " + err.Error())
	// }

	// m.Members, m.Moderators = users, users

	// m.Reviews = make([]model.Review, 0)

	return nil
}

func (r *CompanyRepository) FindAll() ([]model.Company, error) {
	companies :=  make([]model.Company, 0)

	rows, err := r.store.db.Query(
		"SELECT id,avatar,name,description,email, phone, inn, manager_telegram,created_at,updated_at FROM companies")
	if err != nil{
		return nil, err
	}
	defer rows.Close()

	for rows.Next(){
		var company model.Company

		err := rows.Scan(
			&company.ID, 
			&company.Avatar, 
			&company.Name, 
			&company.Description, 
			&company.Email, 
			&company.Phone,
			&company.INN,
			&company.ManagerTelegram,
			&company.CreatedAt, 
			&company.UpdatedAt)
		if err != nil{
			return nil, err
		}

		companies = append(companies, company)
	}

	return companies, nil
}

func (r *CompanyRepository) FindAllToCreateCompanyResponse() ([]dtos.CreateCompanyResponse, error) {
	companies :=  make([]dtos.CreateCompanyResponse, 0)

	rows, err := r.store.db.Query(
		"SELECT id,avatar,name,description, email, phone, inn, manager_telegram, created_at,updated_at FROM companies")
	if err != nil{
		return nil, err
	}
	defer rows.Close()

	for rows.Next(){
		var company model.Company

		err := rows.Scan(
			&company.ID, 
			&company.Avatar, 
			&company.Name, 
			&company.Description, 
			&company.Email, 
			&company.Phone,
			&company.INN,
			&company.ManagerTelegram,
			&company.CreatedAt, 
			&company.UpdatedAt)
		if err != nil{
			return nil, err
		}

		companies = append(companies, dtos.ModelToCreateCompanyResponse(company))
	}

	return companies, nil
}


func (r *CompanyRepository) FindById(id int, membersParticipantTable, moderatorsParticipantTable, reviewsTable string) (model.Company, error) {
	var company model.Company

	company.ID = id
	if err := r.store.db.QueryRow(
		"SELECT avatar,name,description, email, phone, inn, manager_telegram,created_at,updated_at FROM companies WHERE id = $1",
		id).Scan(
		&company.Avatar, 
		&company.Name, 
		&company.Description, 
		&company.Email, 
		&company.Phone,
		&company.INN,
		&company.ManagerTelegram,
		&company.CreatedAt, 
		&company.UpdatedAt); err != nil {
		return model.Company{}, err
	}

	return company, nil
}

func (r *CompanyRepository) FindByName(name string) (model.Company, error) {
	var company model.Company

	company.Name = name
	if err := r.store.db.QueryRow(
		"SELECT avatar,id,description, email, phone, inn, manager_telegram,created_at,updated_at FROM companies WHERE id = $1",
		name).Scan(
		&company.Avatar, 
		&company.ID,
		&company.Description,
		&company.Email, 
		&company.Phone,
		&company.INN,
		&company.ManagerTelegram,
		&company.CreatedAt, 
		&company.UpdatedAt); err != nil {
		return model.Company{}, err
	}
	return company, nil
}

func (r *CompanyRepository) DeleteById(id int) error{
	if _, err := r.store.db.Exec(
		"DELETE FROM companies WHERE id = $1",
		id); err != nil {
		return err
	}
	return nil
}

func (r *CompanyRepository) DeleteByName(name string) error{
	if _, err := r.store.db.Exec(
		"DELETE FROM companies WHERE name = $1",
		name); err != nil {
		return err
	}
	return nil
}