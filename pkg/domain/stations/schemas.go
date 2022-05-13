package stations

import (
	"errors"
	"flip/pkg/domain/cars"
)

type Station struct {
	Name string

	// considering 1D plane for distance calculation
	Location int

	AvailableCars []cars.ICar
}

func NewStation(name string, location int) *Station {
	return &Station{
		Name:     name,
		Location: location,
	}
}

func (s *Station) AddCar(car cars.ICar) error {
	dbCar := s.GetCar(car.GetLicenseNo())

	if dbCar == nil {
		s.AvailableCars = append(s.AvailableCars, car)
		return nil

	}
	return errors.New("Car already Added")
}

func (s *Station) DropCar(i int) error {
	if len(s.AvailableCars) == 1 {
		s.AvailableCars = []cars.ICar{}
		return nil
	}

	s.AvailableCars = append(s.AvailableCars[:i], s.AvailableCars[i+1:]...)
	return nil
}

func (s *Station) GetCar(carID string) cars.ICar {
	for _, c := range s.AvailableCars {
		if carID == c.GetLicenseNo() {
			return c
		}
	}
	return nil
}

func (s *Station) GetStationReport() {

}

func (s *Station) DelistCar(car cars.ICar) error {

	// todo replace with get car and updatecar methods
	for i, c := range s.AvailableCars {
		if c.GetLicenseNo() == car.GetLicenseNo() {
			c.SetCarStatus(true)
			s.DropCar(i)
			return nil

		}
	}
	return errors.New("Car not found")
}

func (s *Station) RelistCar(car cars.ICar) {

	// todo replace with get car and updatecar methods
	car.SetCarStatus(false)
	s.AddCar(car)
}

func (s *Station) ValidateCar(car cars.ICar) bool {
	for _, c := range s.AvailableCars {
		if c.GetLicenseNo() == car.GetLicenseNo() && !c.GetCarStatus() {
			return true

		}
	}
	return false
}
