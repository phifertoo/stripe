package auth

import (
	"fmt"
	"log"

	stripe "github.com/stripe/stripe-go/v76"
	customer "github.com/stripe/stripe-go/v76/customer"
)

func GetStripe() {
	params := &stripe.CustomerParams{
		Email: stripe.String("customer@example.com"),
		Name:  stripe.String("Customer Name"),
	}
	c, err := customer.New(params)
	if err != nil {
		log.Fatalf("customer.New: %v", err)
	}

	fmt.Printf("Customer created: %v\n", c.ID)
	// url := "https://api.stripe.com" // Replace with your desired API endpoint
	// resp, err := http.Get(url)
	// if err != nil {
	// 	fmt.Printf("Error making request: %s\n", err)
	// 	return
	// }
	// defer resp.Body.Close()

	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Printf("Error reading body: %s\n", err)
	// 	return
	// }

	// fmt.Printf("Response Body: %s\n", string(body))
}
