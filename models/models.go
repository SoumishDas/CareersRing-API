package models

import "github.com/jinzhu/gorm"

func MigrateDB(db *gorm.DB){
	db.AutoMigrate(
		&Item{},
		&User{},
		&ProcessInstance{},
		&Company{},
		&CompanyBankAccount{},
		&Client{},
		&ClientPoc{},
		&ClientBillingAddress{},
		&Address{},
		&Candidate{},
		&Requirement{},
		&PurchaseOrder{},
		&CashPosting{},
		&Invoice{},
		&InvoiceItem{},
		&Remittance{},
		&CandidateOppurtunity{},
		
	)
}