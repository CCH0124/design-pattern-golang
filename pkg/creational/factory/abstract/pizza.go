package abstract

import (
	"fmt"
	"strings"
)

// Pizza defines the interface for pizzas in the abstract factory system.
type Pizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
	GetName() string
	String() string
}

// pizza is the base implementation for Pizza.
type pizza struct {
	name      string
	dough     Dough
	sauce     Sauce
	cheese    Cheese
	veggies   []Veggies
	pepperoni Pepperoni
	clams     Clams
}

func (p *pizza) Bake() {
	fmt.Printf("%s: Baking for 25 minutes at 350\n", p.name)
}

func (p *pizza) Cut() {
	fmt.Printf("%s: Cutting the pizza into diagonal slices\n", p.name)
}

func (p *pizza) Box() {
	fmt.Printf("%s: Place pizza in official PizzaStore box\n", p.name)
}

func (p *pizza) GetName() string {
	return p.name
}

func (p *pizza) String() string {
	var result strings.Builder
	result.WriteString(fmt.Sprintf("---- %s ----\n", p.name))
	if p.dough != nil {
		result.WriteString(p.dough.Something() + "\n")
	}
	if p.sauce != nil {
		result.WriteString(p.sauce.Something() + "\n")
	}
	if p.cheese != nil {
		result.WriteString(p.cheese.Something() + "\n")
	}
	if len(p.veggies) > 0 {
		for i, v := range p.veggies {
			result.WriteString(v.Something())
			if i < len(p.veggies)-1 {
				result.WriteString(", ")
			}
		}
		result.WriteString("\n")
	}
	if p.clams != nil {
		result.WriteString(p.clams.Something() + "\n")
	}
	if p.pepperoni != nil {
		result.WriteString(p.pepperoni.Something() + "\n")
	}
	return result.String()
}

// --- Concrete Pizzas ---

// CheesePizza structure
type CheesePizza struct {
	*pizza
	ingredientFactory PizzaIngredientFactory
}

// NewCheesePizza constructor
func NewCheesePizza(name string, factory PizzaIngredientFactory) *CheesePizza {
	return &CheesePizza{
		pizza:             &pizza{name: name},
		ingredientFactory: factory,
	}
}

// Prepare implements the preparation steps for CheesePizza using abstract factory
func (c *CheesePizza) Prepare() {
	fmt.Printf("Preparing %s\n", c.name)
	c.dough = c.ingredientFactory.CreateDough()
	c.sauce = c.ingredientFactory.CreateSauce()
	c.cheese = c.ingredientFactory.CreateCheese()
}

// ClamPizza structure
type ClamPizza struct {
	*pizza
	ingredientFactory PizzaIngredientFactory
}

// NewClamPizza constructor
func NewClamPizza(name string, factory PizzaIngredientFactory) *ClamPizza {
	return &ClamPizza{
		pizza:             &pizza{name: name},
		ingredientFactory: factory,
	}
}

// Prepare implements the preparation steps for ClamPizza
func (c *ClamPizza) Prepare() {
	fmt.Printf("Preparing %s\n", c.name)
	c.dough = c.ingredientFactory.CreateDough()
	c.sauce = c.ingredientFactory.CreateSauce()
	c.cheese = c.ingredientFactory.CreateCheese()
	c.clams = c.ingredientFactory.CreateClams()
}

// VeggiePizza structure
type VeggiePizza struct {
	*pizza
	ingredientFactory PizzaIngredientFactory
}

// NewVeggiePizza constructor
func NewVeggiePizza(name string, factory PizzaIngredientFactory) *VeggiePizza {
	return &VeggiePizza{
		pizza:             &pizza{name: name},
		ingredientFactory: factory,
	}
}

// Prepare implements the preparation steps for VeggiePizza
func (c *VeggiePizza) Prepare() {
	fmt.Printf("Preparing %s\n", c.name)
	c.dough = c.ingredientFactory.CreateDough()
	c.sauce = c.ingredientFactory.CreateSauce()
	c.cheese = c.ingredientFactory.CreateCheese()
	c.veggies = c.ingredientFactory.CreateVeggies()
}

// PepperoniPizza structure
type PepperoniPizza struct {
	*pizza
	ingredientFactory PizzaIngredientFactory
}

// NewPepperoniPizza constructor
func NewPepperoniPizza(name string, factory PizzaIngredientFactory) *PepperoniPizza {
	return &PepperoniPizza{
		pizza:             &pizza{name: name},
		ingredientFactory: factory,
	}
}

// Prepare implements the preparation steps for PepperoniPizza
func (c *PepperoniPizza) Prepare() {
	fmt.Printf("Preparing %s\n", c.name)
	c.dough = c.ingredientFactory.CreateDough()
	c.sauce = c.ingredientFactory.CreateSauce()
	c.cheese = c.ingredientFactory.CreateCheese()
	c.veggies = c.ingredientFactory.CreateVeggies()
	c.pepperoni = c.ingredientFactory.CreatePepperoni()
}
