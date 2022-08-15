package models

import "time"

type Product struct {
	ID        int            `json:"id" gorm:"primary_key:auto_increment"`
	Title     string         `json:"title" gorm:"type: varchar(255)"`
	Price     int            `json:"price" gorm:"type: varchar(255)"`
	Image     string         `json:"image" gorm:"type: varchar(255)"`
	UserID    int            `json:"-"`
	User      UserProfileRel `json:"user"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
}

type ProductUserRel struct {
	ID     int    `json:"id" gorm:"primary_key:auto_increment"`
	Title  string `json:"title" gorm:"type: varchar(255)"`
	Price  int    `json:"price" gorm:"type: varchar(255)"`
	Image  string `json:"image" gorm:"type: varchar(255)"`
	UserID int    `json:"-"`
}

func (ProductUserRel) TableName() string {
	return "products"
}
