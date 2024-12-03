package contract

type NewProductDTO struct {
	Name        string `json:"name" validate:"required,min=2,max=60"`
	Description string `json:"description" validate:"required,min=2,max=400"`
	Price       string `json:"price" validate:"required"`
	Quantity    string `json:"quantity" validate:"required"`
	CategoryID  string `json:"category_id" validate:"required"`
}

type UpdateProductDTO struct {
	Images      []string `json:"images"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       *float64 `json:"price"`
	Quantity    *uint64  `json:"quantity"`
	CategoryID  string   `json:"category_id"`
}
