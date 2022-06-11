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
	se := &SessionEntity{
		ID:        s.ID,
		AccountID: s.AccountID,
		Token:     s.Token,
		Roles:     s.Roles,
		ExpiresAt: s.ExpiresAt,
	}

	if s.Account != nil {
		se.Account = s.Account.ToEntity()
	}

	return se
}

type SessionEntity struct {
	ID        string         `gorm:"column:id; type:text; primaryKey"`
	AccountID string         `gorm:"column:account_id; type:uuid; not null"`
	Token     string         `gorm:"column:token; type:text; unique; not null"`
	Roles     pq.StringArray `gorm:"column:roles; type:text[]; not null"`
	ExpiresAt *time.Time     `gorm:"column:expires_at; type:timestamp"`

	Account *AccountEntity `gorm:"foreignKey:AccountID"`
}

func (SessionEntity) TableName() string {
	return "sessions"
}

func (se *SessionEntity) ToSession() *Session {
	s := &Session{
		ID:        se.ID,
		AccountID: se.AccountID,
		Token:     se.Token,
		Roles:     se.Roles,
		ExpiresAt: se.ExpiresAt,
	}

	if se.Account != nil {
		s.Account = se.Account.ToAccount()
	}

	return s
}
