package factory

import (
	"designpattern/pkg/creatation/factory/factory/base"
	"designpattern/pkg/creatation/factory/factory/ingredient"
	"fmt"
)

type VeggiePizza struct {
	base.Pizza
	PizzaIngredientFactory ingredient.PizzaIngredientFactory
}

func NewVeggiePizza(factory ingredient.PizzaIngredientFactory) *VeggiePizza {
	return &VeggiePizza{
		PizzaIngredientFactory: factory,
	}
}

func (p *VeggiePizza) Prepare() {
	fmt.Printf("Preparing %s\n", p.Name)
	p.Dough = p.PizzaIngredientFactory.CreateDough()
	p.Sauce = p.PizzaIngredientFactory.CreateSauce()
	p.Cheese = p.PizzaIngredientFactory.CreateCheese()
	p.Veggies = p.PizzaIngredientFactory.CreateVeggies()
}
