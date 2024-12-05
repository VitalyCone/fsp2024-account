package dtos

import (
	"time"

	"github.com/VitalyCone/account/internal/app/model"
)

const (
	MembersParticipantTable    = "user_company_members"
	ModeratorsParticipantTable = "user_company_moderators"
)

type CreateCompanyDto struct {
	Avatar      []byte `json:"avatar"`
	Name        string `json:"name" validate:"required,min=1,max=100"`
	Description string `json:"description"`
}

type CreateCompanyResponse struct {
	ID          int       `json:"id"`
	Avatar      []byte    `json:"avatar"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (t *CreateCompanyDto) ToModel() model.Company {
	return model.Company{
		Avatar:      t.Avatar,
		Name:        t.Name,
		Description: t.Description,
	}
}

func ModelToCreateCompanyResponse(company model.Company) CreateCompanyResponse {
	return CreateCompanyResponse{
		ID:          company.ID,
		Avatar:      company.Avatar,
		Name:        company.Name,
		Description: company.Description,
		CreatedAt:   company.CreatedAt,
		UpdatedAt:   company.UpdatedAt,
	}
}
