package concrete

type MarinaraSauce struct{}

func (m *MarinaraSauce) Something() string {
	return "Marinara Sauce"
}

type PlumTomatoSauce struct{}

func (p *PlumTomatoSauce) Something() string {
	return "Plum Tomato Sauce"
}
