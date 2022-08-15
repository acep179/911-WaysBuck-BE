package models

import "time"

type Transaction struct {
	ID        int       `json:"id"`
	BuyerID   int       `json:"buyer_id"`
	CartID    int       `json:"cart_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
