package controllers

import (
    "net/http"
    "time"
    "github.com/gin-gonic/gin"
    "backend/models"
    "backend/config"
    "github.com/google/uuid"
)

func CreatePoll(c *gin.Context) {
    var poll models.Poll
    if err := c.ShouldBindJSON(&poll); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    poll.ID = uuid.NewString()
    poll.CreatedAt = time.Now()
    poll.Status = models.Open

    _, err := config.DB.Exec("INSERT INTO polls (id, title, description, creator_id, created_at, expires_at, status) VALUES (?, ?, ?, ?, ?, ?, ?)",
        poll.ID, poll.Title, poll.Description, poll.CreatorID, poll.CreatedAt, poll.ExpiresAt, poll.Status)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, poll)
}

func ViewPoll(c *gin.Context) {
    id := c.Param("id")
    var poll models.Poll
    row := config.DB.QueryRow("SELECT id, title, description, creator_id, created_at, expires_at, status FROM polls WHERE id = ?", id)
    err := row.Scan(&poll.ID, &poll.Title, &poll.Description, &poll.CreatorID, &poll.CreatedAt, &poll.ExpiresAt, &poll.Status)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Poll not found"})
        return
    }
    c.JSON(http.StatusOK, poll)
}

func ClosePoll(c *gin.Context) {
    id := c.Param("id")
    _, err := config.DB.Exec("UPDATE polls SET status = 'closed' WHERE id = ?", id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Poll closed"})
}
