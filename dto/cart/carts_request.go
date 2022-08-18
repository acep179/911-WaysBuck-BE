package cartsdto

type CartRequest struct {
	ProductID     string `json:"product_id" form:"product_id" validate:"required"`
	TransactionID int    `json:"transaction_id" form:"transaction_id" validate:"required"`
	Topping       string `json:"topping" form:"topping" validate:"required"`
}
