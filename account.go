package main

import (
	"github.com/lib/pq"
	"time"
)

type Account struct {
	ID          string         `json:"id" gorm:"column:id; type:uuid; primaryKey"`
	Username    string         `json:"username" gorm:"column:username; type:text; unique; not null"`
	Roles       pq.StringArray `json:"roles" gorm:"column:roles; type:text[]; not null"`
	IsDeleted   bool           `json:"isDeleted" gorm:"column:deleted; type:boolean; not null; default:false"`
	BannedUntil *time.Time     `json:"bannedUntil,omitempty" gorm:"column:banned_until; type:timestamp"`

	Credentials *Credentials `json:"credentials" gorm:"foreignKey:account_id; constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}

func (Account) TableName() string {
	return "accounts"
}

type Credentials struct {
	AccountID      string `json:"-" gorm:"column:account_id; type:uuid; primaryKey"`
	Email          string `json:"email" gorm:"column:email; type:text; unique; not null"`
	PasswordBcrypt string `json:"passwordBcrypt" gorm:"column:password_bcrypt; type:text; not null"`
}

func (Credentials) TableName() string {
	return "credentials"
}
