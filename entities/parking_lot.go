package entities

type ParkingSlots struct {
	TotalSlots    int
	OccupiedSlots int
	SlotDetails   []SlotDetails
}

type SlotDetails struct {
	Occupied   bool
	CarDetails CarDetails
}

type SlotData struct {
	Level  int
	SlotId int
}

type CarDetails struct {
	CarNo string
	Color string
}
