package authdto

type RegisterResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"fullName" form:"name" validate:"required"`
	Email    string `json:"email"  form:"email" validate:"required"`
	Password string `json:"password"  form:"password" validate:"required"`
}
