package main

import (
	"fmt"
	"practice/parking-lot/entry-gate"
	"practice/parking-lot/exit-gate"
	parking_lot "practice/parking-lot/parking-lot"
	parking_type "practice/parking-lot/parking-type"
	"practice/parking-lot/payment"
	"practice/parking-lot/ticket"
)

func main() {
	carSpot := []string{"A1", "A2", "A3"}
	bikeSpot := []string{"B1"}
	hcSpot := []string{"C1", "C2", "C3"}
	carParkingType := parking_type.CreateNewParkingType("Car", carSpot, 10.0)
	bikeParkingType := parking_type.CreateNewParkingType("Bike", bikeSpot, 5.0)
	hcParkingType := parking_type.CreateNewParkingType("HC", hcSpot, 0.0)

	var tickets []ticket.Ticket
	parkingSpace := []parking_type.Parking{carParkingType, bikeParkingType, hcParkingType}
	parkingLot := parking_lot.CreateNewParkingLot(parkingSpace)
	entry_gate := entry.NewEntry(parkingLot)
	pay := payment.NewPayment(10)
	exit_gate := exit.NewExit(parkingLot, pay)

	e, t := entry_gate.CreateTicket("AB123", "Bike")
	if e != nil {
		fmt.Println("error:", e)
	} else {
		tickets = append(tickets, t)
	}
	e, t = entry_gate.CreateTicket("CD123", "Car")
	if e != nil {
		fmt.Println("error:", e)
	} else {
		tickets = append(tickets, t)
	}
	e, t = entry_gate.CreateTicket("EF123", "HC")
	if e != nil {
		fmt.Println("error:", e)
	} else {
		tickets = append(tickets, t)
	}
	e, t = entry_gate.CreateTicket("GH123", "Bike")
	if e != nil {
		fmt.Println("error:", e)
	} else {
		tickets = append(tickets, t)
	}

	exit_gate.ProcessVehicleExit(tickets[0])
	exit_gate.ProcessVehicleExit(tickets[1])
	for _, ticket := range tickets {
		fmt.Printf("before:%v\n", ticket)
	}

	e, t = entry_gate.CreateTicket("GH123", "Bike")
	if e != nil {
		fmt.Println("error:", e)
	} else {
		tickets = append(tickets, t)
	}
	for _, ticket := range tickets {
		fmt.Printf("after:%v\n", ticket)
	}

}
