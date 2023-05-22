package exit_car_handler

import (
	"errors"
	"fmt"
	"github.com/dunzoit/projects/parking_lot_problem_detailed_sol/entities"
)

type ExitCar struct {
	ParkingLot []entities.ParkingSlots
	CarSlotMap map[string]entities.SlotData
}

func NewExitCarHandler(parkingLot []entities.ParkingSlots, carSlotMap map[string]entities.SlotData) *ExitCar {
	return &ExitCar{
		ParkingLot: parkingLot,
		CarSlotMap: carSlotMap,
	}
}

type ExitCarHandler interface {
	ExitCar(carDetails entities.CarDetails) (*entities.SlotData, error)
}

func (p *ExitCar) ExitCar(carDetails entities.CarDetails) (*entities.SlotData, error) {
	slotData, ok := p.CarSlotMap[fmt.Sprintf("%v|%v", carDetails.CarNo, carDetails.Color)]
	if !ok {
		return nil, errors.New("car not found in parking😞")
	}
	delete(p.CarSlotMap, fmt.Sprintf("%v|%v", carDetails.CarNo, carDetails.Color))
	p.ParkingLot[slotData.Level].SlotDetails[slotData.SlotId] = entities.SlotDetails{}
	p.ParkingLot[slotData.Level].OccupiedSlots--
	return &slotData, nil
}
