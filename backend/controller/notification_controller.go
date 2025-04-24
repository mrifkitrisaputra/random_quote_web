package controllers

import (
    "net/http"
    "backend/config"
    "github.com/gin-gonic/gin"
)

func GetNotifications(c *gin.Context) {
    userID := c.Param("id")
    rows, err := config.DB.Query("SELECT id, message, type, is_read FROM notifications WHERE user_id = ?", userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()

    var notes []map[string]interface{}
    for rows.Next() {
        var id, message, ntype string
        var isRead bool
        rows.Scan(&id, &message, &ntype, &isRead)
        notes = append(notes, gin.H{"id": id, "message": message, "type": ntype, "isRead": isRead})
    }
    c.JSON(http.StatusOK, notes)
}