package models

import "time"

type LoanPayment struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	LoanID    uint      `gorm:"not null" json:"loan_id"`
	Loan      Loan      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Amount    float64   `gorm:"type:numeric;not null" json:"amount"`
	PaidAt    time.Time `json:"payment_date"`
}
