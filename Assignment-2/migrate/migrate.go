package main

import (
	"fmt"

	"github.com/kritika-bipl/Assignment/Assignment-2/initializers"
	"github.com/kritika-bipl/Assignment/Assignment-2/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(
		&models.Bank{},
		&models.Branch{},
		&models.Customer{},
		&models.Account{},
		&models.Transaction{},
		&models.Loan{},
		&models.LoanPayment{},
	)

	fmt.Println("Database migration completed successfully!")
}
