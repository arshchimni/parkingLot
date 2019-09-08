package car

type Car struct {
	RegNo string
	Color string
}

func NewCar(regNo, color string) Car {
	return Car{RegNo: regNo, Color: color}
}

func (c Car) GetRegNo() string {
	return c.RegNo
}

func (c Car) GetColor() string {
	return c.Color
}
