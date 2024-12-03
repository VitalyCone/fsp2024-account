package dtos

import "github.com/VitalyCone/account/internal/app/model"

type CreateCompanyDto struct {
	Avatar []byte `json:"avatar"`
	Name string `json:"name" validate:"required,min=1,max=100"`
	Description string `json:"description"`
}

func (t *CreateCompanyDto) ToModel() model.Company {
	return model.Company{
		Avatar: t.Avatar,
		Name: t.Name,
		Description: t.Description,
	}
}