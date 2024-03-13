package api

import (
	"fmt"
	"log"

	stripe_sdk "github.com/stripe/stripe-go/v76"
	customer "github.com/stripe/stripe-go/v76/customer"
)

func CreateCustomer(email string, customerName string) {
	params := &stripe_sdk.CustomerParams{
		Email: stripe_sdk.String(email),
		Name:  stripe_sdk.String(customerName),
	}
	c, err := customer.New(params)
	if err != nil {
		log.Fatalf("customer.New: %v", err)
	}

	fmt.Printf("Customer created: %v\n", c.ID)
}

func ViewAllCustomers() {
	params := &stripe_sdk.CustomerListParams{}
	customers := customer.List(params)

	fmt.Printf("customers %v\n", customers.List())
}
