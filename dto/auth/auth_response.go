package authdto

type RegisterResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"fullName" form:"name" validate:"required"`
	Email    string `json:"email"  form:"email" validate:"required"`
	Password string `json:"password"  form:"password" validate:"required"`
}

type LoginResponse struct {
	FullName string `gorm:"type: varchar(255)" json:"fullName"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Password string `gorm:"type: varchar(255)" json:"password"`
	Token    string `gorm:"type: varchar(255)" json:"token"`
}
