package models

import "time"

type Candidate struct {
	ID        uint64    `gorm:"primary_key;auto_increment;not_null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Name      string    `json:"name" gorm:"not null"`
	User      User      `json:"user" gorm:"not null"`
	UserID    uint64
	pAddress  Address	`gorm:"polymorphic:Owner"`
	cAddress  Address	`gorm:"polymorphic:Owner"`
}