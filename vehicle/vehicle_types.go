package vehicle

import (
	"errors"
)

type VehicleType struct {
	slug string
}

var(
	CAR = VehicleType{"CAR"}
	BIKE = VehicleType{"BIKE"}
	TRUCK = VehicleType{"TRUCK"}
	UNKNOWN = VehicleType{""}
)


func FromString(s string) (VehicleType, error) {
	switch s {
	case CAR.slug:
		return CAR, nil
	case BIKE.slug:
		return BIKE, nil
	case TRUCK.slug:
		return TRUCK, nil
	}

	return UNKNOWN, errors.New("unknown vehicle type " + s)
}

func (v *VehicleType) FromVehicleType()string{
	return v.slug
}

func NumberOfVehcileTypeSupported() int{
	return 3
}