package stations

import (
	"flip/pkg/domain/cars"
)

type Station struct {
	Name          string
	Location      int
	AvailableCars []cars.ICar
}

func NewStation(name string, location int) *Station {
	return &Station{
		Name:     name,
		Location: location,
	}
}

func (s *Station) AddCar(car cars.ICar) {
	s.AvailableCars = append(s.AvailableCars, car)

	// sR := GetStationDb()
	// sR.UpdateStation(*s)
}

func (s *Station) GetStationReport() {

}

func (s *Station) DelistCar(car cars.ICar) {
	for _, c := range s.AvailableCars {
		if c.GetLicenseNo() == car.GetLicenseNo() {
			c.SetCarStatus(true)

		}
	}
}

func (s *Station) RelistCar(car cars.ICar) {
	for _, c := range s.AvailableCars {
		if c.GetLicenseNo() == car.GetLicenseNo() {
			c.SetCarStatus(false)

		}
	}
}

func (s *Station) ValidateCar(car cars.ICar) bool {
	for _, c := range s.AvailableCars {
		if c.GetLicenseNo() == car.GetLicenseNo() && !c.GetCarStatus() {
			return true

		}
	}
	return false
}
