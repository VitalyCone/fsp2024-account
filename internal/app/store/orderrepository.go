package store

import (
	"github.com/VitalyCone/account/internal/app/apiserver/dtos"
	"github.com/VitalyCone/account/internal/app/model"
)

type OrderRepository struct {
	store *Store
}

func (r *OrderRepository) Create(order *model.Order, user model.User) error{
	tx, err := r.store.db.Begin()
	if err != nil{
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(
		"INSERT INTO orders (username, company_id, service_id, order_status, price, will_be_finished_at) "+
			"VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, created_at, updated_at")
	if err != nil{
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		order.User.Username,
		order.Company.ID,
		order.Service.ID,
		order.OrderStatus,
		order.Price,
		order.WillBeFinishedAt).Scan(&order.ID, &order.CreatedAt, &order.UpdatedAt)
	if err != nil{
		return err
	}

	err = r.store.User().ModifyBalanceWithTx(user.Balance, user.Username, tx)
	if err != nil{
		return err
	}
	
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *OrderRepository) FindAllByUsernameToResponse(username string) ([]dtos.OrderResponse, error){
	orders := make([]dtos.OrderResponse, 0)

	rows, err := r.store.db.Query(
		"SELECT id, username, company_id, service_id, order_status, price, will_be_finished_at, created_at, updated_at FROM orders WHERE username = $1 ORDER BY created_at DESC", username)
	if err != nil{
		return nil, err
	}
	defer rows.Close()

	for rows.Next(){
		var order model.Order

		err := rows.Scan(
			&order.ID, 
			&order.User.Username, 
			&order.Company.ID, 
			&order.Service.ID, 
			&order.OrderStatus, 
			&order.Price,
			&order.WillBeFinishedAt,
			&order.CreatedAt,
			&order.UpdatedAt, 
		)
		if err != nil{
			return nil, err
		}

		orders = append(orders, dtos.OrderModelToResponse(order))
	}
	return orders, nil
}

func (r *OrderRepository) FindAllByCompanyIdToResponse(company_id int) ([]dtos.OrderResponse, error){
	orders := make([]dtos.OrderResponse, 0)

	rows, err := r.store.db.Query(
		"SELECT id, username, company_id, service_id, order_status, price, will_be_finished_at, created_at, updated_at FROM orders WHERE company_id = $1 ORDER BY created_at DESC", company_id)
	if err != nil{
		return nil, err
	}
	defer rows.Close()

	for rows.Next(){
		var order model.Order

		err := rows.Scan(
			&order.ID, 
			&order.User.Username, 
			&order.Company.ID, 
			&order.Service.ID, 
			&order.OrderStatus, 
			&order.Price,
			&order.WillBeFinishedAt,
			&order.CreatedAt,
			&order.UpdatedAt, 
		)
		if err != nil{
			return nil, err
		}

		orders = append(orders, dtos.OrderModelToResponse(order))
	}
	return orders, nil
}

func (r *OrderRepository) FindById(id int) (model.Order, error){
	var order model.Order

	err := r.store.db.QueryRow(
		"SELECT id, username, company_id, service_id, order_status, price, will_be_finished_at, created_at, updated_at "+
		"FROM orders WHERE id = $1", id).Scan(
			&order.ID,
			&order.User.Username,
			&order.Company.ID,
			&order.Service.ID,
			&order.OrderStatus,
			&order.Price,
			&order.WillBeFinishedAt,
			&order.CreatedAt,
			&order.UpdatedAt)
	if err != nil{
		return model.Order{}, err
	}
	return order, nil
}

func (r *OrderRepository) DeleteById(id int) error{
	_, err := r.store.db.Exec(
		"DELETE FROM orders WHERE id = $1", id)
	if err != nil{
		return err
	}
	return nil
}