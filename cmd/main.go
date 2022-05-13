package main

import (
	"flip/pkg/carRental"
	"flip/pkg/domain/cars"
	"flip/pkg/domain/stations"
	"fmt"
)

func main() {

	rentalApp := carRental.NewCarRentalApp("flip")

	c1 := cars.NewCar(cars.SedanCarType)
	c2 := cars.NewCar(cars.SedanCarType)

	st1 := stations.NewStation("bellandur", 0)
	st1.AddCar(c1)

	rentalApp.AddStation(*st1)

	st1.AddCar(c2)

	rentalApp.UpdateStation(*st1)

	st := rentalApp.GetStation("bellandur")

	// cars := st.AvailableCars
	for i, c := range st.AvailableCars {
		fmt.Println(i, c.GetLicenseNo())
	}

	b1, _ := rentalApp.BookVehicle(c1, *st1)
	b3, _ := rentalApp.BookVehicle(c2, *st1)

	b3, err := rentalApp.BookVehicle(c1, *st1)
	fmt.Println(b3 == nil)
	fmt.Println(err)

	fmt.Println(b1.ID)
	rentalApp.CloseBooking(*b1, *st1)
	abc := rentalApp.CloseBooking(*b1, *st1)
	fmt.Println(abc)
}
