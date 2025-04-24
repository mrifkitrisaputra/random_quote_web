package controllers

import (
    "net/http"
    "backend/config"
    "github.com/gin-gonic/gin"
    "backend/models"
)

// GetUserInfo handles fetching the user info
func GetUserInfo(c *gin.Context) {
    userID := c.Param("id")
    var user models.User
    row := config.DB.QueryRow("SELECT id, email, username, fingerprint FROM users WHERE id = ?", userID)
    
    err := row.Scan(&user.ID, &user.Email, &user.Username, &user.Fingerprint)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    
    c.JSON(http.StatusOK, user)
}

// UpdateUserInfo handles updating the user information
func UpdateUserInfo(c *gin.Context) {
    userID := c.Param("id")
    var user models.User
    
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    _, err := config.DB.Exec("UPDATE users SET email = ?, username = ?, fingerprint = ? WHERE id = ?",
        user.Email, user.Username, user.Fingerprint, userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{"message": "User info updated"})
}

