package main

import (
	"github.com/lib/pq"
	"time"
)

type Session struct {
	ID        string         `json:"id" gorm:"column:id; type:text; primaryKey"`
	AccountID string         `json:"accountId" gorm:"column:account_id; type:uuid; not null"`
	Token     string         `json:"token" gorm:"column:token; type:text; unique; not null"`
	Roles     pq.StringArray `json:"roles" gorm:"column:roles; type:text[]; not null"`
	ExpiresAt *time.Time     `json:"expiresAt,omitempty" gorm:"column:expires_at; type:timestamp"`

	Account *Account `json:"-" gorm:"foreignKey:AccountID"`
}

func (Session) TableName() string {
	return "sessions"
}
