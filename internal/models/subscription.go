package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Status string
type Type string

const (
	Active   Status = "active"
	InActive Status = "inactive"
	Deleted  Status = "deleted"
)

const (
	Individual Type = "individual"
	Group      Type = "group"
)

type Subscription struct {
	Id                      uint      `json:"id"`
	OwnerId                 uint      `json:"owner_id"`
	Name                    string    `json:"name"`
	Amount                  float64   `json:"amount"`
	Payment_date            time.Time `json:"payment_date"`
	Payment_interval        uint      `json:"payment_interval"`
	Status                  Status    `json:"status"`
	Type                    Type      `json:"type"`
	DaysNotifyBeforePayment uint      `json:"days_notify_before_payment"`
	Created_at              time.Time `json:"created_at"`
	Updated_at              time.Time `json:"updated_at"`
	Image                   []byte    `json:"image"`
	Color                   string    `json:"color"`
	InvateToken             string    `json:"invate_token"`
}

func (sscpt Subscription) Validate() error {
	return validation.ValidateStruct(&sscpt,
		validation.Field(&sscpt.OwnerId, validation.Required),
		validation.Field(&sscpt.Name, validation.Required),
		validation.Field(&sscpt.Amount, validation.Required),
		validation.Field(&sscpt.Payment_date, validation.Required),
		validation.Field(&sscpt.Payment_interval, validation.Required),
		validation.Field(&sscpt.Type, validation.Required),
		validation.Field(&sscpt.Color, validation.Required))
}
