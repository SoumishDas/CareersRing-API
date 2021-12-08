package models

import "time"

type Address struct {
	ID        uint64    `gorm:"primary_key;auto_increment;not_null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Name      string    `json:"name" gorm:"not null"`
	OwnerID   int
    OwnerType string
}
