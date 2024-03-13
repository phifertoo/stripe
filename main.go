package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/phifertoo/stripe/api"
	stripe_api "github.com/stripe/stripe-go/v76"
)

func main() {
	err := godotenv.Load() // This will load the .env file
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	stripe_api.Key = os.Getenv("STRIPE_SECRET_KEY")
	// api.CreateCustomer()
	api.ViewAllCustomers()

	fmt.Print("test")
}
