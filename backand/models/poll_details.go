package models

type PollDetails struct {
    Poll              Poll         `json:"poll"`
    Options           []PollOption `json:"options"`
    Results           []PollResult `json:"results,omitempty"`
    CanSeeResults     bool         `json:"can_see_results"`
    VotingURL    string `json:"voting_url"`
}