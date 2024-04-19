package parking_type

import "practice/parking-lot/ticket"

type Parking interface {
	GetParkingType() string
	GetRate() float64
	Book() string
	UnBook(t ticket.Ticket)
	IsSpaceLeft() bool
}

type ParkingType struct {
	parkingType string
	spaceMap    map[string]bool
	rate        float64
}

func (p ParkingType) GetParkingType() string {
	return p.parkingType

}
func (p ParkingType) GetRate() float64 {
	return p.rate

}

func (p ParkingType) IsSpaceLeft() bool {
	spaces := p.spaceMap
	for _, v := range spaces {
		if !v {
			return true
		}
	}
	return false

}

func (p ParkingType) Book() string {
	spaces := p.spaceMap
	for k, v := range spaces {
		if !v {
			spot := k
			spaces[k] = true
			return spot
		}
	}
	return ""

}
func (p ParkingType) UnBook(t ticket.Ticket) {
	p.spaceMap[t.GetParkingSpot()] = false

}
func CreateNewParkingType(parkingType string, spaceMap []string, rate float64) Parking {
	mp := make(map[string]bool)
	for _, v := range spaceMap {
		mp[v] = false

	}
	return ParkingType{
		parkingType: parkingType,
		spaceMap:    mp,
		rate:        rate,
	}

}
