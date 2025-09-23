package factory

import (
	"designpattern/pkg/creatation/factory/factory/base"
	"designpattern/pkg/creatation/factory/factory/ingredient"
	"fmt"
)

type CheesePizza struct {
	base.Pizza
	PizzaIngredientFactory ingredient.PizzaIngredientFactory
}

func NewCheesePizza(factory ingredient.PizzaIngredientFactory) *CheesePizza {
	return &CheesePizza{
		PizzaIngredientFactory: factory,
	}
}

func (p *CheesePizza) Prepare() {
	fmt.Printf("Preparing %s\n", p.Name)
	p.Dough = p.PizzaIngredientFactory.CreateDough()
	p.Sauce = p.PizzaIngredientFactory.CreateSauce()
	p.Cheese = p.PizzaIngredientFactory.CreateCheese()
}
