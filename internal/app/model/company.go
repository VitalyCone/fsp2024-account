package model

import "time"

type Company struct {
	ID          int
	Avatar      []byte
	Name        string
	Services    []Service // -db
	Description string
	Tags        []Tag
	Member      []User // -db
	Moderators  []User // -db
	Reviews     []Review
	CreatedAt   time.Time
	UpdatedAt	time.Time
}