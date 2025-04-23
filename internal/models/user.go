package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type User struct {
	Id           uint   `json:"id"`
	Login        string `json:"name"`
	Password     string `json:"password"`
	Email        string `json:"email"`
	DateBirthday string `json:"date_birthday"`
	Created_at   string `json:"created_at"`
}

func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Login, validation.Required),
		validation.Field(&u.Email, validation.Required),
		validation.Field(&u.Password, validation.Required),
		validation.Field(&u.DateBirthday, validation.Required))
}
