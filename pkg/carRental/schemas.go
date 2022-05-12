package carRental

import (
	"flip/pkg/domain/bookings"
	"flip/pkg/domain/cars"
	"flip/pkg/domain/stations"
)

type CarRentalApp struct {
	name     string
	stations []*stations.Station
	bookings []*bookings.Booking
}

type ICarRentalApp interface {
	AddStation(station stations.Station)
	UpdateStation(station stations.Station)
	GetStation(name string) *stations.Station
	// SearchVehicle(Type string)

	BookVehicle(car cars.ICar, station stations.Station) bookings.Booking
	CloseBooking(booking bookings.Booking, drop stations.Station)
}

var CarRentalAppDB *CarRentalApp

func NewCarRentalApp(name string) *CarRentalApp {

	if CarRentalAppDB == nil {
		CarRentalAppDB = &CarRentalApp{
			name: name,
		}
		return CarRentalAppDB

	}
	return CarRentalAppDB
}
