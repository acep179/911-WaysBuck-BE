package models

type Topping struct {
	ID     int    `json:"id"`
	Title  string `json:"title" gorm:"type: varchar(255)"`
	Price  int    `json:"price" `
	Image  string `json:"image" gorm:"type: varchar(255)"`
	UserID int    `json:"user_id"`
}
