package store

import (
	"database/sql"
	"fmt"

	"github.com/VitalyCone/account/internal/app/model"
)

type ParticipantRepository struct {
	store *Store
}

func (r *ParticipantRepository) Create(m model.Participant, tableName string) error {
	if _, err := r.store.db.Exec(
		fmt.Sprintf("INSERT INTO %s (username, company_id) "+
			"VALUES ($1,$2)", tableName),
		m.User.Username, m.Company.ID); err != nil {
		return err
	}

	return nil
}

func (r *ParticipantRepository) CreateWithTx(tx *sql.Tx, m model.Participant, tableName string) error {
	if _, err := tx.Exec(
		fmt.Sprintf("INSERT INTO %s (username, company_id) "+
			"VALUES ($1,$2)", tableName),
		m.User.Username, m.Company.ID); err != nil {
		return err
	}

	return nil
}

func (r *ParticipantRepository) FindByUsername(username, tableName string) ([]model.Participant, error){
	participants := make([]model.Participant, 0)

	rows, err := r.store.db.Query(
		"SELECT company_id FROM $1 WHERE username = $2",
		tableName, username)
	if err != nil{
		return nil, err
	}
	defer rows.Close()

	for rows.Next(){
		var participant model.Participant

		err := rows.Scan(&participant.Company.ID)
		if err != nil{
			return nil, err
		}

		participant.User.Username = username
		participants = append(participants, participant)
	}
	return participants, nil
}

func (r *ParticipantRepository) FindByCompanyId(id int, tableName string) ([]model.Participant, error){
	participants := make([]model.Participant, 0)

	rows, err := r.store.db.Query(
		"SELECT username FROM $1 WHERE company_id = $2",
		tableName, id)
	if err != nil{
		return nil, err
	}
	defer rows.Close()

	for rows.Next(){
		var participant model.Participant

		err := rows.Scan(&participant.User.Username)
		if err != nil{
			return nil, err
		}

		participant.Company.ID = id
		participants = append(participants, participant)
	}
	return participants, nil
}

func (r *ParticipantRepository) IsParticipant(username, tableName string, companyId int) (bool , error){
	var exist bool

	if err := r.store.db.QueryRow(
		fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE company_id = $1 AND username = $2)",tableName),
		companyId, username).Scan(&exist); err != nil {
		return false, err
	}
	return exist, nil
}

func (r *ParticipantRepository) Delete(m model.Participant, tableName string) error{
	if _, err := r.store.db.Exec(
		"DELETE FROM $1 WHERE id = $2 AND company_id = $3",
		tableName, m.User.Username, m.Company.ID); err != nil {
		return err
	}
	return nil
}