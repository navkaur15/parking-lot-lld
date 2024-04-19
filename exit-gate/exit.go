package exit

import (
	"fmt"
	parking_lot "practice/parking-lot/parking-lot"
	"practice/parking-lot/payment"
	"practice/parking-lot/ticket"
	"time"
)

type OfficeExit struct {
	parkingLot parking_lot.ParkingLot
	payment    payment.Payment
}

type Exit interface {
	ProcessVehicleExit(t ticket.OfficeTicket)
}

func (e OfficeExit) ProcessVehicleExit(t ticket.Ticket) {
	parking := e.parkingLot
	rate := parking.UnParkVehicle(t)
	//Payment
	t.SetCheckOutTime(time.Now())
	price := e.payment.CalculatePayment(t, rate)
	fmt.Println("PricePaid", price)
	t.SetAmoutPaid(price)
	t.SetStatus("Exited")
}

func NewExit(parking parking_lot.ParkingLot, payment payment.Payment) OfficeExit {
	return OfficeExit{
		parkingLot: parking,
		payment:    payment,
	}
}
