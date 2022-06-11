package main

import (
	"github.com/lib/pq"
	"time"
)

type Account struct {
	ID          string         `json:"id"`
	Username    string         `json:"username"`
	Roles       pq.StringArray `json:"roles"`
	IsDeleted   bool           `json:"isDeleted"`
	BannedUntil *time.Time     `json:"bannedUntil,omitempty"`

	Credentials *Credentials `json:"credentials"`
}

type Credentials struct {
	AccountID      string `json:"-"`
	Email          string `json:"email"`
	PasswordBcrypt string `json:"passwordBcrypt"`
}

func (a *Account) ToEntity() *AccountEntity {
	ae := &AccountEntity{
		ID:          a.ID,
		Username:    a.Username,
		Roles:       a.Roles,
		IsDeleted:   a.IsDeleted,
		BannedUntil: a.BannedUntil,
	}

	if a.Credentials != nil {
		ae.Credentials = &CredentialsEntity{
			AccountID:      a.ID,
			Email:          a.Credentials.Email,
			PasswordBcrypt: a.Credentials.PasswordBcrypt,
		}
	}

	return ae
}

type AccountEntity struct {
	ID          string             `gorm:"column:id; type:uuid; primaryKey"`
	Username    string             `gorm:"column:username; type:text; unique"`
	Roles       pq.StringArray     `gorm:"column:roles; type:text[]"`
	IsDeleted   bool               `gorm:"column:deleted; type:boolean"`
	BannedUntil *time.Time         `gorm:"column:banned_until; type:timestamp"`
	Credentials *CredentialsEntity `gorm:"foreignKey:account_id; constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}

func (AccountEntity) TableName() string {
	return "accounts"
}

type CredentialsEntity struct {
	AccountID      string `gorm:"column:account_id; type:uuid; primaryKey"`
	Email          string `gorm:"column:email; type:text; unique"`
	PasswordBcrypt string `gorm:"column:password_bcrypt; type:text"`
}

func (CredentialsEntity) TableName() string {
	return "credentials"
}

func (ae *AccountEntity) ToAccount() *Account {
	a := &Account{
		ID:          ae.ID,
		Username:    ae.Username,
		Roles:       ae.Roles,
		IsDeleted:   ae.IsDeleted,
		BannedUntil: ae.BannedUntil,
	}

	if ae.Credentials != nil {
		a.Credentials = &Credentials{
			Email:          ae.Credentials.Email,
			PasswordBcrypt: ae.Credentials.PasswordBcrypt,
		}
	}
	return a
}
