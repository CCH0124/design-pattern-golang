package concrete

type ThickCrustDough struct{}

func (d *ThickCrustDough) Something() string {
	return "Thick Crust Dough"
}

type ThinCrustDough struct{}

func (d *ThinCrustDough) Something() string {
	return "Thin Crust Dough"
}
