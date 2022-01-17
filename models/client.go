package models

import "time"

type Client struct {
	ID        		uint64    `gorm:"primary_key;auto_increment;not_null"`
	CreatedAt 		time.Time `gorm:"autoCreateTime"`
	UpdatedAt 		time.Time `gorm:"autoUpdateTime"`
	ClientName 		string
	LegalName		string  

	AgreementLink	string
	AgreementStart 	time.Time
	AgreementEnd	time.Time

	CompanyID		uint64
	Company			Company

	PaymentTerm		string
}

type ClientPoc struct{
	ID        		uint64    `gorm:"primary_key;auto_increment;not_null"`
	CreatedAt 		time.Time `gorm:"autoCreateTime"`
	UpdatedAt 		time.Time `gorm:"autoUpdateTime"`
	ClientID 		int
	Client			Client
	Name	  		string
	PhoneNo   		string
	BusinessUnit	string
	LineOfBusiness	string
}

type PurchaseOrder struct{
	ID        		uint64    `gorm:"primary_key;auto_increment;not_null"`
	CreatedAt 		time.Time `gorm:"autoCreateTime"`
	UpdatedAt 		time.Time `gorm:"autoUpdateTime"`
	ClientPocID    int
	ClientPoc      ClientPoc
	TotalAmount		int
	AvailableAmount int
	PoDate			time.Time
	PoExpiryDate	time.Time
}

type ClientBillingAddress struct {
	ID        	uint64    `gorm:"primary_key;auto_increment;not_null"`
	CreatedAt 	time.Time `gorm:"autoCreateTime"`
	UpdatedAt 	time.Time `gorm:"autoUpdateTime"`

	ClientID    int
	Client      Client

	Address		string

	Pincode		uint
	GstNo		string

	IsSEZ		bool

}

