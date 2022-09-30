package payment_service

import "time"

type Transactions struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	Email     string    `json:"email"`
	Amount    int       `json:"amount"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Status    string    `json:"status"`
}
