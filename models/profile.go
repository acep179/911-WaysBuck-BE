package models

import "time"

type Profile struct {
	ID        int       `json:"id"`
	Image     string    `json:"image" gorm:"type: varchar(255)"`
	Address   string    `json:"address" gorm:"type: varchar(255)"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
