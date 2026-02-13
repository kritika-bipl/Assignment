package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kritika-bipl/Assignment/Assignment-2/controllers"
	"gorm.io/gorm"
)


func RegisterRoutes(r *gin.Engine , db *gorm.DB){

	api := r.Group("/api")
	{    
		//bank api
		api.POST("/banks", controllers.CreateBank(db))
		api.GET("/banks", controllers.ListBanks(db))

		//branches
		api.POST("/branches", controllers.CreateBranch(db))
		api.GET("/branches/:bankId", controllers.ListBranchesByBank(db))

		// Customers
		api.POST("/customers", controllers.CreateCustomer(db))
		api.GET("/customers/:id", controllers.GetCustomer(db))

		//accounts
		api.POST("/accounts/open", controllers.OpenAccountHandler(db))
		api.GET("/accounts/:id", controllers.GetAccountHandler(db))
		api.GET("/accounts/customer/:customerId", controllers.GetAccountsByCustomerHandler(db))

		// Transactions
		api.POST("/accounts/deposit", controllers.DepositHandler(db))
		api.POST("/accounts/withdraw", controllers.WithdrawHandler(db))
		api.GET("/accounts/:id/transactions", controllers.ListTransactions(db))

		// Loans
		api.POST("/loans/apply", controllers.ApplyLoanHandler(db))
		api.GET("/loans/customer/:customerId", controllers.GetLoansByCustomerHandler(db))
		api.GET("/loans/:loanId", controllers.GetLoanHandler(db))

		// Loan payments
		api.POST("/loans/repay", controllers.RepayLoanHandler(db))

		// Reporting
		api.GET("/loans/:loanId/interest-this-year", controllers.InterestThisYearHandler(db))
		api.GET("/loans/:loanId/pending-amount", controllers.PendingAmountHandler(db))



	}
}