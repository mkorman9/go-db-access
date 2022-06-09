package main

import (
	"strings"
	"time"
)

type Client struct {
	ID          string     `json:"id"`
	Gender      string     `json:"gender"`
	FirstName   string     `json:"firstName"`
	LastName    string     `json:"lastName"`
	Address     string     `json:"address,omitempty"`
	PhoneNumber string     `json:"phoneNumber,omitempty"`
	Email       string     `json:"email,omitempty"`
	BirthDate   *time.Time `json:"birthDate,omitempty"`
	CreditCards []string   `json:"creditCards"`
	IsDeleted   bool       `json:"-"`
}

func (c *Client) ToEntity() *ClientEntity {
	return &ClientEntity{
		ID:                c.ID,
		Gender:            c.Gender,
		FirstName:         c.FirstName,
		LastName:          c.LastName,
		Address:           c.Address,
		PhoneNumber:       c.PhoneNumber,
		Email:             c.Email,
		BirthDate:         c.BirthDate,
		CreditCardsString: strings.Join(c.CreditCards, ";"),
		IsDeleted:         c.IsDeleted,
	}
}
