package concrete

import "designpattern/pkg/creatation/factory/factory/ingredient"

type ChicagoPizzaIngredientFactory struct{}

func (c *ChicagoPizzaIngredientFactory) CreateDough() ingredient.Dough {
	return &ThickCrustDough{}
}

func (c *ChicagoPizzaIngredientFactory) CreateSauce() ingredient.Sauce {
	return &PlumTomatoSauce{}
}

func (c *ChicagoPizzaIngredientFactory) CreateCheese() ingredient.Cheese {
	return &MozzarellaCheese{}
}

func (c *ChicagoPizzaIngredientFactory) CreateVeggies() []ingredient.Veggies {
	return []ingredient.Veggies{
		&BlackOlives{},
		&Spinach{},
		&Eggplant{},
	}
}

func (c *ChicagoPizzaIngredientFactory) CreatePepperoni() ingredient.Pepperoni {
	return &SlicedPepperoni{}
}

func (c *ChicagoPizzaIngredientFactory) CreateClams() ingredient.Clams {
	return &FrozenClams{}
}
