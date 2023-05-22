package park_car_handler

import (
	"errors"
	"fmt"
	"github.com/dunzoit/projects/parking_lot_problem_detailed_sol/entities"
)

type ParkCar struct {
	ParkingLot []entities.ParkingSlots
	CarSlotMap map[string]entities.SlotData
}

func NewParkCarHandler(parkingLot []entities.ParkingSlots, carSlotMap map[string]entities.SlotData) *ParkCar {
	return &ParkCar{
		ParkingLot: parkingLot,
		CarSlotMap: carSlotMap,
	}
}

type ParkCarHandler interface {
	ParkCar(carDetails entities.CarDetails) (*entities.SlotData, error)
}

func (p *ParkCar) ParkCar(carDetails entities.CarDetails) (*entities.SlotData, error) {
	var parkingAllotted bool
	var slotData entities.SlotData
	for level := 0; level < len(p.ParkingLot); level++ {
		if p.ParkingLot[level].TotalSlots > p.ParkingLot[level].OccupiedSlots {
			for slotId := range p.ParkingLot[level].SlotDetails {
				if p.ParkingLot[level].SlotDetails[slotId].Occupied {
					continue
				}
				p.ParkingLot[level].SlotDetails[slotId].CarDetails = carDetails
				p.ParkingLot[level].SlotDetails[slotId].Occupied = true
				p.ParkingLot[level].OccupiedSlots++
				slotData = entities.SlotData{Level: level, SlotId: slotId}
				p.CarSlotMap[fmt.Sprintf("%v|%v", carDetails.CarNo, carDetails.Color)] = slotData
				parkingAllotted = true
				break
			}
			break
		}
	}
	if !parkingAllotted {
		return nil, errors.New("all slots filled 😞")
	}
	return &slotData, nil
}
