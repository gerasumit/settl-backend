package controllers

// import (
//     "net/http"
//     "settl-backend/services"
//     "github.com/gin-gonic/gin"
// )

// func CreateGroup(c *gin.Context) {
//     var request services.GroupRequest
//     if err := c.ShouldBindJSON(&request); err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
//         return
//     }
//     group, err := services.CreateGroup(request)
//     if err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create group"})
//         return
//     }
//     c.JSON(http.StatusOK, gin.H{"group": group})
// }