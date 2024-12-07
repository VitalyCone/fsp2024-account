package dtos

import (
	"strings"
	"time"

	"github.com/VitalyCone/account/internal/app/model"
)

type CreateUserDto struct {
	Username   string `json:"username" form:"username" validate:"required,alphanum,min=3,max=32"`
	Password   string `json:"password" form:"password" validate:"required,min=3,max=32"`
	FirstName  string `json:"first_name" form:"first_name" validate:"required,max=50"`
	SecondName string `json:"second_name" form:"second_name" validate:"required,max=50"`
	Role       string `json:"role" form:"role" validate:"required,oneof=user admin"` //"user"/"admin"
}

func (c *CreateUserDto) ToModel(passHash string) model.User {
	avatar := make([]byte, 0)
	return model.User{
		Username:     strings.ToLower(c.Username),
		PasswordHash: passHash,
		FirstName:    c.FirstName,
		SecondName:   c.SecondName,
		Role:         c.Role,
		Avatar:       avatar,
	}
}

type ModifyUserDto struct {
	Username    string `json:"username" form:"username"`
	OldPassword string `json:"old_password" form:"old_password"`
	NewPassword string `json:"new_password" form:"new_password"`
	Avatar      []byte `json:"avatar" form:"avatar"`
	FirstName   string `json:"first_name" form:"first_name"`
	SecondName  string `json:"second_name" form:"second_name"`
	Balance     float32    `json:"balance"`
}

func (u *ModifyUserDto) ToModel(passHash string) model.User {
	return model.User{}
}

type UserDto struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type UserResponse struct {
	ID         int       `json:"id"`
	Avatar     []byte    `json:"avatar"`
	Username   string    `json:"username"`
	FirstName  string    `json:"first_name"`
	SecondName string    `json:"second_name"`
	Balance    float32       `json:"balance"`
	Role       string    `json:"role"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func UserToResponse(m model.User) UserResponse {
	return UserResponse{
		ID:         m.ID,
		Avatar:     m.Avatar,
		Username:   m.Username,
		FirstName:  m.FirstName,
		SecondName: m.SecondName,
		Balance:    m.Balance,
		Role:       m.Role,
		CreatedAt:  m.CreatedAt,
		UpdatedAt:  m.UpdatedAt,
	}
}
