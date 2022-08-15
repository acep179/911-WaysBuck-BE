package models

import "time"

type Cart struct {
	ID        int       `json:"id"`
	ProductID int       `json:"prouduct_id"`
	ToppingID int       `json:"topping_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
