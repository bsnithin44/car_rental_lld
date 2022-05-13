package cars

type CarType string

type carPrice int

const (
	SUVCarType       CarType = "SUV"
	HatchBackCarType CarType = "Hatchback"
	SedanCarType     CarType = "sedan"

	SUVPrice   carPrice = 11
	SedanPrice carPrice = 20
)

type Car struct {
	Name    string
	Brand   string
	License string

	ID       string
	Type     CarType
	Price    int
	isBooked bool
}

type ICar interface {
	GetLicenseNo() string
	SetCarStatus(bool)
	GetCarStatus() bool
}

func NewCar(typeOfCar CarType) ICar {

	switch typeOfCar {

	case SedanCarType:
		return NewSedan()

	case SUVCarType:
		return NewSedan()
	}
	return nil
}

func (c *Car) SetCarStatus(booked bool) {
	c.isBooked = booked

}

func (c *Car) GetCarStatus() bool {
	return c.isBooked

}
