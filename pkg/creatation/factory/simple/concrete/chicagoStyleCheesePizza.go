package concrete

import (
	"designpattern/pkg/creatation/factory/simple/base"
	"fmt"
)

type ChicagoStyleCheesePizza struct {
	*Pizza
}

func (n *ChicagoStyleCheesePizza) Prepare() {
	n.Pizza.Prepare()
	fmt.Println("Extra step for Chicago Style Deep Dish Cheese Pizza!")
}

func NewChicagoStyleCheesePizza() base.PizzaFlow {
	p := &Pizza{
		name:     "Chicago Style Deep Dish Cheese Pizza",
		dough:    "Extra Thick Crust Dough",
		sauce:    "Plum Tomato Sauce",
		toppings: []string{"Shredded Mozzarella Cheese"},
	}

	return &ChicagoStyleCheesePizza{
		Pizza: p,
	}
}
