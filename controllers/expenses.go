package controllers

// import (
//     "net/http"
//     "settl-backend/services"
//     "github.com/gin-gonic/gin"
// )

// func AddExpense(c *gin.Context) {
//     var request services.ExpenseRequest
//     if err := c.ShouldBindJSON(&request); err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
//         return
//     }
//     expense, err := services.AddExpense(request)
//     if err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add expense"})
//         return
//     }
//     c.JSON(http.StatusOK, gin.H{"expense": expense})
// }
