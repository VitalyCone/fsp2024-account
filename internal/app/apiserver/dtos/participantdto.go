package dtos

import "github.com/VitalyCone/account/internal/app/model"

type CreateParticipantDto struct {
	Username string `json:"username" validate:"required"`
}

func (cpd *CreateParticipantDto) ToModel(companyId int) model.Participant{
	return model.Participant{
		User: model.User{Username: cpd.Username},
		Company: model.Company{ID: companyId},
	}
}

type ParticipantResponse struct {
	User UserResponse `json:"user"`
	CompanyId int `json:"company_id"`
}

func ParticipantToResponse (m model.Participant) ParticipantResponse{
	return ParticipantResponse{
		User: UserToResponse(m.User),
		CompanyId: m.Company.ID,
	}
}
