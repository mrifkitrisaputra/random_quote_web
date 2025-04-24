package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "backend/config"
)

func ViewResults(c *gin.Context) {
    pollID := c.Param("id")
    rows, err := config.DB.Query(`
        SELECT o.id, o.text, COUNT(v.id) as votes
        FROM options o
        LEFT JOIN votes v ON o.id = v.option_id
        WHERE o.poll_id = ?
        GROUP BY o.id
    `, pollID)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()

    var results []map[string]interface{}
    for rows.Next() {
        var id, text string
        var votes int
        rows.Scan(&id, &text, &votes)
        results = append(results, gin.H{"optionId": id, "text": text, "votes": votes})
    }

    c.JSON(http.StatusOK, gin.H{"results": results})
}
