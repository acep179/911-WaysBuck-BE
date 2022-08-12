package productsdto

type ProductResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title" form:"title" validate:"required"`
	Price string `json:"price"  form:"price" validate:"required"`
	Image string `json:"image"  form:"image" validate:"required"`
}
