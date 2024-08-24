package openclosed

// OCP suggests that software entities should be open for extension but closed for modification.
// Go's interfaces make this principle easy to implement.

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
