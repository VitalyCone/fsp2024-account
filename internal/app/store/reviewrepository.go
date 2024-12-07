package store

import (
	"fmt"

	"github.com/VitalyCone/account/internal/app/apiserver/dtos"
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

	user, _ := r.store.User().FindUserByUsername(m.CreatorUser.Username)
	m.CreatorUser = user

	return nil
}


func (r *ReviewRepository) FindAllByObjectId(tableName string, objId int)([]model.Review, error){
	reviews := make([]model.Review, 0)
	//JOIN users u ON r.creator_username = u.username
	rows, err := r.store.db.Query(
		fmt.Sprintf("SELECT id, rating, header, text, created_at, updated_at, "+
			"u.id, u.username, u.first_name, u.secondname, u.role, u.created_at, u.updated_at, u.avatar, u.balance FROM %s "+
			"JOIN users u ON %s.creator_username = u.username " +
			"WHERE object_id = $1", tableName, tableName),
		objId)
	if err != nil{
		return nil, err
	}
	defer rows.Close()

	for rows.Next(){
		var review model.Review

		err := rows.Scan(
			&review.ID, &review.Rating, &review.Header,
			&review.Text, &review.CreatedAt, &review.UpdatedAt,
			&review.CreatorUser.ID, &review.CreatorUser.Username, &review.CreatorUser.FirstName,
			&review.CreatorUser.SecondName, &review.CreatorUser.Role, &review.CreatorUser.CreatedAt,
			&review.CreatorUser.UpdatedAt, &review.CreatorUser.Avatar, &review.CreatorUser.Balance)
		if err != nil{
			return nil, err
		}

		review.ObjectId = objId
		reviews = append(reviews, review)
	}
	return reviews, nil
}

func (r *ReviewRepository) FindAllByObjectIdToResponse(tableName string, objId int)([]dtos.ReviewResponce, error){
	reviews := make([]dtos.ReviewResponce, 0)
	//JOIN users u ON r.creator_username = u.username
	rows, err := r.store.db.Query(
		fmt.Sprintf("SELECT s.id, s.rating, s.header, s.text, s.created_at, s.updated_at, "+
			"u.id, u.username, u.first_name, u.second_name, u.role, u.created_at, u.updated_at, u.avatar, u.balance FROM %s s "+
			"JOIN users u ON s.creator_username = u.username " +
			"WHERE object_id = $1", tableName),
		objId)
	if err != nil{
		return nil, err
	}
	defer rows.Close()

	for rows.Next(){
		var review model.Review

		err := rows.Scan(
			&review.ID, &review.Rating, &review.Header,
			&review.Text, &review.CreatedAt, &review.UpdatedAt,
			&review.CreatorUser.ID, &review.CreatorUser.Username, &review.CreatorUser.FirstName,
			&review.CreatorUser.SecondName, &review.CreatorUser.Role, &review.CreatorUser.CreatedAt,
			&review.CreatorUser.UpdatedAt, &review.CreatorUser.Avatar, &review.CreatorUser.Balance)
		if err != nil{
			return nil, err
		}

		review.ObjectId = objId
		reviews = append(reviews, dtos.ReviewToResponce(review))
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

	user, _ := r.store.User().FindUserByUsername(review.CreatorUser.Username)
	review.CreatorUser = user
	
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
