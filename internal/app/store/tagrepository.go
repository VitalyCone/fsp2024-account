package store

import "github.com/VitalyCone/account/internal/app/model"

type TagRepository struct {
	store *Store
}

func (r *TagRepository) Create(m *model.Tag) error{
	if err := r.store.db.QueryRow(
		"INSERT INTO tags (name) "+
			"VALUES ($1) RETURNING id",
		m.Name).Scan(&m.ID); err != nil {
		return err
	}

	return nil
}

func (r *TagRepository) FindById(id int) (model.Tag, error){
	var m model.Tag

	m.ID = id
	if err := r.store.db.QueryRow(
		"SELECT (name) FROM tags WHERE id = $1",
		id).Scan(
		&m.Name); err != nil {
		return model.Tag{}, err
	}
	return m, nil
}

func (r *TagRepository) DeleteById(id int) error{
	if _, err := r.store.db.Exec(
		"DELETE FROM tags WHERE id = $1",
		id); err != nil {
		return err
	}

	return nil
}