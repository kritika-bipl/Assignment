package models

import "time"

type Loan struct {
	ID              uint          `gorm:"primaryKey" json:"id"`
	CustomerID      uint          `gorm:"not null" json:"customer_id"`
	Customer        Customer      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	BranchID        uint          `gorm:"not null" json:"branch_id"`
	Branch          Branch        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	PrincipalAmount float64       `gorm:"type:numeric;not null" json:"principal_amount"`
	InterestRate    float64       `gorm:"type:numeric;default:12" json:"interest_rate"` // percent per year
	TotalInterest   float64       `gorm:"type:numeric;default:0" json:"total_interest"`
	AmountPaid      float64       `gorm:"type:numeric;default:0" json:"amount_paid"`
	RemainingAmount float64       `gorm:"type:numeric;default:0" json:"remaining_amount"`
	StartDate       time.Time     `json:"start_date"`
	EndDate         *time.Time    `json:"end_date"`
    Status          string        `gorm:"size:50;default:'active'" json:"status"`
	Payments        []LoanPayment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"payments"`
	CreatedAt       time.Time     `json:"created_at"`
}
