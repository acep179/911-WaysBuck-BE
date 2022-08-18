package authdto

type RegisterRequest struct {
	FullName string `json:"fullName" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"loginEmail" form:"email" validate:"required"`
	Password string `json:"loginPassword" form:"password" validate:"required"`
}
