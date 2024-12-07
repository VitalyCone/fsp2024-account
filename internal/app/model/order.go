package model

import "time"

type Order struct {
	ID               int       `json:"id"`
	User             User      `json:"user"`
	Company          Company   `json:"company"`
	Service          Service   `json:"service"`
	Price            float32   `json:"price"`
	OrderStatus      string    `json:"order_status"`
	WillBeFinishedAt time.Time `json:"will_be_finished_at"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
