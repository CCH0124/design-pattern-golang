package simple

import "fmt"

// SimplePizzaFactory creates instances of Pizza.
type SimplePizzaFactory struct{}

// NewSimplePizzaFactory is the constructor for SimplePizzaFactory.
func NewSimplePizzaFactory() *SimplePizzaFactory {
	return &SimplePizzaFactory{}
}

// CreatePizza creates a Pizza based on type name.
func (f *SimplePizzaFactory) CreatePizza(pizzaType string) (Pizza, error) {
	switch pizzaType {
	case "nycheese":
		return NewNYStyleCheesePizza(), nil
	case "chicago":
		return NewChicagoStyleCheesePizza(), nil
	default:
		return nil, fmt.Errorf("wrong pizza type: %s", pizzaType)
	}
}
