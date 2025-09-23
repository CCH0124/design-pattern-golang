package base

type PizzaFlow interface {
	Prepare()
	Bake()
	Cut()
	Box()
	GetName() string
}
