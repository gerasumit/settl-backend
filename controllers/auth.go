package controllers

import (
//     "net/http"
//     "settl-backend/services"
    "github.com/gin-gonic/gin"
)

func GoogleLogin(c *gin.Context) {
//     var request services.GoogleLoginRequest
//     if err := c.ShouldBindJSON(&request); err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
//         return
//     }
//     token, err := services.HandleGoogleLogin(request)
//     if err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to authenticate"})
//         return
//     }
//     c.JSON(http.StatusOK, gin.H{"token": token})
}

func Logout(c *gin.Context) {
//     err := services.HandleLogout(c)
//     if err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to log out"})
//         return
//     }
//     c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
