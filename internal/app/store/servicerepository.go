package store

import (
	"github.com/VitalyCone/account/internal/app/apiserver/dtos"
	"github.com/VitalyCone/account/internal/app/model"
)

const tagsTableService = "tags_services"

type ServiceRepository struct {
	store *Store
}

func (r *ServiceRepository) Create(m *model.Service) error {
	tagsObj := make([]model.TagForObject, 0)

	if err := r.store.db.QueryRow(
		"INSERT INTO services (company_id, service_type_id, text, price) "+
			"VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at",
		m.Company.ID, m.ServiceType.ID, m.Text, m.Price).Scan(
		&m.ID, &m.CreatedAt, &m.UpdatedAt); err != nil {
		return err
	}
	
	for _, tag := range m.Tags{
		tagObj := model.TagForObject{
			ObjectId: m.ID,
			Tag: tag,
		}

		err := r.store.Tag().CreateForObject(&tagObj, tagsTableService)
		if err != nil{
			return err
		}
		tagsObj = append(tagsObj, tagObj)
	}

	tags , err := r.store.Tag().TagObjToTag(tagsObj)
	if err != nil{
		return err
	}
	m.Tags = tags

	servicetype, _ := r.store.ServiceType().FindById(m.ServiceType.ID)
	m.ServiceType.Name = servicetype.Name

	return nil
}

func (r *ServiceRepository) FindByCompanyId(id int) ([]model.Service, error) {
	services := make([]model.Service, 0)

	rows, err := r.store.db.Query(
		"SELECT s.id, s.company_id, st.id, st.name, s.text, s.price, s.created_at, s.updated_at "+
		"FROM services s " +
		"JOIN service_types st ON s.service_type_id = st.id "+
		"WHERE s.company_id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var service model.Service

		err := rows.Scan(
			&service.ID, &service.Company.ID, &service.ServiceType.ID, &service.ServiceType.Name,  &service.Text,
			&service.Price, &service.CreatedAt, &service.UpdatedAt)
		if err != nil {
			return nil, err
		}

		tagObj, _ := r.store.Tag().FindByObjectIdForObject(service.ID, tagsTableService)
		service.Tags, _ =  r.store.Tag().TagObjToTag(tagObj)

		services = append(services, service)
	}

	return services, nil
}

func (r *ServiceRepository) FindByCompanyIdToResponse(id int) ([]dtos.ServiceResponse, error) {
	services := make([]dtos.ServiceResponse, 0)

	rows, err := r.store.db.Query(
		"SELECT s.id, s.company_id, st.id, st.name, s.text, s.price, s.created_at, s.updated_at "+
		"FROM services s " +
		"JOIN service_types st ON s.service_type_id = st.id "+
		"WHERE s.company_id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var service model.Service

		err := rows.Scan(
			&service.ID, &service.Company.ID, &service.ServiceType.ID, &service.ServiceType.Name, &service.Text,
			&service.Price, &service.CreatedAt, &service.UpdatedAt)
		if err != nil {
			return nil, err
		}

		tagObj, _ := r.store.Tag().FindByObjectIdForObject(service.ID, tagsTableService)
		service.Tags, _ =  r.store.Tag().TagObjToTag(tagObj)

		services = append(services, dtos.ModelServiceToResponse(service))
	}

	return services, nil
}


func (r *ServiceRepository) FindById(id int) (model.Service, error) {
	var service model.Service


	if err := r.store.db.QueryRow(
		"SELECT s.id, s.company_id, st.id, st.name, s.text, s.price, s.created_at, s.updated_at "+
		"FROM services s " +
		"JOIN service_types st ON s.service_type_id = st.id "+
		"WHERE s.id = $1",id).Scan(
		&service.ID, &service.Company.ID, &service.ServiceType.ID, &service.ServiceType.Name,
		&service.Text, &service.Price, &service.CreatedAt, &service.UpdatedAt); err != nil {
		return model.Service{}, err
	}

	tagObj, _ := r.store.Tag().FindByObjectIdForObject(service.ID, tagsTableService)
	service.Tags, _ =  r.store.Tag().TagObjToTag(tagObj)

	return service, nil
}

func (r *ServiceRepository) DeleteById(id int) error {

	if _, err := r.store.db.Exec(
		"DELETE FROM services WHERE id = $1",
		id); err != nil {
		return err
	}

	return nil
}
