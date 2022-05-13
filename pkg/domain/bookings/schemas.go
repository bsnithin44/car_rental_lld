package bookings

import (
	"flip/pkg/domain/cars"
	"flip/pkg/domain/stations"

	"github.com/google/uuid"
)

type Booking struct {
	ID            string
	car           cars.ICar
	pickUpStation stations.Station
	dropStation   stations.Station
	isActive      bool
}

func NewBooking(car cars.ICar) Booking {
	return Booking{
		car:      car,
		ID:       uuid.NewString(),
		isActive: true,
	}
}

func (b *Booking) SetPickup(station stations.Station) {
	b.pickUpStation = station
}

func (b *Booking) SetDrop(station stations.Station) {
	b.dropStation = station
}

func (b *Booking) CloseBooking() {
	b.isActive = false
}
func (b *Booking) GetCar() cars.ICar {
	return b.car
}
