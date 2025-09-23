package concrete

type MozzarellaCheese struct{}

func (m *MozzarellaCheese) Something() string {
	return "Mozzarella Cheese"
}

type ReggianoCheese struct{}

func (r *ReggianoCheese) Something() string {
	return "Reggiano Cheese"
}
