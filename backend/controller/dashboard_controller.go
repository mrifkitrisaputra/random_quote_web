package controllers

import (
    "net/http"
    "backend/config"
    "github.com/gin-gonic/gin"
)

func ViewDashboard(c *gin.Context) {
    userID := c.Param("id")
    rows, err := config.DB.Query("SELECT id, title, status FROM polls WHERE creator_id = ?", userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()

    var polls []map[string]interface{}
    for rows.Next() {
        var id, title, status string
        rows.Scan(&id, &title, &status)
        polls = append(polls, gin.H{"id": id, "title": title, "status": status})
    }
    c.JSON(http.StatusOK, gin.H{"dashboard": polls})
}