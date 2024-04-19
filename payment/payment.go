package payment

import (
	"fmt"
	"math"
	"practice/parking-lot/ticket"
)

type Payment interface {
	CalculatePayment(t ticket.Ticket, rate float64) float64
	IsPaid(ticket.Ticket) bool
}
type OfficePayment struct {
	defaultRate float64
}

func (p OfficePayment) CalculatePayment(t ticket.Ticket, rate float64) float64 {
	t1 := t.GetCheckInTime()
	t2 := t.GetCheckOutTime()
	timeElapsed := math.Ceil(t2.Sub(t1).Minutes())
	return (timeElapsed) * (rate)
}

func (p OfficePayment) IsPaid(t ticket.Ticket) bool {
	var isPaid bool
	fmt.Scanln("payment done for ", t.GetVehicleID(), &isPaid)
	return isPaid
}

func NewPayment(rate float64) Payment {
	return OfficePayment{defaultRate: rate}
}
