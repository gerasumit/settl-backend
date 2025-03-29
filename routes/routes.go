package routes

import (
	"net/http"
    "github.com/gin-gonic/gin"
    "settl-backend/controllers"
)

func SetupRoutes(router *gin.Engine) {
    api := router.Group("/api")
    {
        auth := api.Group("/auth")
        {
            auth.POST("/google-login", controllers.GoogleLogin)
            auth.POST("/logout", controllers.Logout)
        }

        // groups := api.Group("/groups")
        // {
        //     groups.POST("/", controllers.CreateGroup)
        //     groups.GET("/", controllers.GetAllGroups)
        //     groups.GET("/:groupId", controllers.GetGroupDetails)
        //     groups.DELETE("/:groupId", controllers.DeleteGroup)
        // }

        // expenses := api.Group("/groups/:groupId/expenses")
        // {
        //     expenses.POST("/", controllers.AddExpense)
        //     expenses.GET("/", controllers.GetAllExpenses)
        //     expenses.GET("/:expenseId", controllers.GetExpenseDetails)
        //     expenses.PUT("/:expenseId", controllers.EditExpense)
        //     expenses.DELETE("/:expenseId", controllers.DeleteExpense)
        // }

		api.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		})
    }
}