package contract

type PaymentDTO struct {
	Token             string  `json:"token"`
	IssuerId          string  `json:"issuer_id"`
	PaymentMethodId   string  `json:"payment_method_id"`
	TransactionAmount float64 `json:"transaction_amount"`
	Installments      uint64  `json:"installments"`
	Payer             struct {
		Email          string `json:"email"`
		Identification struct {
			Type   string `json:"type"`
			Number string `json:"number"`
		}
	}
}
