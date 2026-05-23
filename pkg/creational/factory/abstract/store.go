package abstract

import "fmt"

// PizzaCreator defines a function type for creating a Pizza.
type PizzaCreator func(pizzaType string) Pizza

// PizzaStore is a concrete type that uses composition rather than inheritance.
type PizzaStore struct {
	createPizza PizzaCreator
}

// NewPizzaStore is the constructor for PizzaStore.
func NewPizzaStore(creator PizzaCreator) *PizzaStore {
	return &PizzaStore{createPizza: creator}
}

// OrderPizza is the core workflow for ordering and making pizza.
func (s *PizzaStore) OrderPizza(pizzaType string) Pizza {
	pizza := s.createPizza(pizzaType)
	if pizza == nil {
		fmt.Printf("Error: Pizza type '%s' is not supported by this store.\n", pizzaType)
		return nil
	}
	fmt.Printf("--- Making a %s ---\n", pizza.GetName())
	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()
	return pizza
}

// --- Specific PizzaStore Creators ---

// NewNYPizzaStore creates a PizzaStore configured with New York style creation logic.
func NewNYPizzaStore() *PizzaStore {
	factory := &NYPizzaIngredientFactory{}
	return NewPizzaStore(func(pizzaType string) Pizza {
		switch pizzaType {
		case "cheese":
			return NewCheesePizza("New York Style Cheese Pizza", factory)
		case "pepperoni":
			return NewPepperoniPizza("New York Style Pepperoni Pizza", factory)
		case "veggie":
			return NewVeggiePizza("New York Style Veggie Pizza", factory)
		case "clam":
			return NewClamPizza("New York Style Clam Pizza", factory)
		default:
			return nil
		}
	})
}

// NewChicagoPizzaStore creates a PizzaStore configured with Chicago style creation logic.
func NewChicagoPizzaStore() *PizzaStore {
	factory := &ChicagoPizzaIngredientFactory{}
	return NewPizzaStore(func(pizzaType string) Pizza {
		switch pizzaType {
		case "cheese":
			return NewCheesePizza("Chicago Style Cheese Pizza", factory)
		case "pepperoni":
			return NewPepperoniPizza("Chicago Style Pepperoni Pizza", factory)
		case "veggie":
			return NewVeggiePizza("Chicago Style Veggie Pizza", factory)
		case "clam":
			return NewClamPizza("Chicago Style Clam Pizza", factory)
		default:
			return nil
		}
	})
}
