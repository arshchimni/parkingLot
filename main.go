package main

import "github.com/parkingLot/pkg"

func main() {
	pkg.Create_parking_lot(10)
	pkg.Park("KA-04-JQ-4731", "grey")
	pkg.Park("KA-04-JQ-4732", "grey")
	pkg.Park("KA-04-JQ-4731", "pink")
	//pkg.Leave(1)
	pkg.Status()
	pkg.Registration_numbers_for_cars_with_colour("grey")
	pkg.Slot_numbers_for_cars_with_colour("pink")
	pkg.Slot_number_for_registration_number("KA-04-JQ-4732")
}
