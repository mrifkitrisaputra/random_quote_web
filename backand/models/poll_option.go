package models

type PollOption struct {
    ID    int    `json:"id"`
    Text  string `json:"text"`
    Order int    `json:"order_index"`
}