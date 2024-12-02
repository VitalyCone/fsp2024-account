package store

import "github.com/VitalyCone/account/internal/app/model"

type ServiceTypeRepository struct {
	store *Store
}

func (r *ServiceTypeRepository) Create(m *model.ServiceType) error{
	if err := r.store.db.QueryRow(
		"INSERT INTO service_types (name) "+
			"VALUES ($1) RETURNING id",
		m.Name).Scan(&m.ID); err != nil {
		return err
	}

	return nil
}

func (r *ServiceTypeRepository) FindById(id int) (model.ServiceType,error){
	var m model.ServiceType

	m.ID = id
	if err := r.store.db.QueryRow(
		"SELECT (name) FROM service_types WHERE id = $1",
		id).Scan(
		&m.Name); err != nil {
		return model.ServiceType{}, err
	}
	return m, nil
}

func (r *ServiceTypeRepository) FindAll() ([]model.ServiceType,error){
	var array []model.ServiceType

	rows, err := r.store.db.Query(
		"SELECT id, name FROM service_types")
	if err != nil{
		return nil, err
	}
	defer rows.Close()

	for rows.Next(){
		var serviceType model.ServiceType

		err := rows.Scan(&serviceType.ID, &serviceType.Name)
		if err != nil{
			return nil, err
		}

		array = append(array, serviceType)
	}

	return array, nil
}

func (r *ServiceTypeRepository) DeleteById(id int) error {
	if _, err := r.store.db.Exec(
		"DELETE FROM service_types WHERE id = $1",
		id); err != nil {
		return err
	}

	return nil
}