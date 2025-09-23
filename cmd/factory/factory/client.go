package main

import (
	"designpattern/pkg/creatation/factory/factory"
	"fmt"
)

func main() {
	store := factory.NewChicagoPizzaStore()
	pizza := store.OrderPizza("cheese")
	fmt.Println(pizza.String())
	pizza = store.OrderPizza("veggie")
	fmt.Println(pizza.String())
}
