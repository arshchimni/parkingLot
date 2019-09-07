package pkg

import (
	"container/heap"
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"

	"github.com/parkingLot/car"
	"github.com/parkingLot/util"
)

/*
Commands needed to be made
done 1.create_parking_lot
done 2.park
done 3.leave
done 4.status
done 5.registration_numbers_for_cars_with_colour
done 6.slot_numbers_for_cars_with_colour
7.slot_number_for_registration_number
*/

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

func Create_parking_lot(number int) {
	pl := GetInstance()
	if number <= 0 {
		log.Println("cannot create parking lot with given slots")
		return
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
	fmt.Println(pl)

	fmt.Println("Created a parking lot with " + strconv.Itoa(number) + " slots")

}

func Park(regNo, color string) {
	pl := GetInstance()
	if !pl.Constructed {
		log.Println("Parking lot not constructed yet cannot park")
		return
	}
	if pl.EmptySlots.Len() == 0 {
		log.Println("Parking lot has no empty slots")
		return
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
	log.Println("Allocated slot number: " + strconv.Itoa(allotedSlot.(int)))
}

func Leave(delSlot int) {
	pl := GetInstance()
	if !pl.Constructed {
		log.Println("Parking lot not constructed yet cannot park")
		return
	}
	if exitCar, ok := pl.SlotCar[delSlot]; ok {
		heap.Push(&pl.EmptySlots, delSlot)
		delete(pl.SlotCar, delSlot)
		delete(pl.ColorRegNo[exitCar.GetColor()], exitCar.GetRegNo())
		log.Println("Slot number " + strconv.Itoa(delSlot) + " is free")

	}
	log.Println(pl)
}

func Status() {
	pl := GetInstance()
	if !pl.Constructed {
		log.Println("Parking lot not constructed yet cannot park")
		return
	}
	fmt.Println("Slot No.\tRegistration No.\tColour")
	for slot, parkCar := range pl.SlotCar {

		fmt.Println(strconv.Itoa(slot) + " \t " + strings.ToUpper(parkCar.GetRegNo()) + " \t " + (parkCar.GetColor()))
	}
}

func Registration_numbers_for_cars_with_colour(color string) {
	pl := GetInstance()
	if !pl.Constructed {
		log.Println("Parking lot not constructed yet cannot park")
		return
	}

	var prettyRegNo []string

	if regMap, ok := pl.ColorRegNo[color]; ok {
		for regNo := range regMap {
			prettyRegNo = append(prettyRegNo, regNo)
		}
		fmt.Println(strings.Join(prettyRegNo, ", "))
	} else {
		log.Println("Car with that color not present in the parking lot")
	}

}

func Slot_numbers_for_cars_with_colour(color string) {
	pl := GetInstance()
	if !pl.Constructed {
		log.Println("Parking lot not constructed yet cannot park")
		return
	}

	var prettySlot []string

	if regMap, ok := pl.ColorRegNo[color]; ok {
		for _, slot := range regMap {
			prettySlot = append(prettySlot, strconv.Itoa(slot))
		}
		fmt.Println(strings.Join(prettySlot, ", "))
	} else {
		log.Println("Car with that color not present in the parking lot")
	}
}

func Slot_number_for_registration_number(regNo string) {
	pl := GetInstance()
	if !pl.Constructed {
		log.Println("Parking lot not constructed yet cannot park")
		return
	}

	if slot, ok := pl.RegNoSlot[regNo]; ok {
		log.Println(slot)
	} else {
		log.Println("Not found")
	}

}
