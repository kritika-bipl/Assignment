package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/kritika-bipl/Assignment/Assignment-2/models"
	"gorm.io/gorm"
)

func OpenAccount(db *gorm.DB, acc *models.Account) error {
	if acc.CustomerID == 0 || acc.BranchID == 0 {
		return errors.New("customer_id and branch_id are required")
	}

	acc.AccountNumber = fmt.Sprintf("AC-%s", uuid.New().String())
	acc.Type = "savings"
	acc.Status = "active"
	acc.Balance = 0
	acc.CreatedAt = time.Now().UTC()

	return db.Create(acc).Error
}

func GetAccountWithTransactions(db *gorm.DB, accountID uint) (models.Account, []models.Transaction, error) {
	var acc models.Account
	if err := db.Preload("Transactions").First(&acc, accountID).Error; err != nil {
		return models.Account{}, nil, err
	}
	return acc, acc.Transactions, nil
}

func GetAccountsByCustomer(db *gorm.DB, customerID uint) ([]models.Account, error) {
	var accs []models.Account
	if err := db.Where("customer_id = ?", customerID).Find(&accs).Error; err != nil {
		return nil, err
	}
	return accs, nil
}


