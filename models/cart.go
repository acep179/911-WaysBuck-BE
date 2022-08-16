package models

import "time"

type Cart struct {
	ID            int                `json:"id" gorm:"primary_key:auto_increment"`
	ProductID     int                `json:"prouduct_id"`
	Product       ProductCartRel     `json:"product"`
	TransactionID int                `json:"-"`
	Transaction   TransactionCartRel `json:"transaction"`
	Topping       []Topping          `json:"category" gorm:"many2many:cart_toppings"`
	CreatedAt     time.Time          `json:"-"`
	UpdatedAt     time.Time          `json:"-"`
}

type CartTransaction struct {
	ID int `json:"prouduct_id"`
}

func (CartTransaction) TableName() string {
	return "carts"
}
