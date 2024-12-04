package dtos

import "github.com/VitalyCone/account/internal/app/model"

const(
	MembersParticipantTable = "user_company_members"
	ModeratorsParticipantTable = "user_company_moderators"
)

type CreateCompanyDto struct {
	Avatar string `json:"avatar"`
	Name string `json:"name" validate:"required,min=1,max=100"`
	Description string `json:"description"`
}

func (t *CreateCompanyDto) ToModel() model.Company {
	return model.Company{
		Avatar: []byte(t.Avatar),
		Name: t.Name,
		Description: t.Description,
	}
}