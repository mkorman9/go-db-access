package main

import (
	"github.com/lib/pq"
	"time"
)

type Session struct {
	ID        string     `json:"id"`
	AccountID string     `json:"accountId"`
	Token     string     `json:"token"`
	Roles     []string   `json:"roles"`
	ExpiresAt *time.Time `json:"expiresAt"`

	Account *Account `json:"-"`
}

func (s *Session) ToEntity() *SessionEntity {
	return &SessionEntity{
		ID:        s.ID,
		AccountID: s.AccountID,
		Token:     s.Token,
		Roles:     s.Roles,
		ExpiresAt: s.ExpiresAt,
		Account:   s.Account.ToEntity(),
	}
}

type SessionEntity struct {
	ID        string         `gorm:"column:id; type:text; primaryKey"`
	AccountID string         `gorm:"column:account_id; type:uuid"`
	Token     string         `gorm:"column:token; type:text; unique"`
	Roles     pq.StringArray `gorm:"column:roles; type:text[]"`
	ExpiresAt *time.Time     `gorm:"column:expires_at; type:timestamp"`

	Account *AccountEntity `gorm:"foreignKey:AccountID"`
}

func (SessionEntity) TableName() string {
	return "sessions"
}

func (se *SessionEntity) ToSession() *Session {
	return &Session{
		ID:        se.ID,
		AccountID: se.AccountID,
		Token:     se.Token,
		Roles:     se.Roles,
		ExpiresAt: se.ExpiresAt,
		Account:   se.Account.ToAccount(),
	}
}
