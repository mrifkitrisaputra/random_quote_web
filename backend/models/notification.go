package models

type NotificationType string

const (
    Info    NotificationType = "info"
    Warning NotificationType = "warning"
    Alert   NotificationType = "alert"
)

type Notification struct {
    ID      string            `json:"id"`
    UserID  string            `json:"userId"`
    PollID  string            `json:"pollId"`
    Message string            `json:"message"`
    Type    NotificationType  `json:"type"`
    IsRead  bool              `json:"isRead"`
}
