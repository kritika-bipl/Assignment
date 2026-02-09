package models

import "time"

type Branch struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	BankID    uint      `gorm:"not null" json:"bank_id"`
	Bank      Bank      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Name      string    `gorm:"size:255;not null" json:"name"`
	IFSCCode  string    `gorm:"size:100;not null;unique" json:"ifsc_code"`
	Address   string    `gorm:"type:text" json:"address"`
	CreatedAt time.Time `json:"created_at"`
}
