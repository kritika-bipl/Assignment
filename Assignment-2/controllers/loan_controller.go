package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kritika-bipl/Assignment/Assignment-2/models"
	"github.com/kritika-bipl/Assignment/Assignment-2/services"
	"gorm.io/gorm"
)

type applyLoanRequest struct {
	CustomerID      uint    `json:"customer_id" binding:"required"`
	BranchID        uint    `json:"branch_id" binding:"required"`
	PrincipalAmount float64 `json:"principal_amount" binding:"required,gt=0"`
	DurationYears   float64 `json:"duration_years" binding:"required,gt=0"`
}

func ApplyLoanHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var r applyLoanRequest
		if err := c.ShouldBindJSON(&r); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		loan := models.Loan{
			CustomerID:      r.CustomerID,
			BranchID:        r.BranchID,
			PrincipalAmount: r.PrincipalAmount,
		}
		if err := services.ApplyLoan(db, &loan, r.DurationYears); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, loan)
	}
}

func GetLoansByCustomerHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("customerId")
		id, _ := strconv.Atoi(idStr)
		loans, err := services.GetLoansByCustomer(db, uint(id))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, loans)
	}
}

func GetLoanHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("loanId")
		id, _ := strconv.Atoi(idStr)
		loan, err := services.GetLoan(db, uint(id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, loan)
	}
}

func RepayLoanHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var r struct {
			LoanID uint    `json:"loan_id" binding:"required"`
			Amount float64 `json:"amount" binding:"required,gt=0"`
		}
		if err := c.ShouldBindJSON(&r); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := services.RepayLoan(db, r.LoanID, r.Amount); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "repayment successful"})
	}
}

func InterestThisYearHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("loanId")
		id, _ := strconv.Atoi(idStr)
		interest, err := services.InterestThisYear(db, uint(id))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"interest_this_year": interest})
	}
}

func PendingAmountHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("loanId")
		id, _ := strconv.Atoi(idStr)
		pending, err := services.PendingAmount(db, uint(id))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"pending_amount": pending})
	}
}

func UpdateLoanHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("loanId")
		id, _ := strconv.Atoi(idStr)
		var payload struct {
			Status string `json:"status"`
		}
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var loan models.Loan
		if err := db.First(&loan, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "loan not found"})
			return
		}
		if payload.Status != "" {
			loan.Status = payload.Status
			if payload.Status == "closed" && loan.RemainingAmount <= 0 {
				now := time.Now().UTC()
				loan.EndDate = &now
			}
		}
		if err := db.Save(&loan).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, loan)
	}
}

func DeleteLoanHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("loanId")
		id, _ := strconv.Atoi(idStr)
		var loan models.Loan
		if err := db.First(&loan, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "loan not found"})
			return
		}
		if loan.RemainingAmount > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "cannot delete active loan with pending amount"})
			return
		}
		if err := db.Delete(&loan).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "loan deleted"})
	}
}
