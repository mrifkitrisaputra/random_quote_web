package service

import (
    "backend/config"
    "backend/models"
    "fmt"
    "time"

    "github.com/google/uuid"
)

type CreatePollService struct{}

func NewCreatePollService() *CreatePollService {
    return &CreatePollService{}
}

func (s *CreatePollService) CreatePoll(title, description string, options []string) (*models.PollDetails, error) {
    if len(title) == 0 || len(options) < 2 {
        return nil, fmt.Errorf("judul dan minimal 2 opsi wajib diisi")
    }

    pollID := uuid.New().String()
    userID := uuid.New().String()
    isOpen := true
    createdAt := time.Now()

    optionsJSON, err := models.SerializeOptions(options)
    if err != nil {
        return nil, fmt.Errorf("gagal memproses opsi")
    }

    query := `
        INSERT INTO polls (id, title, description, options, is_open, created_at, user_id)
        VALUES (?, ?, ?, ?, ?, ?, ?)
    `

    _, err = config.DB.Exec(query, pollID, title, description, optionsJSON, isOpen, createdAt, userID)
    if err != nil {
        return nil, fmt.Errorf("gagal menyimpan poll: %v", err)
    }

    return &models.PollDetails{
        PollID:  pollID,
        PollURL: fmt.Sprintf("https://quickpoll/polls/%s", pollID),
    }, nil
}