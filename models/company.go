package models

import "time"

type Company struct {
	ID        uint64    `gorm:"primary_key;auto_increment;not_null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	Name			string
	LegalName		string
	ContactNo		string
	AltContactNo	string
	GstNo			string

	Address			string
}

type CompanyBankAccount struct {
	ID        uint64    `gorm:"primary_key;auto_increment;not_null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	CompanyID	uint64
	Company 	Company

	Bank		string
	Branch		string
	IfscCode	string
	AccNo 		string
	AccType		string
	
}