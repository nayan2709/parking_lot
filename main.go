package main

import (
	"fmt"
	"github.com/dunzoit/projects/parking_lot_problem_detailed_sol/display_car_details"
	"github.com/dunzoit/projects/parking_lot_problem_detailed_sol/entities"
	"github.com/dunzoit/projects/parking_lot_problem_detailed_sol/exit_car_handler"
	"github.com/dunzoit/projects/parking_lot_problem_detailed_sol/park_car_handler"
	"log"
	"os"
)

//setParkingLot: sets the parking lot with given slots
//-----------------
//INPUTS:
//Enter your operator id: Metro:123
//level: 3, slot: 2
//Operation: parkCar 1 10
//Operation: parkCar 2 20
//Operation: parkCar 3 30
//Operation: exitCar 2 20
//Operation: parkCar 4 40
//Operation: exit 1  1
//------------------
//OUTPUTS:
//😎FedUp with Handling your car in parking, HERE comes the ParkingExpertBot😎
//Valid Operation and Format::
//1.parkCar carNo color
//2.displayDetails carNo color
//3.exitCar carNo color
//4.Exit
//OP: parkCar, Slot Data: &{Level:0 SlotId:0}
//OP: parkCar, Slot Data: &{Level:0 SlotId:1}
//OP: parkCar, Slot Data: &{Level:1 SlotId:0}
//OP: exitCar, Slot Data: &{Level:0 SlotId:1}
//OP: parkCar, Slot Data: &{Level:0 SlotId:1}
//Done for the day, our service will resume our services sometime.🙏

// Running Implementation:
// go run main.go
func main() {
	// open input/output file in split tab
	//openFilesInSplitTabs("/Users/nayan/go/src/github.com/dunzoit/projects/parking_lot_problem_detailed_sol/input.txt", "/Users/nayan/go/src/github.com/dunzoit/projects/parking_lot_problem_detailed_sol/output.txt")
	//Redirect Input
	inputFile, err := os.Open("parking_lot_problem_detailed_sol/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()
	os.Stdin = inputFile

	//Redirect Output
	outputFile, err := os.Create("parking_lot_problem_detailed_sol/output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()
	os.Stdout = outputFile

	fmt.Print("😎FedUp with Handling your car in parking, HERE comes the ParkingExpertBot😎\n")
	var operatorId string
	fmt.Scanf("Enter your operator id: %s\n", &operatorId)
	if operatorId[0:6] != "Metro:" {
		fmt.Println("Not Valid Operator id: ", operatorId)
		return
	}

	var level, slots int
	fmt.Scanf("level: %d, slot: %d\n", &level, &slots)
	parkingLot := make([]entities.ParkingSlots, level)
	setParkingLot(slots, parkingLot)
	carSlotMap := make(map[string]entities.SlotData)
	parkHandler := park_car_handler.NewParkCarHandler(parkingLot, carSlotMap)
	exitHandler := exit_car_handler.NewExitCarHandler(parkingLot, carSlotMap)
	displayHandler := display_car_details.NewDisplayCarHandler(carSlotMap)

	fmt.Printf("\n\nValid Operation and Format::\n" +
		"1.parkCar carNo color\n" +
		"2.displayDetails carNo color\n" +
		"3.exitCar carNo color\n" +
		"4.Exit\n",
	)
	for {
		var op, carNo, color string
		var slotData *entities.SlotData
		var err error
		fmt.Scanf("Operation: %s %s %s\n", &op, &carNo, &color)
		carDetails := entities.CarDetails{
			CarNo: carNo,
			Color: color,
		}
		stop := false
		switch {
		case op == "parkCar":
			slotData, err = parkHandler.ParkCar(carDetails)
		case op == "displayDetails":
			slotData, err = displayHandler.DisplayCarDetails(carDetails)
		case op == "exitCar":
			slotData, err = exitHandler.ExitCar(carDetails)
		default:
			stop = true
		}
		if stop {
			break
		}
		if err != nil {
			fmt.Println(fmt.Sprintf("Error while doing op: %v, err: %v", op, err))
		}
		fmt.Println(fmt.Sprintf("OP: %v, Slot Data: %+v", op, slotData))
	}
	fmt.Println("Done for the day, our service will resume in sometime.🙏")
}

func setParkingLot(slots int, parkingSlots []entities.ParkingSlots) {
	for i := 0; i < len(parkingSlots); i++ {
		var slotDetails []entities.SlotDetails
		for j := 0; j < slots; j++ {
			slotDetails = append(slotDetails, entities.SlotDetails{})
		}
		parkingSlots[i] = entities.ParkingSlots{
			TotalSlots:    slots,
			OccupiedSlots: 0,
			SlotDetails:   slotDetails,
		}
	}
}

//
//func openFilesInSplitTabs(inputFile, outputFile string) error {
//	// Construct the command to open the input and output files in split tabs
//	command := exec.Command("goland", "-p", fmt.Sprintf("-e %s -e %s", inputFile, outputFile))
//
//	// Run the command
//	err := command.Run()
//	if err != nil {
//		return err
//	}
//
//	return nil
//}

/*
// Parking Lot Solution using channels
func main() {
	//Redirect Input
	inputFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()
	os.Stdin = inputFile

	//Redirect Output
	outputFile, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()
	os.Stdout = outputFile

	fmt.Print("😎FedUp with Handling your car in parking, HERE comes the ParkingExpertBot😎\n")
	var operatorId string
	fmt.Scanf("Enter your operator id: %s\n", &operatorId)
	if operatorId[0:6] != "Metro:" {
		fmt.Println("Not Valid Operator id: ", operatorId)
		return
	}

	var level, slots int
	fmt.Scanf("level: %d, slot: %d\n", &level, &slots)
	parkingLot := make([]entities.ParkingSlots, level)
	setParkingLot(slots, parkingLot)
	carSlotMap := make(map[string]entities.SlotData)
	parkHandler := park_car_handler.NewParkCarHandler(parkingLot, carSlotMap)
	exitHandler := exit_car_handler.NewExitCarHandler(parkingLot, carSlotMap)
	displayHandler := display_car_details.NewDisplayCarHandler(carSlotMap)

	fmt.Printf("\n\nValid Operation and Format::\n" +
		"1.parkCar carNo color\n" +
		"2.displayDetails carNo color\n" +
		"3.exitCar carNo color\n" +
		"4.Exit\n",
	)
	for {
		var op, carNo, color string
		var slotData *entities.SlotData
		var err error
		fmt.Scanf("Operation: %s %s %s\n", &op, &carNo, &color)
		carDetails := entities.CarDetails{
			CarNo: carNo,
			Color: color,
		}
		stop := false
		switch {
		case op == "parkCar":
			slotData, err = parkHandler.ParkCar(carDetails)
		case op == "displayDetails":
			slotData, err = displayHandler.DisplayCarDetails(carDetails)
		case op == "exitCar":
			slotData, err = exitHandler.ExitCar(carDetails)
		default:
			stop = true
		}
		if stop {
			break
		}
		if err != nil {
			fmt.Println(fmt.Sprintf("Error while doing op: %v, err: %v", op, err))
		}
		fmt.Println(fmt.Sprintf("OP: %v, Slot Data: %+v", op, slotData))
	}
	fmt.Println("Done for the day, our service will resume in sometime.🙏")
}
*/
