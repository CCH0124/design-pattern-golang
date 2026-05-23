package simple

import "fmt"

// Pizza defines the interface that all pizzas must implement.
type Pizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
	GetName() string
}

// pizza represents the base fields and default implementations for a Pizza.
// By making it lowercase, we keep it package-private, avoiding external abuse,
// while exposing the Pizza interface.
type pizza struct {
	name     string
	dough    string
	sauce    string
	toppings []string
}

func (p *pizza) Prepare() {
	fmt.Printf("Preparing %s Pizza\n", p.name)
	fmt.Printf("Tossing %s dough\n", p.dough)
	fmt.Printf("Adding %s sauce\n", p.sauce)
	fmt.Print("Adding toppings: \n")
	for _, t := range p.toppings {
		fmt.Printf("\t%s, ", t)
	}
	fmt.Println()
}

func (p *pizza) Bake() {
	fmt.Println(p.name + " Baking Pizza for 25 min at 350")
}

func (p *pizza) Cut() {
	fmt.Println(p.name + " Cutting the pizza into diagonal slices")
}

func (p *pizza) Box() {
	fmt.Println(p.name + " Place pizza in official PizzaStore box")
}

func (p *pizza) GetName() string {
	return p.name
}

// NYStyleCheesePizza is a specific NY style pizza.
type NYStyleCheesePizza struct {
	*pizza
}

func (n *NYStyleCheesePizza) Prepare() {
	n.pizza.Prepare()
	fmt.Println("Extra step for NY Style Cheese Pizza!")
}

// NewNYStyleCheesePizza constructor
func NewNYStyleCheesePizza() Pizza {
	return &NYStyleCheesePizza{
		pizza: &pizza{
			name:     "NY Style Sauce and Cheese Pizza",
			dough:    "Thin Crust Dough",
			sauce:    "Marinara Sauce",
			toppings: []string{"Grated Reggiano Cheese"},
		},
	}
}

// ChicagoStyleCheesePizza is a specific Chicago style pizza.
type ChicagoStyleCheesePizza struct {
	*pizza
}

func (c *ChicagoStyleCheesePizza) Prepare() {
	c.pizza.Prepare()
	fmt.Println("Extra step for Chicago Style Deep Dish Cheese Pizza!")
}

// NewChicagoStyleCheesePizza constructor
func NewChicagoStyleCheesePizza() Pizza {
	return &ChicagoStyleCheesePizza{
		pizza: &pizza{
			name:     "Chicago Style Deep Dish Cheese Pizza",
			dough:    "Extra Thick Crust Dough",
			sauce:    "Plum Tomato Sauce",
			toppings: []string{"Shredded Mozzarella Cheese"},
		},
	}
}
