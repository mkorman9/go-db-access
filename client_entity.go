package main

import (
	"github.com/lib/pq"
	"time"
)

type ClientEntity struct {
	ID          string         `gorm:"column:id; type:uuid; primaryKey"`
	Gender      string         `gorm:"column:gender; type:char(1)"`
	FirstName   string         `gorm:"column:first_name; type:varchar(255)"`
	LastName    string         `gorm:"column:last_name; type:varchar(255)"`
	Address     string         `gorm:"column:home_address; type:varchar(1024)"`
	PhoneNumber string         `gorm:"column:phone_number; type:varchar(64)"`
	Email       string         `gorm:"column:email; type:varchar(64)"`
	BirthDate   *time.Time     `gorm:"column:birth_date; type:timestamp"`
	CreditCards pq.StringArray `gorm:"column:credit_cards; type:text[]"`
	IsDeleted   bool           `gorm:"column:deleted; type:boolean"`
}

func (ClientEntity) TableName() string {
	return "clients"
}

func (e *ClientEntity) ToClient() *Client {
	return &Client{
		ID:          e.ID,
		Gender:      e.Gender,
		FirstName:   e.FirstName,
		LastName:    e.LastName,
		Address:     e.Address,
		PhoneNumber: e.PhoneNumber,
		Email:       e.Email,
		BirthDate:   e.BirthDate,
		CreditCards: e.CreditCards,
		IsDeleted:   e.IsDeleted,
	}
}
