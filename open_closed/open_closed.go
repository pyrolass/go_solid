package openclosed

// Open Closed says that classes should be open for extension but not modification
// if we have a single struct for calculating the horse power we would need an
// if statement for each car which then modifies the struct each time we need to
// calculate with this approach we fix this issue by having a single interface
// that each car needs to implement.

type HorsePowerCalculator interface {
	CalculateHorsePower() float64
}

type M3 struct{}

func NewM3() HorsePowerCalculator {
	return &M3{}
}

func (*M3) CalculateHorsePower() float64 {
	return 473
}

type Porche911 struct{}

func NewPorche911() HorsePowerCalculator {
	return &Porche911{}
}

func (*Porche911) CalculateHorsePower() float64 {
	return 388
}
