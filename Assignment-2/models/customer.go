package models

import "time"

type Customer struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:255;not null" json:"name"`
	Email     string    `gorm:"size:255;unique;not null" json:"email"`
	Phone     string    `gorm:"size:50" json:"phone"`
	Address   string    `gorm:"type:text" json:"address"`
	Accounts  []Account `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"accounts"`
	Loans     []Loan    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"loans"`
	CreatedAt time.Time `json:"created_at"`
}
