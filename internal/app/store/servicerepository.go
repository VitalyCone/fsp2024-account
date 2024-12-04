package store

import "github.com/VitalyCone/account/internal/app/model"

type ServiceRepository struct {
	store *Store
}

func (r *ServiceRepository) Create(m *model.Service) error {
	if err := r.store.db.QueryRow(
		"INSERT INTO service (company_id, service_type_id, text, price) "+
			"VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at",
		m.Company.ID, m.ServiceType.ID, m.Text, m.Price).Scan(
		&m.ID, &m.CreatedAt, &m.UpdatedAt); err != nil {
		return err
	}

	return nil
}

func (r *ServiceRepository) FindByCompanyId(id int) ([]model.Service, error) {
	services := make([]model.Service, 0)

	rows, err := r.store.db.Query(
		"SELECT id, company_id, service_type_id, text, price, created_at, updated_at FROM service")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var service model.Service

		err := rows.Scan(
			&service.ID, &service.Company.ID, &service.ServiceType.ID, &service.Text,
			&service.Price, &service.CreatedAt, &service.UpdatedAt, &service.UpdatedAt)
		if err != nil {
			return nil, err
		}
		services = append(services, service)
	}

	return services, nil
}

func (r *ServiceRepository) FindById(id int) (model.Service, error) {
	var service model.Service

	if err := r.store.db.QueryRow(
		"SELECT id, company_id, service_type_id, text, price, created_at, updated_at "+
			"FROM service WHERE id == $1",
		id).Scan(
		&service.ID, &service.Company.ID, &service.ServiceType.ID,
		&service.Text, &service.Price, &service.CreatedAt, &service.UpdatedAt); err != nil {
		return model.Service{}, err
	}

	return service, nil
}

func (r *ServiceRepository) DeleteById(id int) error {

	if _, err := r.store.db.Exec(
		"DELETE FROM service WHERE id == $1",
		id); err != nil {
		return err
	}

	return nil
}
