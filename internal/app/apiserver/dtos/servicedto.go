package dtos

import (
	"time"

	"github.com/VitalyCone/account/internal/app/model"
)

type CreateServiceDto struct {
	CompanyID     int     `json:"company_id" validate:"required"`
	ServiceTypeID int     `json:"service_type_id" validate:"required"`
	Text          string  `json:"text" validate:"required"`
	Price         float32 `json:"price" validate:"required"`
}
type ServiceDetails struct {
	ID int `json:"id"`
	CompanyID int `json:"company_id"`
	ServiceTypeID int `json:"service_type_id"`
	Text string `json:"text"`
	Price float32 `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (t *CreateServiceDto) ToModel() model.Service {
	return model.Service{
		Company: model.Company{ID : t.CompanyID},
		ServiceType: model.ServiceType{ID : t.ServiceTypeID},
		Text: t.Text,
		Price: t.Price,
	}
}

func ServiceToDto(m model.Service) ServiceDetails {
	return ServiceDetails{
		ID: m.ID,
		CompanyID: m.Company.ID,
		ServiceTypeID: m.ServiceType.ID,
		Text: m.Text,
		Price: m.Price,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
