package main

import (
	payment_service "github.com/yerassylbolatov/payment-service"
	"github.com/yerassylbolatov/payment-service/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(payment_service.Server)
	err := srv.Run("8080", handlers.InitRoutes())
	if err != nil {
		log.Fatalf("error occured: %s", err.Error())
	}
}
