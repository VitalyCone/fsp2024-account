package dtos

import "github.com/VitalyCone/account/internal/app/model"

type CreateTagDto struct {
	Name string `json:"name" validate:"required,min=1,max=50,alphanum"`
}

func (t *CreateTagDto) ToModel() model.Tag {
	return model.Tag{
		Name: t.Name,
	}
}