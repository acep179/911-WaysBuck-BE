package cartsdto

type CartRequest struct {
	ProductID     int `json:"product_id" form:"product_id" validate:"required"`
	TransactionID int `json:"transaction_id" form:"transaction_id"`
}
