package main

import (
	"github.com/lib/pq"
	"time"
)

type ClientEntity struct {
	ID          string         `gorm:"column:id; type:uuid; primaryKey"`
	Gender      string         `gorm:"column:gender; type:char(1)"`
	FirstName   string         `gorm:"column:first_name; type:text"`
	LastName    string         `gorm:"column:last_name; type:text"`
	Address     string         `gorm:"column:home_address; type:text"`
	PhoneNumber string         `gorm:"column:phone_number; type:text"`
	Email       string         `gorm:"column:email; type:text"`
	BirthDate   *time.Time     `gorm:"column:birth_date; type:timestamp"`
	CreditCards pq.StringArray `gorm:"column:credit_cards; type:text[]"`
	IsDeleted   bool           `gorm:"column:deleted; type:boolean"`
}

type ClientEntities []*ClientEntity

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

func (es ClientEntities) ToClients() Clients {
	clients := make([]*Client, len(es))
	for i, entity := range es {
		clients[i] = entity.ToClient()
	}

	return clients
}
