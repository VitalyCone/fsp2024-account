package dtos

import (
	"time"

	"github.com/VitalyCone/account/internal/app/model"
)

type CreateServiceDto struct {
	ServiceTypeID int     `json:"service_type_id" validate:"required"`
	Text          string  `json:"text" validate:"required"`
	Price         float32 `json:"price" validate:"required"`
	TagsIds       []int   `json:"tagsIds"`
}
type ServiceResponse struct {
	ID          int               `json:"id"`
	CompanyID   int               `json:"company_id"`
	ServiceType model.ServiceType `json:"service_type"`
	Text        string            `json:"text"`
	Price       float32           `json:"price"`
	Rating      float32           `json:"rating"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
	Tags        []model.Tag       `json:"tags"`
}

func (t *CreateServiceDto) ToModel(company_id int) model.Service {
	tags := make([]model.Tag, len(t.TagsIds))
	for i, id := range t.TagsIds {
		tags[i] = model.Tag{ID: id}
	}
	return model.Service{
		Company:     model.Company{ID: company_id},
		ServiceType: model.ServiceType{ID: t.ServiceTypeID},
		Text:        t.Text,
		Price:       t.Price,
		Tags:        tags,
	}
}

func ModelServiceToResponse(m model.Service) ServiceResponse {
	return ServiceResponse{
		ID:          m.ID,
		CompanyID:   m.Company.ID,
		ServiceType: m.ServiceType,
		Text:        m.Text,
		Rating:      m.Rating,
		Price:       m.Price,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
		Tags:        m.Tags,
	}
}
