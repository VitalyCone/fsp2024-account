package dtos

import "github.com/VitalyCone/account/internal/app/model"

const(
	ReviewServicesTable = "review_services"
	ReviewCompaniesTable = "review_companies"
)

type CreateReviewServiceDto struct {
	ServiceId int `json:"service_id" validate:"required, numeric"`
	Rating int `json:"rating" validate:"required,min=1,max=5"`
	CreatorUsername string `json:"creator_username" validate:"required"`
	Header string `json:"header"`
	Text string `json:"text"`
}

type CreateReviewCompanyDto struct {
	CompanyId int `json:"service_id" validate:"required, numeric"`
	Rating int `json:"rating" validate:"required,min=1,max=5"`
	CreatorUsername string `json:"creator_username" validate:"required"`
	Header string `json:"header"`
	Text string `json:"text"`
}
// ID          int
// ReviewType  ReviewType
// TypeID     int
// Rating      int
// CreatorUser User
// Header      string
// Text        string
// CreatedAt   time.Time
// UpdatedAt   time.Time

func (t *CreateReviewServiceDto) ToModel() model.Review {
	return model.Review{
		ObjectId: t.ServiceId,
		TableName: ReviewServicesTable,
		Rating: t.Rating,
		CreatorUser: model.User{Username: t.CreatorUsername},
		Header: t.Header,
		Text: t.Text,
	}
}

func (t *CreateReviewCompanyDto) ToModel() model.Review {
	return model.Review{
		ObjectId: t.CompanyId,
		TableName: ReviewCompaniesTable,
		Rating: t.Rating,
		CreatorUser: model.User{Username: t.CreatorUsername},
		Header: t.Header,
		Text: t.Text,
	}
}