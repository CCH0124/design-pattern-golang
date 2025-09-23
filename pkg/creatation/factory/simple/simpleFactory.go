package simple

import (
	"designpattern/pkg/creatation/factory/simple/base"
	"designpattern/pkg/creatation/factory/simple/concrete"
	"fmt"
)

// SimplePizzaFactory struct
type SimplePizzaFactory struct{}

// CreatePizza is the public method for creating pizzas
func (f *SimplePizzaFactory) CreatePizza(pizzaType string) (base.PizzaFlow, error) {
	switch pizzaType {
	case "nycheese":
		return concrete.NewNYStyleCheesePizza(), nil
	case "chicago":
		return concrete.NewChicagoStyleCheesePizza(), nil
	default:
		return nil, fmt.Errorf("Wrong pizza type passed")
	}
}
