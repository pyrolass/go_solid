package liskov

// Liskov's principle says that child class should be able to do everything that the
// parent class does
// in this example both M3 and Altim extend a common interface which is drive
// both cars are drivable but on is fast

type ICar interface {
	Drive() bool
}

type IFastCar interface {
	ICar
	Fast() bool
}

type M3 struct{}

type Altima struct{}

func NewM3() IFastCar {
	return &M3{}
}

func (*M3) Drive() bool {
	return true
}

func (*M3) Fast() bool {
	return true
}

func NewAltima() ICar {
	return &Altima{}
}

func (*Altima) Drive() bool {
	return true
}
