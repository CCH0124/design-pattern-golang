package factory

import (
	"designpattern/pkg/creatation/factory/factory/base"
	"designpattern/pkg/creatation/factory/factory/ingredient"
	"fmt"
)

type PepperoniPizza struct {
	base.Pizza
	PizzaIngredientFactory ingredient.PizzaIngredientFactory
}

func NewPepperoniPizza(factory ingredient.PizzaIngredientFactory) *PepperoniPizza {
	return &PepperoniPizza{
		PizzaIngredientFactory: factory,
	}
}

func (p *PepperoniPizza) Prepare() {
	fmt.Printf("Preparing %s\n", p.Name)
	p.Dough = p.PizzaIngredientFactory.CreateDough()
	p.Sauce = p.PizzaIngredientFactory.CreateSauce()
	p.Cheese = p.PizzaIngredientFactory.CreateCheese()
	p.Veggies = p.PizzaIngredientFactory.CreateVeggies()
	p.Pepperoni = p.PizzaIngredientFactory.CreatePepperoni()
}
