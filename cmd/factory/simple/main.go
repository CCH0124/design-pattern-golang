package main

import (
	"designpattern/pkg/creational/factory/simple"
	"fmt"
)

func main() {
	factory := simple.NewSimplePizzaFactory()

	pizzaTypes := []string{"nycheese", "chicago", "unknown"}
	for _, t := range pizzaTypes {
		pizza, err := factory.CreatePizza(t)
		if err != nil {
			fmt.Printf("Failed to create pizza: %v\n", err)
			continue
		}
		// Fix: Print the actual type of the pizza structure instead of the string name.
		fmt.Printf("Created pizza type: %T (%s)\n", pizza, pizza.GetName())
		pizza.Prepare()
		pizza.Bake()
		pizza.Cut()
		pizza.Box()
		fmt.Println()
	}
}
