package main

import (
	"time"
)

type Gender = string

const (
	GenderUndefined = Gender("-")
	GenderMale      = Gender("M")
	GenderFemale    = Gender("F")
)

type Client struct {
	ID          string     `json:"id"`
	Gender      Gender     `json:"gender"`
	FirstName   string     `json:"firstName"`
	LastName    string     `json:"lastName"`
	Address     string     `json:"address,omitempty"`
	PhoneNumber string     `json:"phoneNumber,omitempty"`
	Email       string     `json:"email,omitempty"`
	BirthDate   *time.Time `json:"birthDate,omitempty"`
	CreditCards []string   `json:"creditCards"`
	IsDeleted   bool       `json:"-"`
}

type Clients []*Client

func (c *Client) ToEntity() *ClientEntity {
	return &ClientEntity{
		ID:          c.ID,
		Gender:      c.Gender,
		FirstName:   c.FirstName,
		LastName:    c.LastName,
		Address:     c.Address,
		PhoneNumber: c.PhoneNumber,
		Email:       c.Email,
		BirthDate:   c.BirthDate,
		CreditCards: c.CreditCards,
		IsDeleted:   c.IsDeleted,
	}
}

func (cs Clients) ToEntities() ClientEntities {
	entities := make([]*ClientEntity, len(cs))
	for i, entity := range cs {
		entities[i] = entity.ToEntity()
	}

	return entities
}
