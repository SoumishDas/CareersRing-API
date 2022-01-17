package models

import "time"

type ProcessInstance struct {
	ID        uint64    `gorm:"primary_key;auto_increment;not_null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	BusinessKey string
	ProcessID   string
	ProcessType string
	Owner string
	OwnerType string
}