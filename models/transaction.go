package models

import "time"

type Transaction struct {
	ID        int            `json:"id" gorm:"primary_key:auto_increment"`
	Amount    int            `json:"amount"`
	UserID    int            `json:"user_id"`
	User      UserProfileRel `json:"user"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
}

type TransactionCartRel struct {
	ID     int `json:"id"`
	UserId int `json:"user_id"`
}
type TransactionUserRel struct {
	ID     int `json:"id"`
	UserId int `json:"user_id"`
}

func (TransactionCartRel) TableName() string {
	return "transactions"
}
func (TransactionUserRel) TableName() string {
	return "transactions"
}