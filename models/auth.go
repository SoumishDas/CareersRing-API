package models

import "time"

type User struct {
	ID        uint64    `gorm:"primary_key;auto_increment;not_null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Email     string    `json:"email" gorm:"not null;unique"`
	Password  string    `json:"password" gorm:"not null"`
	FirstName string    `json:"firstName" `
	LastName  string    `json:"lastName" `
	IsActive  bool      `json:"isActive" gorm:"not null;default:true"`
}