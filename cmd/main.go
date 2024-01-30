package main

import (
	"context"
	"log"

	"backend/internal/app"
)

// @title ShipStation RestAPI
// @version 1.0
// @description API server for ShipStation application

// @host http://localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	log.Println("Application start!")
	// Создаем контекст
	ctx := context.Background()

	// Создаем Aplication
	application, err := app.New(ctx)
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	// Запустите сервер, вызвав метод StartServer у объекта Application
	application.Run()
	log.Println("Application terminated!")
}
