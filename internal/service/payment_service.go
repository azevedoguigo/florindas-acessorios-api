package service

import (
	"context"

	"github.com/azevedoguigo/florindas-acessorios-api/internal/contract"
	"github.com/azevedoguigo/florindas-acessorios-api/pkg"
	"github.com/mercadopago/sdk-go/pkg/preference"
)

type PaymentService interface {
	Pay(paymentDTO contract.CreatePreferenceDTO) (*preference.Response, error)
}

type paymentService struct {
	client preference.Client
}

func NewPaymentService(client preference.Client) PaymentService {
	return paymentService{client: client}
}

func (p paymentService) Pay(createPreferenceDTO contract.CreatePreferenceDTO) (*preference.Response, error) {
	if err := pkg.ValidateStruct(createPreferenceDTO); err != nil {
		return &preference.Response{}, err
	}

	var items []preference.ItemRequest
	for _, dtoItem := range createPreferenceDTO.Items {
		items = append(items, preference.ItemRequest{
			Title:     dtoItem.Title,
			Quantity:  dtoItem.Quantity,
			UnitPrice: dtoItem.Price,
		})
	}

	request := preference.Request{
		Items: items,
	}

	resource, err := p.client.Create(context.Background(), request)
	if err != nil {
		return nil, err
	}

	return resource, nil
}
