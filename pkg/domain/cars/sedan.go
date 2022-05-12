package cars

import (
	"github.com/google/uuid"
)

type SedanCar struct {
	Car
}

func NewSedan() ICar {
	return &SedanCar{
		Car: Car{
			ID:   uuid.New().String(),
			Type: SedanCarType,
		},
	}
}

func (nS *SedanCar) GetLicenseNo() string {
	return nS.ID
}
