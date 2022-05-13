package carRental

import (
	"errors"
	"flip/pkg/domain/bookings"
	"flip/pkg/domain/cars"
	"flip/pkg/domain/stations"
)

func (cR *CarRentalApp) AddStation(station stations.Station) {

	dbStation := cR.GetStation(station.Name)
	if dbStation == nil {
		cR.stations = append(cR.stations, &station)

	}

}

func (cR *CarRentalApp) GetStation(name string) *stations.Station {
	for _, st := range cR.stations {
		if name == st.Name {
			return st
		}
	}
	return nil
}

func (cR *CarRentalApp) UpdateStation(station stations.Station) {

	for _, st := range cR.stations {
		if station.Name == st.Name {
			st.AvailableCars = station.AvailableCars

		}
	}
}

func (cR *CarRentalApp) BookVehicle(car cars.ICar, pickup stations.Station) (*bookings.Booking, error) {

	if pickup.ValidateCar(car) {
		newBooking := bookings.NewBooking(car)
		newBooking.SetPickup(pickup)

		cR.bookings = append(cR.bookings, &newBooking)
		pickup.DelistCar(car)

		return &newBooking, nil

	}
	return nil, errors.New("Car not available")

}

func (cR *CarRentalApp) CloseBooking(booking bookings.Booking, drop stations.Station) error {

	for _, b := range cR.bookings {
		if b.ID == booking.ID {
			b.SetDrop(drop)
			err := drop.RelistCar(b.GetCar())
			if err != nil {

				return err
			}
			b.CloseBooking()
			return nil
		}

	}

	return errors.New("Invalid booking")
}
