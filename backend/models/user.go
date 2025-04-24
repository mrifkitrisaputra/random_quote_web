package models

type User struct {
    ID         string `json:"id"`
    Email      string `json:"email"`
    Username   string `json:"username"`
    Password   string `json:"password"`
    Fingerprint string `json:"fingerprint"`
    Role       string `json:"role"`
}