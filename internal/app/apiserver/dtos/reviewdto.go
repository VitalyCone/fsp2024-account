package dtos

import (
	"time"

	"github.com/VitalyCone/account/internal/app/model"
)

const (
	ReviewServicesTable  = "review_services"
	ReviewCompaniesTable = "review_companies"
)

type CreateReviewServiceDto struct {
	Rating          int    `json:"rating" validate:"required,min=1,max=5"`
	Header          string `json:"header"`
	Text            string `json:"text"`
}

type CreateReviewCompanyDto struct {
	Rating          int    `json:"rating" validate:"required,min=1,max=5"`
	Header          string `json:"header"`
	Text            string `json:"text"`
}

type ReviewResponce struct{
	ID        int
	TableName string
	ObjectId  int
	Rating      int
	CreatorUser UserResponse
	Header      string
	Text        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func ReviewToResponce(m model.Review) ReviewResponce{

	return ReviewResponce{
		ID:        m.ID,
		TableName: m.TableName,
		ObjectId:  m.ObjectId,
		Rating:      m.Rating,
		CreatorUser: UserToResponse(m.CreatorUser),
		Header:      m.Header,
		Text:        m.Text,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}
	// ID        int
	// TableName string
	// ObjectId  int
	// // ReviewType  ReviewType
	// // TypeID     int
	// Rating      int
	// CreatorUser User
	// Header      string
	// Text        string
	// CreatedAt   time.Time
	// UpdatedAt   time.Time

func (t *CreateReviewServiceDto) ToModel(serviceId int, username string) model.Review {
	return model.Review{
		ObjectId:    serviceId,
		TableName:   ReviewServicesTable,
		Rating:      t.Rating,
		CreatorUser: model.User{Username: username},
		Header:      t.Header,
		Text:        t.Text,
	}
}

func (t *CreateReviewCompanyDto) ToModel(companyId int, username string) model.Review {
	return model.Review{
		ObjectId:    companyId,
		TableName:   ReviewCompaniesTable,
		Rating:      t.Rating,
		CreatorUser: model.User{Username: username},
		Header:      t.Header,
		Text:        t.Text,
	}
}


