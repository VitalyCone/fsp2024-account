package store

import "github.com/VitalyCone/account/internal/app/model"

type ReviewRepository struct {
	store *Store
}

func (r *ReviewRepository) Create(m *model.Review) error{
	if err := r.store.db.QueryRow(
		"INSERT INTO $1 (object_id, rating, creator_username, header, text) "+
			"VALUES ($2,$3,$4,$5,$6) RETURNING id, created_at, updated_at",
		m.TableName, m.ObjectId, m.Rating, m.CreatorUser.Username, m.Header, m.Text).Scan(
			&m.ID); err != nil {
		return err
	}

	return nil
}


func (r *ReviewRepository) FindAllByObjectId(tableName string, objId int)([]model.Review, error){
	var reviews []model.Review

	rows, err := r.store.db.Query(
		"SELECT (id, rating, creator_username, header, text, created_at, updated_at) FROM $1 WHERE object_id = $2",
		tableName, objId)
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
		"SELECT (object_id, rating, creator_username, header, text, created_at, updated_at) FROM $1 WHERE id = $2",
		tableName, id).Scan(
			&review.ObjectId, &review.Rating, &review.CreatorUser.Username, 
			&review.Header,&review.Text, &review.CreatedAt, &review.UpdatedAt); err != nil {
		return model.Review{}, err
	}
	return review, nil
}

func (r *ReviewRepository) DeleteById(tableName string, id int) error{
	if _, err := r.store.db.Exec(
		"DELETE FROM $1 WHERE id = $2",
		tableName, id); err != nil {
		return err
	}
	return nil
}
