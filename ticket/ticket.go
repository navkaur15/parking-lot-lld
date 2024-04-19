package ticket

import (
	"time"
)

type Ticket interface {
	GetVehicleID() string
	GetVehicleType() string
	GetCheckInTime() time.Time
	GetCheckOutTime() time.Time
	GetParkingSpot() string
	SetCheckOutTime(checkoutTime time.Time)
	SetAmoutPaid(price float64)
	SetStatus(status string)
}
type OfficeTicket struct {
	vehicleID    string
	vehicleType  string
	checkinTime  time.Time
	parkingSpot  string
	checkoutTime time.Time
	amountPaid   float64
	status       string
}

func (t *OfficeTicket) GetVehicleID() string {
	return t.vehicleID
}
func (t *OfficeTicket) GetVehicleType() string {
	return t.vehicleType
}
func (t *OfficeTicket) GetCheckInTime() time.Time {
	return t.checkinTime
}
func (t *OfficeTicket) GetCheckOutTime() time.Time {
	return t.checkoutTime
}
func (t *OfficeTicket) GetParkingSpot() string {
	return t.parkingSpot
}
func (t *OfficeTicket) SetCheckOutTime(checkoutTime time.Time) {
	t.checkoutTime = checkoutTime
}
func (t *OfficeTicket) SetAmoutPaid(price float64) {
	t.amountPaid = price
}
func (t *OfficeTicket) SetStatus(status string) {
	t.status = status
}

func GetNewOfficeTicket(vehicleID, vehicleType, parkingSpot, status string) Ticket {
	return &OfficeTicket{
		vehicleID:   vehicleID,
		vehicleType: vehicleType,
		checkinTime: time.Now(),
		parkingSpot: parkingSpot,
		status:      status,
	}

}
