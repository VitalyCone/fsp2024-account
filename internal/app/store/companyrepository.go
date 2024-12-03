package store

import "github.com/VitalyCone/account/internal/app/model"

type CompanyRepository struct {
	store *Store
}

func (r *CompanyRepository) Create(*model.Company) error{
	// if err := r.store.db.QueryRow(
	// 	"INSERT INTO users (username, password_hash, first_name, second_name) "+
	// 		"VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at",
	// 	m.Username, m.PasswordHash, m.FirstName, m.SecondName).Scan(
	// 	&m.ID, &m.CreatedAt, &m.UpdatedAt); err != nil {
	// 	return err
	// }

	// return nil
	return nil
}