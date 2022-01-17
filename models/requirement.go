package models

import "time"

type Requirement struct {
	ID        		uint64    `gorm:"primary_key;auto_increment;not_null"`
	CreatedAt 		time.Time `gorm:"autoCreateTime"`
	UpdatedAt 		time.Time `gorm:"autoUpdateTime"`
	ClientPocID    int
	ClientPoc      ClientPoc
	JobDescription	string
	TotalOpenings	int
	CurrentOpenings	int
	MinSalary		int
	MaxSalary		int
	MinExp			int
	MaxExp			int
	MaxJoiningDate	time.Time
	Role			string
	Skill			string
	OwnerID			uint64
	Owner			User	`gorm:"foreignKey:OwnerID"`

	PaymentType		string
	InvoicePercentage	int
	InvoiceAmount	int
	GstPercentage	int
}

type CandidateOppurtunity struct {
	ID        		uint64    	`gorm:"primary_key;auto_increment;not_null"`
	CreatedAt 		time.Time 	`gorm:"autoCreateTime"`
	UpdatedAt 		time.Time 	`gorm:"autoUpdateTime"`

	CandidateID		uint64
	Candidate		Candidate

	RequirementID	uint64
	Requirement 	Requirement

	RecruiterID		uint64
	Recruiter		User		`gorm:"foreignKey:RecruiterID"`

	PurchaseOrderID	uint64
	PurchaseOrder   PurchaseOrder

	Stage			string
	State			string

	
	InvoiceID		uint64
	Invoice			Invoice

	CTC				int
	JoiningDate     time.Time

	Joined 			bool
	ClientCandidateID		int
}

