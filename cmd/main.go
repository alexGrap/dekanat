package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
	"ozon/internal/models"
	"ozon/internal/restDelivery"
)

func main() {
	err := godotenv.Load("cmd/.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	models.DB = os.Getenv("DATABASE")
	log.Print(fmt.Sprintf("Storage database - %s", models.DB))
	//viperConf, err := config.LoadConfig()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//conf, err := config.ParseConfig(viperConf)
	//if err != nil {
	//	log.Fatal(err)
	//}
	// if models.DB == "postgres" {
	//	postgresConnection(conf)
	//} else {
	//	log.Fatal("Error with db params")
	//}
	service := os.Getenv("SERVICE")
	log.Print(fmt.Sprintf("Sevice - %s", service))
	if service == "rest" {
		restService()
	} else {
		log.Fatal("Error with service params")
	}
}

func restService() {
	var app = fiber.New()
	restDelivery.Hearing(app)
}
