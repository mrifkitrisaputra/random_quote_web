package controllers

import (
    "net/http"
    "time"
    "github.com/gin-gonic/gin"
    "backend/models"
    "backend/config"
    "github.com/google/uuid"
)

func VotePoll(c *gin.Context) {
    var vote models.Vote
    if err := c.ShouldBindJSON(&vote); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    vote.ID = uuid.NewString()
    vote.Timestamp = time.Now()
    vote.PollID = c.Param("id")

    _, err := config.DB.Exec("INSERT INTO votes (id, poll_id, option_id, user_identifier, fingerprint, timestamp) VALUES (?, ?, ?, ?, ?, ?)",
        vote.ID, vote.PollID, vote.OptionID, vote.UserIdentifier, vote.Fingerprint, vote.Timestamp)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Vote submitted"})
}
