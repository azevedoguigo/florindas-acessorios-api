package config

import (
	"log"
	"os"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/preference"
)

func InitMercadoPago() preference.Client {
	accessToken := os.Getenv("MERCADO_PAGO_ACCESS_TOKEN")

	cfg, err := config.New(accessToken)
	if err != nil {
		log.Fatalln(err.Error())
		return nil
	}

	client := preference.NewClient(cfg)
	return client
}
