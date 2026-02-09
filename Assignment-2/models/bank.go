package models

import "time"

type Bank struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:255;not null" json:"name"`
	Code      string    `gorm:"size:100;unique;not null" json:"code"`
	Branches  []Branch  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"branches"`
	CreatedAt time.Time `json:"created_at"`
}
