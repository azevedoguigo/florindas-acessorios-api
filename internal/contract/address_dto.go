package contract

type NewAddress struct {
	CEP        string `json:"cep" validate:"required"`
	UF         string `json:"uf" validate:"required"`
	City       string `json:"city" validate:"required"`
	Street     string `json:"street" validate:"required"`
	Number     string `json:"number" validate:"required"`
	Complement string `json:"complement"`
	UserID     string `json:"user_id" validate:"required"`
}
