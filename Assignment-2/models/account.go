package models

import "time"

type Account struct {
	ID            uint          `gorm:"primaryKey" json:"id"`
	CustomerID    uint          `gorm:"not null" json:"customer_id"`
	Customer      Customer      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	BranchID      uint          `gorm:"not null" json:"branch_id"`
	Branch        Branch        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	AccountNumber string        `gorm:"size:100;unique;not null" json:"account_number"`
	Type          string        `gorm:"size:50;not null;default:'savings'" json:"type"`
	Balance       float64       `gorm:"type:numeric;default:0" json:"balance"`
	Status        string        `gorm:"size:50;default:'active'" json:"status"`
	Transactions  []Transaction `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"transactions"`
	CreatedAt     time.Time     `json:"created_at"`
}
