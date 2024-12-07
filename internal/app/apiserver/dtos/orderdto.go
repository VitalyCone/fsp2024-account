package dtos

import (
	"time"

	"github.com/VitalyCone/account/internal/app/model"
)

type CreateOrderDto struct {
	CompanyId        int       `json:"company_id"`
	ServiceId        int       `json:"service_id"`
	WillBeFinishedAt time.Time `json:"will_be_finished_at"`
}

func (d *CreateOrderDto) ToModel(username string) model.Order {
	return model.Order{
		User:             model.User{Username: username},
		Company:          model.Company{ID: d.CompanyId},
		Service:          model.Service{ID: d.ServiceId},
		WillBeFinishedAt: d.WillBeFinishedAt,
	}
}

type OrderResponse struct {
	ID               int       `json:"id"`
	Username         string    `json:"username"`
	CompanyId        int       `json:"company_id"`
	ServiceID        int       `json:"service_id"`
	Price            float32   `json:"price"`
	OrderStatus      string    `json:"order_status"`
	WillBeFinishedAt time.Time `json:"will_be_finished_at"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func OrderModelToResponse(m model.Order) OrderResponse {
	return OrderResponse{
		ID:               m.ID,
		Username:         m.User.Username,
		CompanyId:        m.Company.ID,
		ServiceID:        m.Service.ID,
		Price:            m.Price,
		OrderStatus:      m.OrderStatus,
		WillBeFinishedAt: m.WillBeFinishedAt,
		CreatedAt:        m.CreatedAt,
		UpdatedAt:        m.UpdatedAt,
	}
}