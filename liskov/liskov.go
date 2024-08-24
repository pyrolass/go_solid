package liskov

// LSP states that objects of a superclass should be replaceable with objects of its subclasses without affecting the correctness of the program.
// In Go, we can demonstrate this with interfaces.
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
