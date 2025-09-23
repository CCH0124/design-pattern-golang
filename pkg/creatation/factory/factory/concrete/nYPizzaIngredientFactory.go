package concrete

import "designpattern/pkg/creatation/factory/factory/ingredient"

type NYPizzaIngredientFactory struct{}

func (n *NYPizzaIngredientFactory) CreateDough() ingredient.Dough {
	return &ThinCrustDough{}
}

func (n *NYPizzaIngredientFactory) CreateSauce() ingredient.Sauce {
	return &MarinaraSauce{}
}

func (n *NYPizzaIngredientFactory) CreateCheese() ingredient.Cheese {
	return &ReggianoCheese{}
}

func (n *NYPizzaIngredientFactory) CreateVeggies() []ingredient.Veggies {
	return []ingredient.Veggies{
		&Garlic{},
		&Onion{},
		&Mushroom{},
		&RedPepper{},
	}
}

func (n *NYPizzaIngredientFactory) CreatePepperoni() ingredient.Pepperoni {
	return &SlicedPepperoni{}
}

func (n *NYPizzaIngredientFactory) CreateClams() ingredient.Clams {
	return &FreshClams{}
}
