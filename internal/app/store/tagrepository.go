package store

import (
	"database/sql"
	"fmt"

	"github.com/VitalyCone/account/internal/app/model"
)

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

func (r *TagRepository) FindByIdWithTx(id int, tx *sql.Tx) (model.Tag, error){
	var m model.Tag

	m.ID = id
	if err := tx.QueryRow(
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

/*


Object Tags


*/

func (r *TagRepository) CreateForObject(m *model.TagForObject, tableName string) error{
	if _, err := r.store.db.Exec(
		fmt.Sprintf("INSERT INTO %s (tag_id, object_id) "+
			"VALUES ($1, $2)", tableName),
		m.Tag.ID, m.ObjectId); err != nil {
		return err
	}

	return nil
}

func (r *TagRepository) CreateForObjectWithTx(m *model.TagForObject,tx *sql.Tx, tableName string) error{
	if _, err := tx.Exec(
		fmt.Sprintf("INSERT INTO %s (tag_id, object_id) "+
			"VALUES ($1, $2)", tableName),
		m.Tag.ID, m.ObjectId); err != nil {
		return err
	}

	return nil
}

func (r *TagRepository) FindByTagIdForObject(id int, tableName string) ([]model.TagForObject, error){
	//		fmt.Sprintf("INSERT INTO %s (tag_id, service_id) "+
	//"VALUES ($1, $2)", tableName),
	tags := make([]model.TagForObject, 0)
	//"SELECT (id, name) FROM tags WHERE id = $1"
	rows, err := r.store.db.Query(fmt.Sprintf("SELECT tag_id, object_id FROM %s WHERE tag_id = $1", tableName), id)
	if err != nil{
		return nil, err
	}
	defer rows.Close()
	
	for rows.Next() {
		var tag model.TagForObject
		
		err := rows.Scan(&tag.Tag.ID, &tag.ObjectId)
		if err != nil{
			return nil, err
		}
		tags = append(tags, tag)
	}
	
	return tags, nil
}

func (r *TagRepository) FindByObjectIdForObject(id int, tableName string) ([]model.TagForObject, error){
	//		fmt.Sprintf("INSERT INTO %s (tag_id, service_id) "+
	//"VALUES ($1, $2)", tableName),
	tags := make([]model.TagForObject, 0)
	//"SELECT (id, name) FROM tags WHERE id = $1"
	// rows, err := r.store.db.Query(
	// 	fmt.Sprintf("SELECT id, rating, header, text, created_at, updated_at, "+
	// 		"u.id, u.username, u.first_name, u.secondname, u.role, u.created_at, u.updated_at, u.avatar FROM %s "+
	// 		"JOIN users u ON %s.creator_username = u.username " +
	// 		"WHERE object_id = $1", tableName, tableName),
	// 	objId)
	// if err != nil{
	// 	return nil, err
	// }
	rows, err := r.store.db.Query(
		fmt.Sprintf("SELECT tag_id, object_id FROM %s WHERE object_id = $1", tableName), id)
	if err != nil{
		return nil, err
	}
	defer rows.Close()
	
	for rows.Next() {
		var tag model.TagForObject
		
		err := rows.Scan(&tag.Tag.ID, &tag.ObjectId)
		if err != nil{
			return nil, err
		}
		tags = append(tags, tag)
	}
	
	return tags, nil
}

func (r *TagRepository) DeleteByObjectIdForObject(id int, tableName string) error{
	_, err := r.store.db.Exec(
		fmt.Sprintf("DELETE FROM %s WHERE object_id = $1", tableName), id)
	if err != nil{
		return err
	}
	return nil
}

func (r *TagRepository) DeleteByTagIdForObject(id int, tableName string) error{
	_, err := r.store.db.Exec(
		fmt.Sprintf("DELETE FROM %s WHERE tag_id = $1", tableName), id)
	if err != nil{
		return err
	}
	return nil
}

func (r *TagRepository) TagObjToTag(tagsForObj []model.TagForObject) ([]model.Tag, error) {
    tags := make([]model.Tag, 0)

    for _, tagObj := range tagsForObj {
        tag, err := r.FindById(tagObj.Tag.ID)
        if err != nil {
            return nil, err
        }

        tags = append(tags, tag)
    }

    return tags, nil
}

func (r *TagRepository) TagObjToTagWithTx(tx *sql.Tx, tagsForObj []model.TagForObject) ([]model.Tag, error) {
    tags := make([]model.Tag, 0)

    for _, tagObj := range tagsForObj {
        tag, err := r.FindByIdWithTx(tagObj.Tag.ID, tx)
        if err != nil {
            return nil, err
        }

        tags = append(tags, tag)
    }

    return tags, nil
}