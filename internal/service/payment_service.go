package service

import (
	"context"

	"github.com/azevedoguigo/florindas-acessorios-api/internal/contract"
	"github.com/azevedoguigo/florindas-acessorios-api/pkg"
	"github.com/mercadopago/sdk-go/pkg/payment"
)

type PaymentService interface {
	Pay(paymentDTO contract.PaymentDTO) (payment.Response, error)
}

type paymentService struct {
	client payment.Client
}

func NewPaymentService(client payment.Client) PaymentService {
	return paymentService{client: client}
}

func (p paymentService) Pay(paymentDTO contract.PaymentDTO) (payment.Response, error) {
	if err := pkg.ValidateStruct(paymentDTO); err != nil {
		return payment.Response{}, err
	}

	request := payment.Request{
		TransactionAmount: paymentDTO.TransactionAmount,
		PaymentMethodID:   paymentDTO.PaymentMethodId,
		Payer: &payment.PayerRequest{
			Email: paymentDTO.Payer.Email,
		},
		Token:        paymentDTO.Token,
		Installments: int(paymentDTO.Installments),
	}

	resource, err := p.client.Create(context.Background(), request)
	if err != nil {
		return payment.Response{}, err
	}

	return *resource, nil
}
