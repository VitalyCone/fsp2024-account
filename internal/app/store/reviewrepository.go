package store

import (
	"fmt"

	"github.com/VitalyCone/account/internal/app/model"
)

type ReviewRepository struct {
	store *Store
}

func (r *ReviewRepository) Create(m *model.Review) error{
	if err := r.store.db.QueryRow(
		fmt.Sprintf("INSERT INTO %s (object_id, rating, creator_username, header, text) "+
			"VALUES ($1,$2,$3,$4,$5) RETURNING id, created_at, updated_at", m.TableName),
		m.ObjectId, m.Rating, m.CreatorUser.Username, m.Header, m.Text).Scan(
			&m.ID, &m.CreatedAt, &m.UpdatedAt); err != nil {
		return err
	}

	return nil
}


func (r *ReviewRepository) FindAllByObjectId(tableName string, objId int)([]model.Review, error){
	reviews := make([]model.Review, 0)

	rows, err := r.store.db.Query(
		fmt.Sprintf("SELECT (id, rating, creator_username, header, text, created_at, updated_at) FROM %s WHERE object_id = $1", tableName),
		objId)
	if err != nil{
		return nil, err
	}
	defer rows.Close()

	for rows.Next(){
		var review model.Review

		err := rows.Scan(
			&review.ID, &review.Rating, &review.CreatorUser.Username, &review.Header,
			&review.Text, &review.CreatedAt, &review.UpdatedAt)
		if err != nil{
			return nil, err
		}

		review.ObjectId = objId
		reviews = append(reviews, review)
	}
	return reviews, nil
}


func (r *ReviewRepository) FindById(tableName string, id int)(model.Review, error){
	var review model.Review

	review.ID = id
	if err := r.store.db.QueryRow(
		fmt.Sprintf("SELECT (object_id, rating, creator_username, header, text, created_at, updated_at) FROM %s WHERE id = $1", tableName),
		id).Scan(
			&review.ObjectId, &review.Rating, &review.CreatorUser.Username, 
			&review.Header,&review.Text, &review.CreatedAt, &review.UpdatedAt); err != nil {
		return model.Review{}, err
	}
	return review, nil
}

func (r *ReviewRepository) DeleteById(tableName string, id int) error{
	if _, err := r.store.db.Exec(
		fmt.Sprintf("DELETE FROM %s WHERE id = $1", tableName),
		id); err != nil {
		return err
	}
	return nil
}
