package parking_lot

import (
	parking_type "practice/parking-lot/parking-type"
	"practice/parking-lot/ticket"
)

type ParkingInterface interface {
	ParkVehicle() string
	UnParkVehicle(t ticket.Ticket) float64
}
type ParkingLot struct {
	parkingSpaces []parking_type.Parking
}

func (p ParkingLot) ParkVehicle(vehicleType string) string {
	var parkSpace parking_type.Parking
	for _, parkType := range p.parkingSpaces {
		if parkType.GetParkingType() == vehicleType {
			parkSpace = parkType
			break
		}
	}
	if parkSpace != nil {
		if parkSpace.IsSpaceLeft() {
			spot := parkSpace.Book()
			return spot
		}
	}
	return ""

}

func (p ParkingLot) UnParkVehicle(t ticket.Ticket) float64 {
	var parkSpace parking_type.Parking
	rate := 0.0
	for _, parkType := range p.parkingSpaces {
		if parkType.GetParkingType() == t.GetVehicleType() {
			parkSpace = parkType
			break
		}
	}
	if parkSpace != nil {
		rate = parkSpace.GetRate()
		parkSpace.UnBook(t)

	}
	return rate

}

func CreateNewParkingLot(parkingSpaces []parking_type.Parking) ParkingLot {
	return ParkingLot{
		parkingSpaces: parkingSpaces,
	}

}
