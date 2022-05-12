package cars

import "github.com/google/uuid"

type SUVCar struct {
	Car
}

func NewSUV() ICar {
	return &SUVCar{
		Car: Car{
			ID:   uuid.New().String(),
			Type: SUVCarType,
		},
	}
}

func (nS *SUVCar) GetLicenseNo() string {
	return nS.License
}
