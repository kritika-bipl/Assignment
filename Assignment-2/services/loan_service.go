package services

import (
	"errors"
	"math"
	"time"

	"github.com/kritika-bipl/Assignment/Assignment-2/models"
	"gorm.io/gorm"
)

const InterestRate = 12.0 

func ApplyLoan(db *gorm.DB, loan *models.Loan, durationYears float64) error {
	if loan.PrincipalAmount <= 0 {
		return errors.New("principal must be positive")
	}
	loan.InterestRate = InterestRate
	loan.TotalInterest = calculateInterest(loan.PrincipalAmount, loan.InterestRate, durationYears)
	loan.AmountPaid = 0
	loan.RemainingAmount = loan.PrincipalAmount + loan.TotalInterest
	loan.StartDate = time.Now().UTC()
	loan.Status = "active"
	return db.Create(loan).Error
}

func calculateInterest(principal float64, ratePercent float64, years float64) float64 {
	rate := ratePercent / 100.0
	interest := principal * rate * years
	// round to 2 decimals
	return math.Round(interest*100) / 100
}

func RepayLoan(db *gorm.DB, loanID uint, amount float64) error {
	if amount <= 0 {
		return errors.New("repayment amount must be positive")
	}
	return db.Transaction(func(tx *gorm.DB) error {
		var loan models.Loan
		if err := tx.First(&loan, loanID).Error; err != nil {
			return err
		}
		if loan.RemainingAmount <= 0 {
			return errors.New("loan already paid")
		}
		paid := amount
		if amount > loan.RemainingAmount {
			paid = loan.RemainingAmount
		}
		loan.AmountPaid += paid
		loan.RemainingAmount -= paid
		if loan.RemainingAmount <= 0 {
			loan.Status = "closed"
			now := time.Now().UTC()
			loan.EndDate = &now
		}
		if err := tx.Save(&loan).Error; err != nil {
			return err
		}
		lp := models.LoanPayment{
			LoanID:    loan.ID,
			Amount:    paid,
			PaidAt:    time.Now().UTC(),
		}
		if err := tx.Create(&lp).Error; err != nil {
			return err
		}
		return nil
	})
}

func GetLoansByCustomer(db *gorm.DB, customerID uint) ([]models.Loan, error) {
	var loans []models.Loan
	if err := db.Where("customer_id = ?", customerID).Preload("Payments").Find(&loans).Error; err != nil {
		return nil, err
	}
	return loans, nil
}

func GetLoan(db *gorm.DB, loanID uint) (models.Loan, error) {
	var loan models.Loan
	if err := db.Preload("Payments").First(&loan, loanID).Error; err != nil {
		return models.Loan{}, err
	}
	return loan, nil
}

func InterestThisYear(db *gorm.DB, loanID uint) (float64, error) {
	loan, err := GetLoan(db, loanID)
	if err != nil {
		return 0, err
	}

	rate := loan.InterestRate / 100 

	interest := loan.PrincipalAmount * rate
	return interest, nil
}

func PendingAmount(db *gorm.DB, loanID uint) (float64, error) {
	loan, err := GetLoan(db, loanID)
	if err != nil {
		return 0, err
	}
	return loan.RemainingAmount, nil
}
