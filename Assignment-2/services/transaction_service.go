package services

import (
	"errors"
	"time"

	"github.com/kritika-bipl/Assignment/Assignment-2/models"
	"gorm.io/gorm"
)

func Deposit(db *gorm.DB, accountID uint, amount float64) error {

	//check amount
	if amount <= 0 {
		return errors.New("deposit amount must be greater than zero")
	}

	//now do 2 works
	//1. update kro account balance
	//2. create transaction record

	//start transaction
	return db.Transaction(func(tx *gorm.DB) error {
		//get account from db by accid
		var acc models.Account
		if err := tx.First(&acc, accountID).Error; err != nil {
			return err
		}

		//increase balance
		acc.Balance += amount

		//save kro updated account
		if err := tx.Save(&acc).Error; err != nil {
			return err
		}

		//create transaction
		tr := models.Transaction{
			AccountID: accountID,
			Type:      "deposit",
			Amount:    amount,
			CreatedAt: time.Now().UTC(),
		}
		if err := tx.Create(&tr).Error; err!=nil {
			return err
		}
		return nil
	})
}

func Withdraw(db *gorm.DB, accountID uint, amount float64) error {
	if amount <= 0 {
		return errors.New("withdraw amount must be positive")
	}
	return db.Transaction(func(tx *gorm.DB) error {
		var acc models.Account
		if err := tx.First(&acc, accountID).Error; err != nil {
			return err
		}
		if acc.Balance < amount {
			return errors.New("insufficient balance")
		}
		acc.Balance -= amount
		if err := tx.Save(&acc).Error; err != nil {
			return err
		}
		tr := models.Transaction{
			AccountID: accountID,
			Type:      "withdraw",
			Amount:    amount,
			CreatedAt: time.Now().UTC(),
		}
		if err := tx.Create(&tr).Error; err != nil {
			return err
		}
		return nil
	})
}

