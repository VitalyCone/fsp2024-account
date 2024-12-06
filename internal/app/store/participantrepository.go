package store

import (
	"database/sql"
	"fmt"

	"github.com/VitalyCone/account/internal/app/apiserver/dtos"
	"github.com/VitalyCone/account/internal/app/model"
)

type ParticipantRepository struct {
	store *Store
}

func (r *ParticipantRepository) Create(m *model.Participant, tableName string) error {
	if _, err := r.store.db.Exec(
		fmt.Sprintf("INSERT INTO %s (username, company_id) "+
			"VALUES ($1,$2)", tableName),
		m.User.Username, m.Company.ID); err != nil {
		return err
	}
	m.User, _ = r.store.usersRepository.FindUserByUsername(m.User.Username)

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

func (r *ParticipantRepository) FindByUsername(username, tableName string) ([]model.Participant, error) {
	participants := make([]model.Participant, 0)

	rows, err := r.store.db.Query(
		fmt.Sprintf("SELECT company_id FROM %s WHERE username = $2",tableName),
		tableName, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var participant model.Participant

		err := rows.Scan(&participant.Company.ID)
		if err != nil {
			return nil, err
		}

		participant.User.Username = username
		participants = append(participants, participant)
	}
	return participants, nil
}

func (r *ParticipantRepository) FindByCompanyId(id int, tableName string) ([]model.Participant, error) {
	participants := make([]model.Participant, 0)

	rows, err := r.store.db.Query(
		fmt.Sprintf("SELECT username FROM %s WHERE company_id = $1", tableName),id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var participant model.Participant

		err := rows.Scan(&participant.User.Username)
		if err != nil {
			return nil, err
		}

		participant.Company.ID = id
		participants = append(participants, participant)
	}
	return participants, nil
}

func (r *ParticipantRepository) FindByCompanyToResponse(id int, tableName string) ([]dtos.ParticipantResponse, error) {
	participants := make([]dtos.ParticipantResponse, 0)

	rows, err := r.store.db.Query(
		fmt.Sprintf("SELECT s.company_id, u.id, u.username, u.first_name, u.second_name, u.role, u.created_at, u.updated_at, u.avatar "+
			"FROM %s s "+
			"JOIN users u ON s.username = u.username "+
			"WHERE s.company_id = $1", tableName),
		id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var participant model.Participant

		err := rows.Scan(
			&participant.Company.ID,
			&participant.User.ID,
			&participant.User.Username,
			&participant.User.FirstName,
			&participant.User.SecondName,
			&participant.User.Role,
			&participant.User.CreatedAt,
			&participant.User.UpdatedAt,
			&participant.User.Avatar,
		)
		if err != nil {
			return nil, err
		}

		participants = append(participants, dtos.ParticipantToResponse(participant))
	}
	return participants, nil
}

func (r *ParticipantRepository) IsParticipant(username, tableName string, companyId int) (bool, error) {
	var exist bool

	if err := r.store.db.QueryRow(
		fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE company_id = $1 AND username = $2)", tableName),
		companyId, username).Scan(&exist); err != nil {
		return false, err
	}
	return exist, nil
}

func (r *ParticipantRepository) Delete(m model.Participant, tableName string) error {
	if _, err := r.store.db.Exec(
		fmt.Sprintf("DELETE FROM %s WHERE username = $1 AND company_id = $2",tableName),
		m.User.Username, m.Company.ID); err != nil {
		return err
	}
	return nil
}
