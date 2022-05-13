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
	// c2 := cars.NewCar(cars.SUVCarType)

	st1 := stations.NewStation("bellandur", 0)
	st2 := stations.NewStation("hsr", 0)
	st1.AddCar(c1)

	rentalApp.AddStation(*st1)
	rentalApp.AddStation(*st2)

	// st1.AddCar(c2)

	rentalApp.UpdateStation(*st1)

	st := rentalApp.GetStation("bellandur")

	// cars := st.AvailableCars
	for i, c := range st.AvailableCars {
		fmt.Println(i, c.GetLicenseNo())
	}

	fmt.Println("c1 booked", c1.GetCarStatus())
	b1, _ := rentalApp.BookVehicle(c1, st1.Name)
	fmt.Println("c1 booked", c1.GetCarStatus())

	_, err := rentalApp.BookVehicle(c1, st1.Name)
	fmt.Println(err)

	rentalApp.CloseBooking(*b1, st2.Name)
	// err = rentalApp.CloseBooking(*b1, *st2)
	// fmt.Println(err)
	fmt.Println("c1 booked", c1.GetCarStatus())

	_, _ = rentalApp.BookVehicle(c1, st1.Name)
	fmt.Println("c1 booked", c1.GetCarStatus())

	_, _ = rentalApp.BookVehicle(c1, st2.Name)
	fmt.Println("c2 booked", c1.GetCarStatus())

}
