package contract

type NewCartProductDTO struct {
	CartID    string `json:"cart_id" validate:"required"`
	ProductID string `json:"product_id" validate:"required"`
}

type UpdateCartProductQuantityDTO struct {
	CartProductID string `json:"cart_product_id" validate:"required"`
	Quantity      uint64 `json:"quantity" validate:"required"`
}
