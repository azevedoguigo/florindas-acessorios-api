package config

import (
	"log"
	"os"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/payment"
)

func InitMercadoPago() payment.Client {
	accessToken := os.Getenv("MERCADO_PAGO_ACCESS_TOKEN")

	cfg, err := config.New(accessToken)
	if err != nil {
		log.Fatalln(err.Error())
		return nil
	}

	client := payment.NewClient(cfg)
	return client
}
