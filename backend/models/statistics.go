package models

type Statistics struct {
    PollID               string         `json:"pollId"`
    TotalParticipants    int            `json:"totalParticipants"`
    ParticipationOverTime map[string]int `json:"participationOverTime"`
}
