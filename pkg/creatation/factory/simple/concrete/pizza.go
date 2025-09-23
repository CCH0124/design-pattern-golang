package concrete

import (
	"fmt"
)

type Pizza struct {
	name     string
	dough    string
	sauce    string
	toppings []string
}

// Implementing the default preparation steps for pizza
func (p *Pizza) Prepare() {
	fmt.Printf("Preparing %s Pizza\n", p.name)
	fmt.Printf("Tossing %s dough\n", p.dough)
	fmt.Printf("Adding %s sauce\n", p.sauce)
	fmt.Print("Adding toppings: \n")
	for _, t := range p.toppings {
		fmt.Printf("\t%s, ", t)
	}
	fmt.Println()
}

func (p *Pizza) Bake() {
	fmt.Println(p.name + " Baking Pizza for 25 min at 350")
}

func (p *Pizza) Cut() {
	fmt.Println(p.name + " Cutting the pizza into diagonal slices")
}

func (p *Pizza) Box() {
	fmt.Println(p.name + " Place pizza in official PizzaStore box")
}

func (p *Pizza) GetName() string {
	return p.name
}
