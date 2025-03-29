package controllers

// import (
//     "net/http"
//     "settl-backend/services"
//     "github.com/gin-gonic/gin"
// )

// func GetGroupBalance(c *gin.Context) {
//     groupID := c.Param("groupId")
//     balances, err := services.CalculateGroupBalance(groupID)
//     if err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to calculate balances"})
//         return
//     }
//     c.JSON(http.StatusOK, gin.H{"balances": balances})
// }
