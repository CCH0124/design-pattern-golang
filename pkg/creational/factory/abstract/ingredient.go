package abstract

// --- Ingredient Interfaces ---

// Dough defines the interface for dough ingredients.
type Dough interface {
	Something() string
}

// Sauce defines the interface for sauce ingredients.
type Sauce interface {
	Something() string
}

// Cheese defines the interface for cheese ingredients.
type Cheese interface {
	Something() string
}

// Veggies defines the interface for veggie ingredients.
type Veggies interface {
	Something() string
}

// Pepperoni defines the interface for pepperoni ingredients.
type Pepperoni interface {
	Something() string
}

// Clams defines the interface for clams ingredients.
type Clams interface {
	Something() string
}

// --- Abstract Factory Interface ---

// PizzaIngredientFactory defines the interface for creating families of ingredients.
type PizzaIngredientFactory interface {
	CreateDough() Dough
	CreateSauce() Sauce
	CreateCheese() Cheese
	CreateVeggies() []Veggies
	CreatePepperoni() Pepperoni
	CreateClams() Clams
}

// --- Concrete Ingredients ---

// ThickCrustDough implementation
type ThickCrustDough struct{}

func (d *ThickCrustDough) Something() string { return "Thick Crust Dough" }

// ThinCrustDough implementation
type ThinCrustDough struct{}

func (d *ThinCrustDough) Something() string { return "Thin Crust Dough" }

// PlumTomatoSauce implementation
type PlumTomatoSauce struct{}

func (s *PlumTomatoSauce) Something() string { return "Plum Tomato Sauce" }

// MarinaraSauce implementation
type MarinaraSauce struct{}

func (s *MarinaraSauce) Something() string { return "Marinara Sauce" }

// MozzarellaCheese implementation
type MozzarellaCheese struct{}

func (c *MozzarellaCheese) Something() string { return "Shredded Mozzarella Cheese" }

// ReggianoCheese implementation
type ReggianoCheese struct{}

func (c *ReggianoCheese) Something() string { return "Grated Reggiano Cheese" }

// SlicedPepperoni implementation
type SlicedPepperoni struct{}

func (p *SlicedPepperoni) Something() string { return "Sliced Pepperoni" }

// FrozenClams implementation
type FrozenClams struct{}

func (c *FrozenClams) Something() string { return "Frozen Clams" }

// FreshClams implementation
type FreshClams struct{}

func (c *FreshClams) Something() string { return "Fresh Clams" }

// Garlic implementation
type Garlic struct{}

func (v *Garlic) Something() string { return "Garlic" }

// Onion implementation
type Onion struct{}

func (v *Onion) Something() string { return "Onion" }

// Mushroom implementation
type Mushroom struct{}

func (v *Mushroom) Something() string { return "Mushroom" }

// RedPepper implementation
type RedPepper struct{}

func (v *RedPepper) Something() string { return "Red Pepper" }

// BlackOlives implementation
type BlackOlives struct{}

func (v *BlackOlives) Something() string { return "Black Olives" }

// Eggplant implementation
type Eggplant struct{}

func (v *Eggplant) Something() string { return "Eggplant" }

// Spinach implementation
type Spinach struct{}

func (v *Spinach) Something() string { return "Spinach" }

// --- Concrete Ingredient Factories ---

// NYPizzaIngredientFactory implementation
type NYPizzaIngredientFactory struct{}

func (n *NYPizzaIngredientFactory) CreateDough() Dough { return &ThinCrustDough{} }
func (n *NYPizzaIngredientFactory) CreateSauce() Sauce { return &MarinaraSauce{} }
func (n *NYPizzaIngredientFactory) CreateCheese() Cheese { return &ReggianoCheese{} }
func (n *NYPizzaIngredientFactory) CreateVeggies() []Veggies {
	return []Veggies{&Garlic{}, &Onion{}, &Mushroom{}, &RedPepper{}}
}
func (n *NYPizzaIngredientFactory) CreatePepperoni() Pepperoni { return &SlicedPepperoni{} }
func (n *NYPizzaIngredientFactory) CreateClams() Clams         { return &FreshClams{} }

// ChicagoPizzaIngredientFactory implementation
type ChicagoPizzaIngredientFactory struct{}

func (c *ChicagoPizzaIngredientFactory) CreateDough() Dough { return &ThickCrustDough{} }
func (c *ChicagoPizzaIngredientFactory) CreateSauce() Sauce { return &PlumTomatoSauce{} }
func (c *ChicagoPizzaIngredientFactory) CreateCheese() Cheese { return &MozzarellaCheese{} }
func (c *ChicagoPizzaIngredientFactory) CreateVeggies() []Veggies {
	return []Veggies{&BlackOlives{}, &Spinach{}, &Eggplant{}}
}
func (c *ChicagoPizzaIngredientFactory) CreatePepperoni() Pepperoni { return &SlicedPepperoni{} }
func (c *ChicagoPizzaIngredientFactory) CreateClams() Clams         { return &FrozenClams{} }
