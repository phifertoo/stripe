package main

import (
	"fmt"
	"os"

	stripe_api "github.com/stripe/stripe-go/v76"
)

func main() {
	stripe_api.Key = os.Getenv("STRIPE_TEST_KEY")

	fmt.Print("test")
}
