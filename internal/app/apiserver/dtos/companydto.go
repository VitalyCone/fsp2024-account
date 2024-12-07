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
	Avatar          []byte `json:"avatar"`
	Name            string `json:"name" validate:"required,min=1,max=100"`
	Description     string `json:"description"`
	Email           string `json:"email" validate:"required,email"`
	Phone           string `json:"phone" validate:"required"`
	INN             string `json:"inn" validate:"required"`
	ManagerTelegram string `json:"manager_telegram"`
}

type CreateCompanyResponse struct {
	ID              int       `json:"id"`
	Avatar          []byte    `json:"avatar"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	Email           string    `json:"email"`
	Phone           string    `json:"phone"`
	INN             string    `json:"inn"`
	ManagerTelegram string    `json:"manager_telegram"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (t *CreateCompanyDto) ToModel() model.Company {
	return model.Company{
		Avatar:          t.Avatar,
		Name:            t.Name,
		Description:     t.Description,
		Email:           t.Email,
		Phone:           t.Phone,
		INN:             t.INN,
		ManagerTelegram: t.ManagerTelegram,
	}
}

func ModelToCreateCompanyResponse(company model.Company) CreateCompanyResponse {
	return CreateCompanyResponse{
		ID:              company.ID,
		Avatar:          company.Avatar,
		Name:            company.Name,
		Description:     company.Description,
		Email:           company.Email,
		Phone:           company.Phone,
		INN:             company.INN,
		ManagerTelegram: company.ManagerTelegram,
		CreatedAt:       company.CreatedAt,
		UpdatedAt:       company.UpdatedAt,
	}
}
