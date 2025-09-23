package factory

import (
	"designpattern/pkg/creatation/factory/factory/base"
	"designpattern/pkg/creatation/factory/factory/concrete"
)

type ChicagoPizzaStore struct {
	base.AbstractPizzaStore
}

func NewChicagoPizzaStore() *ChicagoPizzaStore {
	store := &ChicagoPizzaStore{}
	store.AbstractPizzaStore.PizzaStore = store
	return store
}

func (s *ChicagoPizzaStore) CreatePizza(pizzaType string) base.IPizza {
	var pizza base.IPizza
	factory := &concrete.ChicagoPizzaIngredientFactory{}
	switch pizzaType {
	case "cheese":
		pizza = NewCheesePizza(factory)
		if cp, ok := pizza.(*CheesePizza); ok {
			cp.Name = "Chicago Style Cheese Pizza"
		}
	case "pepperoni":
		pizza = NewPepperoniPizza(factory)
		if cp, ok := pizza.(*PepperoniPizza); ok {
			cp.Name = "Chicago Style Pepperoni Pizza"
		}
	case "veggie":
		pizza = NewVeggiePizza(factory)
		if cp, ok := pizza.(*VeggiePizza); ok {
			cp.Name = "Chicago Style Veggie Pizza"
		}
	case "clam":
		pizza = NewClamPizza(factory)
		if cp, ok := pizza.(*ClamPizza); ok {
			cp.Name = "Chicago Style Clam Pizza"
		}
	}
	return pizza
}
