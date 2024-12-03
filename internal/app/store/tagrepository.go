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

func (r *TagRepository) FindAll() ([]model.Tag, error){
	tags := make([]model.Tag, 0)
	//"SELECT (id, name) FROM tags WHERE id = $1"
	rows, err := r.store.db.Query("SELECT id, name FROM tags")
	if err != nil{
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var tag model.Tag

		err := rows.Scan(&tag.ID, &tag.Name)
		if err != nil{
			return nil, err
		}
		tags = append(tags, tag)
	}

	return tags, nil
}

func (r *TagRepository) DeleteById(id int) error{
	if _, err := r.store.db.Exec(
		"DELETE FROM tags WHERE id = $1",
		id); err != nil {
		return err
	}

	return nil
}