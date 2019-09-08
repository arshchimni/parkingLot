package pkg

import (
	"container/heap"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/parkingLot/car"
	"github.com/parkingLot/util"
)

//create a parking lot struct that gets initialized with values suppled.

type parkingLot struct {
	Capacity    int
	Constructed bool
	EmptySlots  util.IntHeap
	RegNoSlot   map[string]int
	SlotCar     map[int]car.Car
	ColorRegNo  map[string]map[string]int
}

var instance *parkingLot
var once sync.Once

// Get Instance of Parking lot
func GetInstance() *parkingLot {
	once.Do(func() {
		instance = &parkingLot{}
	})
	return instance
}

func Create_parking_lot(number int) error {
	pl := GetInstance()
	if number <= 0 {
		return fmt.Errorf("cannot create parking lot with given slots")
	}

	for i := 1; i <= number; i++ {
		pl.EmptySlots = append(pl.EmptySlots, i)
	}
	heap.Init(&pl.EmptySlots)
	pl.Capacity = number
	pl.Constructed = true
	pl.RegNoSlot = map[string]int{}
	pl.ColorRegNo = map[string]map[string]int{}
	pl.SlotCar = map[int]car.Car{}

	fmt.Println("Created a parking lot with " + strconv.Itoa(number) + " slots")
	return nil

}

func Park(regNo, color string) error {
	pl := GetInstance()
	if !pl.Constructed {
		return fmt.Errorf("Parking lot not constructed yet cannot park")
	}
	if pl.EmptySlots.Len() == 0 {
		return fmt.Errorf("Sorry, parking lot is full")
	}

	car := car.NewCar(regNo, color)
	allotedSlot := heap.Pop(&pl.EmptySlots)
	pl.SlotCar[allotedSlot.(int)] = car
	pl.RegNoSlot[regNo] = allotedSlot.(int)

	_, exists := pl.ColorRegNo[car.GetColor()]
	if exists {
		pl.ColorRegNo[car.GetColor()][car.GetRegNo()] = allotedSlot.(int)
	} else {
		pl.ColorRegNo[car.GetColor()] = map[string]int{car.GetRegNo(): allotedSlot.(int)}
	}
	fmt.Println("Allocated slot number: " + strconv.Itoa(allotedSlot.(int)))
	return nil
}

func Leave(delSlot int) error {
	pl := GetInstance()
	if !pl.Constructed {

		return fmt.Errorf("Parking lot not constructed yet cannot park")
	}
	if exitCar, ok := pl.SlotCar[delSlot]; ok {
		heap.Push(&pl.EmptySlots, delSlot)
		delete(pl.SlotCar, delSlot)
		delete(pl.RegNoSlot, exitCar.GetRegNo())
		delete(pl.ColorRegNo[exitCar.GetColor()], exitCar.GetRegNo())
		fmt.Println("Slot number " + strconv.Itoa(delSlot) + " is free")

	}
	return nil
}

func Status() error {
	pl := GetInstance()
	if !pl.Constructed {

		return fmt.Errorf("Parking lot not constructed yet cannot park")
	}
	fmt.Println("Slot No.\tRegistration No.\tColour")

	//needed to always have list in a sorted manner
	for i := 1; i <= pl.Capacity; i++ {
		if parkCar, exists := pl.SlotCar[i]; exists {

			fmt.Println(strconv.Itoa(i) + " \t " + strings.ToUpper(parkCar.GetRegNo()) + " \t " + (parkCar.GetColor()))
		}
	}
	return nil
}

func Registration_numbers_for_cars_with_colour(color string) error {
	pl := GetInstance()
	if !pl.Constructed {

		return fmt.Errorf("Parking lot not constructed yet cannot park")
	}

	var prettyRegNo []string

	if regMap, ok := pl.ColorRegNo[color]; ok {
		for regNo := range regMap {
			prettyRegNo = append(prettyRegNo, strings.ToUpper(regNo))
		}
		sort.Strings(prettyRegNo)
		fmt.Println(strings.Join(prettyRegNo, ", "))
	} else {
		fmt.Println("Car with that color not present in the parking lot")
	}
	return nil
}

func Slot_numbers_for_cars_with_colour(color string) error {
	pl := GetInstance()
	if !pl.Constructed {

		return fmt.Errorf("Parking lot not constructed yet cannot park")
	}

	var prettySlot []string

	if regMap, ok := pl.ColorRegNo[color]; ok {
		for _, slot := range regMap {
			prettySlot = append(prettySlot, strconv.Itoa(slot))
		}
		sort.Strings(prettySlot)
		fmt.Println(strings.Join(prettySlot, ", "))
	} else {
		fmt.Println("Car with that color not present in the parking lot")
	}
	return nil
}

func Slot_number_for_registration_number(regNo string) error {
	pl := GetInstance()
	if !pl.Constructed {

		return fmt.Errorf("Parking lot not constructed yet cannot park")
	}

	if slot, ok := pl.RegNoSlot[regNo]; ok {
		fmt.Println(slot)
	} else {
		fmt.Println("Not found")
	}
	return nil
}
