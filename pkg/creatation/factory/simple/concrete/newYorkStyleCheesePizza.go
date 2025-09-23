package concrete

import (
	"designpattern/pkg/creatation/factory/simple/base"
	"fmt"
)

type NewYorkStyleCheesePizza struct {
	*Pizza
}

func (n *NewYorkStyleCheesePizza) Prepare() {
	n.Pizza.Prepare()
	fmt.Println("Extra step for NY Style Cheese Pizza!")
}

func NewNYStyleCheesePizza() base.PizzaFlow {
	p := &Pizza{
		name:     "NY Style Sauce and Cheese Pizza",
		dough:    "Thin Crust Dough",
		sauce:    "Marinara Sauce",
		toppings: []string{"Grated Reggiano Cheese"},
	}

	return &NewYorkStyleCheesePizza{
		Pizza: p,
	}
}
