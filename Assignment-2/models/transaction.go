package models

import "time"

type Transaction struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	AccountID uint      `gorm:"not null" json:"account_id"`
	Account   Account   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Type      string    `gorm:"size:50;not null" json:"type"` // deposit, withdraw, loan_repay
	Amount    float64   `gorm:"type:numeric;not null" json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}
