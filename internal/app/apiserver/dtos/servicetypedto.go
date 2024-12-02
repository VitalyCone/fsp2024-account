package dtos

import "github.com/VitalyCone/account/internal/app/model"

type CreateServiceTypeDto struct {
	Name string `json:"name" validate:"required,min=1,max=50"`
}

func (t *CreateServiceTypeDto) ToModel() model.ServiceType {
	return model.ServiceType{
		Name: t.Name,
	}
}