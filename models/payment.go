package models

type Payment struct {
    Username string `json:"username"`
    Amount   float64 `json:"amount"`
}
