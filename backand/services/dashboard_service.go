package services

import (
    "database/sql"
    "fmt"
    "time"

    "backand/config"
)

type PollSummary struct {
    ID                   string `json:"id"`
    Title                string `json:"title"`
    Description          string `json:"description"`
    Status               string `json:"status"`         // open / closed
    VoteCount            int    `json:"vote_count"`     // jumlah voter
    OptionCount          int    `json:"option_count"`   // jumlah opsi jawaban
    ExpiresAt            string `json:"expires_at"`     // format user-friendly
    AllowRealtimeResults bool   `json:"allow_realtime"` // apakah peserta bisa lihat hasil saat aktif
}

func GetPollsForDashboard() ([]PollSummary, error) {
    rows, err := config.DB.Query(`
        SELECT 
            p.id, 
            p.title, 
            p.description, 
            p.status,
            p.expires_at,
            p.allow_realtime_results
        FROM polls p
    `)
    if err != nil {
        return nil, fmt.Errorf("gagal mengambil daftar polling: %v", err)
    }
    defer rows.Close()

    var polls []PollSummary
    for rows.Next() {
        var (
            id             string
            title          string
            desc           sql.NullString
            status         string
            expiresAt      string
            allowRealtime  bool
        )

        err := rows.Scan(&id, &title, &desc, &status, &expiresAt, &allowRealtime)
        if err != nil {
            return nil, err
        }

        // Hitung vote_count per polling
        var voteCount int
        err = config.DB.QueryRow("SELECT COUNT(*) FROM votes WHERE poll_id = ?", id).Scan(&voteCount)
        if err != nil {
            return nil, fmt.Errorf("gagal menghitung jumlah suara untuk polling %s: %v", id, err)
        }

        // Hitung option_count per polling
        var optionCount int
        err = config.DB.QueryRow("SELECT COUNT(*) FROM poll_options WHERE poll_id = ?", id).Scan(&optionCount)
        if err != nil {
            return nil, fmt.Errorf("gagal menghitung jumlah opsi untuk polling %s: %v", id, err)
        }

        // Format waktu
        formattedTime := parseTime(expiresAt)

        polls = append(polls, PollSummary{
            ID:                 id,
            Title:              title,
            Description:        desc.String,
            Status:             status,
            VoteCount:          voteCount,
            OptionCount:        optionCount,
            ExpiresAt:          formattedTime,
            AllowRealtimeResults: allowRealtime,
        })
    }

    return polls, nil
}

func GetPollDetails(pollID string) (*PollSummary, error) {
    row := config.DB.QueryRow(`
        SELECT 
            id, 
            title, 
            description, 
            status, 
            expires_at,
            allow_realtime_results
        FROM polls WHERE id = ?`, pollID)

    var (
        id             string
        title          string
        desc           sql.NullString
        status         string
        expiresAt      string
        allowRealtime  bool
    )

    if err := row.Scan(&id, &title, &desc, &status, &expiresAt, &allowRealtime); err != nil {
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf("poll not found")
        }
        return nil, fmt.Errorf("gagal mendapatkan detail polling: %v", err)
    }

    // Hitung jumlah opsi jawaban
    var optionCount int
    err := config.DB.QueryRow("SELECT COUNT(*) FROM poll_options WHERE poll_id = ?", pollID).Scan(&optionCount)
    if err != nil {
        return nil, fmt.Errorf("gagal menghitung jumlah opsi: %v", err)
    }

    // Hitung jumlah suara
    var voteCount int
    err = config.DB.QueryRow("SELECT COUNT(*) FROM votes WHERE poll_id = ?", pollID).Scan(&voteCount)
    if err != nil {
        return nil, fmt.Errorf("gagal menghitung jumlah suara: %v", err)
    }

    // Format waktu
    formattedTime := parseTime(expiresAt)

    poll := &PollSummary{
        ID:                 id,
        Title:              title,
        Description:        desc.String,
        Status:             status,
        ExpiresAt:          formattedTime,
        AllowRealtimeResults: allowRealtime,
        OptionCount:        optionCount,
        VoteCount:          voteCount,
    }

    return poll, nil
}

// Fungsi bantuan parsing waktu
func parseTime(timeStr string) string {
    t, err := time.Parse("2006-01-02 15:04:05", timeStr)
    if err != nil {
        // Coba parsing ulang dengan format yang berbeda (misalnya ISO8601)
        t2, err2 := time.Parse(time.RFC3339, timeStr)
        if err2 != nil {
            return timeStr
        }
        return t2.Format("02 Januari 2006, 15:04 WIB")
    }
    return t.Format("02 Januari 2006, 15:04 WIB")
}