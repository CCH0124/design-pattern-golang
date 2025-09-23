package concrete

type Garlic struct{}

func (g *Garlic) Something() string {
	return "Garlic"
}

type Mushroom struct{}

func (m *Mushroom) Something() string {
	return "Mushroom"
}

type Onion struct{}

func (o *Onion) Something() string {
	return "Onion"
}

type RedPepper struct{}

func (r *RedPepper) Something() string {
	return "Red Pepper"
}

type Spinach struct{}

func (s *Spinach) Something() string {
	return "Spinach"
}
