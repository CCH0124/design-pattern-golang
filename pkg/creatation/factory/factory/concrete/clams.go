package concrete

type FreshClams struct{}

func (f *FreshClams) Something() string {
	return "Fresh Clams"
}

type FrozenClams struct{}

func (f *FrozenClams) Something() string {
	return "Frozen Clams"
}
