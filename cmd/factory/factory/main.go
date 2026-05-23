package main

import (
	"designpattern/pkg/creational/factory/abstract"
	"fmt"
)

func main() {
	// NY Pizza Store
	nyStore := abstract.NewNYPizzaStore()
	fmt.Println("=== Ordering Pizza from New York Pizza Store ===")
	nyPizza := nyStore.OrderPizza("cheese")
	if nyPizza != nil {
		fmt.Println(nyPizza.String())
	}

	// Chicago Pizza Store
	chicagoStore := abstract.NewChicagoPizzaStore()
	fmt.Println("=== Ordering Pizza from Chicago Pizza Store ===")
	chicagoPizza := chicagoStore.OrderPizza("veggie")
	if chicagoPizza != nil {
		fmt.Println(chicagoPizza.String())
	}
}
