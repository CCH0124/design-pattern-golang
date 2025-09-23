package factory

import (
	"designpattern/pkg/creatation/factory/factory/base"
	"designpattern/pkg/creatation/factory/factory/ingredient"
	"fmt"
)

type ClamPizza struct {
	base.Pizza
	PizzaIngredientFactory ingredient.PizzaIngredientFactory
}

func NewClamPizza(factory ingredient.PizzaIngredientFactory) *ClamPizza {
	return &ClamPizza{
		PizzaIngredientFactory: factory,
	}
}

func (p *ClamPizza) Prepare() {
	fmt.Printf("Preparing %s\n", p.Name)
	p.Dough = p.PizzaIngredientFactory.CreateDough()
	p.Sauce = p.PizzaIngredientFactory.CreateSauce()
	p.Cheese = p.PizzaIngredientFactory.CreateCheese()
	p.Clams = p.PizzaIngredientFactory.CreateClams()
}
