package base

import (
	"fmt"
)

type PizzaStore interface {
	CreatePizza(item string) IPizza
	OrderPizza(typeName string) IPizza
}

type AbstractPizzaStore struct {
	PizzaStore
}

func (s *AbstractPizzaStore) OrderPizza(typeName string) IPizza {
	pizza := s.CreatePizza(typeName)
	fmt.Printf("--- Making a %s ---\n", pizza.GetName())
	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()
	return pizza
}
