package base

import (
	"designpattern/pkg/creatation/factory/factory/ingredient"
	"fmt"
	"strings"
)

type Pizza struct {
	Name      string
	Dough     ingredient.Dough
	Sauce     ingredient.Sauce
	Veggies   []ingredient.Veggies
	Cheese    ingredient.Cheese
	Pepperoni ingredient.Pepperoni
	Clams     ingredient.Clams
}

type IPizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
	String() string
	GetName() string
}

func (p *Pizza) Bake() {
	fmt.Printf("%s Bake for 25 minutes at 350\n", p.Name)
}

func (p *Pizza) Cut() {
	fmt.Printf("%s Cutting the pizza into diagonal slices\n", p.Name)
}

func (p *Pizza) Box() {
	fmt.Printf("%s Place pizza in official PizzaStore box\n", p.Name)
}

func (p *Pizza) GetName() string {
	return p.Name
}

func (p *Pizza) String() string {
	var result strings.Builder
	result.WriteString(fmt.Sprintf("---- %s ----\n", p.Name))
	if p.Dough != nil {
		result.WriteString(p.Dough.Something() + "\n")
	}
	if p.Sauce != nil {
		result.WriteString(p.Sauce.Something() + "\n")
	}
	if p.Cheese != nil {
		result.WriteString(p.Cheese.Something() + "\n")
	}
	if p.Veggies != nil {
		for i, v := range p.Veggies {
			result.WriteString(v.Something())
			if i < len(p.Veggies)-1 {
				result.WriteString(", ")
			}
		}
		result.WriteString("\n")
	}
	if p.Clams != nil {
		result.WriteString(p.Clams.Something() + "\n")
	}
	if p.Pepperoni != nil {
		result.WriteString(p.Pepperoni.Something() + "\n")
	}
	return result.String()
}
