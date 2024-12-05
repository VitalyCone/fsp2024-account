package dtos

import (
	"time"

	"github.com/VitalyCone/account/internal/app/model"
)

type CreateServiceDto struct {
	ServiceTypeID int     `json:"service_type_id" validate:"required"`
	Text          string  `json:"text" validate:"required"`
	Price         float32 `json:"price" validate:"required"`
}
type ServiceResponse struct {
	ID int `json:"id"`
	CompanyID int `json:"company_id"`
	ServiceType model.ServiceType `json:"service_type"`
	Text string `json:"text"`
	Price float32 `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (t *CreateServiceDto) ToModel(company_id int) model.Service {
	return model.Service{
		Company: model.Company{ID : company_id},
		ServiceType: model.ServiceType{ID : t.ServiceTypeID},
		Text: t.Text,
		Price: t.Price,
	}
}

func ModelServiceToResponse(m model.Service) ServiceResponse {
	return ServiceResponse{
		ID: m.ID,
		CompanyID: m.Company.ID,
		ServiceType: m.ServiceType,
		Text: m.Text,
		Price: m.Price,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
