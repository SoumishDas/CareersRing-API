package models

import "time"

type Invoice struct {
	ID        uint64    `gorm:"primary_key;auto_increment;not_null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	Number 		int
	Date		time.Time
	Type  		string

	CompanyID	uint64
	Company		Company

	PurchaseOrderID	uint64
	PurchaseOrder 	PurchaseOrder
	Note			string

	CashPostingID	uint64
	CashPosting		CashPosting

	Status		string
	

}


type InvoiceItem struct {
	ID        		uint64    	`gorm:"primary_key;auto_increment;not_null"`
	CreatedAt 		time.Time 	`gorm:"autoCreateTime"`
	UpdatedAt 		time.Time 	`gorm:"autoUpdateTime"`

	InvoiceID		uint64
	Invoice			Invoice

	PurchaseOrderID	uint64
	PurchaseOrder	PurchaseOrder

	CandidateOppurtunityID	uint64
	CandidateOppurtunity	CandidateOppurtunity

	Item			string
	Quantity		int
	Rate			int
}



type CashPosting struct {
	ID        		uint64    	`gorm:"primary_key;auto_increment;not_null"`
	CreatedAt 		time.Time 	`gorm:"autoCreateTime"`
	UpdatedAt 		time.Time 	`gorm:"autoUpdateTime"`

	TotalAmount		int		
	TransactionID	string
	
	ClientBillingAddressID	int
	ClientBillingAddress	ClientBillingAddress

	CompanyBankAccountID	uint64
	CompanyBankAccount		CompanyBankAccount

	PaymentDate		time.Time
	Comment 		string
}

type Remittance struct {
	ID        		uint64    	`gorm:"primary_key;auto_increment;not_null"`
	CreatedAt 		time.Time 	`gorm:"autoCreateTime"`
	UpdatedAt 		time.Time 	`gorm:"autoUpdateTime"`

	CashPostingID	uint64
	CashPosting		CashPosting


}