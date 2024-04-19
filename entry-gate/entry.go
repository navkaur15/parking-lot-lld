package entry

import (
	"errors"
	parking_lot "practice/parking-lot/parking-lot"
	"practice/parking-lot/ticket"
)

type OfficeEntry struct {
	parkingLot parking_lot.ParkingLot
}
type Entry interface {
	CreateTicket(vehicleID, vehicleType string) (error, ticket.Ticket)
}

func (e OfficeEntry) CreateTicket(vehicleID, vehicleType string) (error, ticket.Ticket) {
	parking := e.parkingLot
	spot := parking.ParkVehicle(vehicleType)
	status := "Parked"
	if spot != "" {
		ticket := ticket.GetNewOfficeTicket(vehicleID, vehicleType, spot, status)
		return nil, ticket
	}
	return errors.New("No space left"), &ticket.OfficeTicket{}

}

func NewEntry(parking parking_lot.ParkingLot) OfficeEntry {
	return OfficeEntry{
		parkingLot: parking,
	}
}
