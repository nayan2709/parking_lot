package display_car_details

import (
	"errors"
	"fmt"
	"github.com/dunzoit/projects/parking_lot_problem_detailed_sol/entities"
)

type DisplayCar struct {
	CarSlotMap map[string]entities.SlotData
}

func NewDisplayCarHandler(carSlotMap map[string]entities.SlotData) *DisplayCar {
	return &DisplayCar{
		CarSlotMap: carSlotMap,
	}
}

type DisplayCarHandler interface {
	DisplayCarDetails(carDetails entities.CarDetails) (*entities.SlotData, error)
}

func (p *DisplayCar) DisplayCarDetails(carDetails entities.CarDetails) (*entities.SlotData, error) {
	slotData, ok := p.CarSlotMap[fmt.Sprintf("%v|%v", carDetails.CarNo, carDetails.Color)]
	if !ok {
		return nil, errors.New("car not found in parking😞")
	}
	return &slotData, nil
}
