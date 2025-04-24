package routes

import (
    "github.com/gin-gonic/gin"
    "backend/controller"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    poll := r.Group("/poll")
    {
        poll.POST("/create", controllers.CreatePoll)
        poll.GET("/:id", controllers.ViewPoll)
        poll.PUT("/close/:id", controllers.ClosePoll)
        poll.POST("/:id/vote", controllers.VotePoll)
        poll.GET("/:id/results", controllers.ViewResults)
    }

    // User routes
    user := r.Group("/user")
    {
        user.GET("/:id", controllers.GetUserInfo)  // Get user info by ID
        user.PUT("/:id", controllers.UpdateUserInfo)  // Update user info
    }

    r.GET("/dashboard/:id", controllers.ViewDashboard)
    r.GET("/notifications/:id", controllers.GetNotifications)

    return r
}
