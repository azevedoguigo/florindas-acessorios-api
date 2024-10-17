package contract

type NewClientDTO struct {
	Name        string `json:"name" validate:"required,min=2,max=100"`
	Email       string `json:"email" validate:"required,email,max=100"`
	Password    string `json:"password" validate:"required,min=6,max=30"`
	CPF         string `json:"cpf" validate:"required"`
	UF          string `json:"uf" validate:"required"`
	CEP         string `json:"cep" validate:"required"`
	City        string `json:"city" validate:"required"`
	Address     string `json:"address" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}
