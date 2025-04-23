package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type PaymentStatus string

const (
	Paid    PaymentStatus = "paid"
	NotPaid PaymentStatus = "not paid"
)

type SubscriptionMember struct {
	Id                      uint          `json:"id"`
	UserId                  uint          `json:"user_id"`
	SubscriptionId          uint          `json:"subscription_id"`
	Amount                  float64       `json:"amount"`
	Joined_at               time.Time     `json:"joined_at"`
	DaysNotifyBeforePayment uint          `json:"days_notify_before_payment"`
	Color                   string        `json:"color"`
	Payment_status          PaymentStatus `json:"payment_status"`
}

func (sscptm SubscriptionMember) Validate() error {
	return validation.ValidateStruct(&sscptm,
		validation.Field(&sscptm.UserId, validation.Required),
		validation.Field(&sscptm.SubscriptionId, validation.Required),
		validation.Field(&sscptm.Amount, validation.Required))
}
