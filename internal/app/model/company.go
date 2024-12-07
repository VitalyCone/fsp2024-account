package model

import "time"

type Company struct {
	ID          int `json:"id"`
	Avatar      []byte `json:"avatar"`
	Name        string `json:"name"`
	Services    []Service `json:"services"`// -db
	Description string `json:"description"`
	Tags        []Tag
	Members    []User `json:"members"`// -db 
	Moderators []User `json:"moderators"`// -db
	Reviews    []Review `json:"reviews"`
	INN string `json:"inn"` //
	Email string `json:"email"` //
	ManagerTelegram string `json:"manager_telegram"` //
	Phone string `json:"phone"` // 
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
