package main

import (
	"designpattern/pkg/creatation/factory/simple"
	"fmt"
)

func main() {
	factory := &simple.SimplePizzaFactory{}

	pizzaTypes := []string{"nycheese", "chicago", "unknown"}
	for _, t := range pizzaTypes {
		pizza, err := factory.CreatePizza(t)
		if err != nil {
			fmt.Printf("Failed to create pizza: %v\n", err)
			continue
		}
		fmt.Printf("Created pizza: %T\n", pizza.GetName())
		pizza.Prepare()
		pizza.Bake()
		pizza.Cut()
		pizza.Box()
	}
}
