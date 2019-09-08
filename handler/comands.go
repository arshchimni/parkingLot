package handler

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/parkingLot/pkg"
)

//map of allowed commands along with the arguments to read
var allowedCommands = map[string]int{
	"create_parking_lot": 1,
	"park":               2,
	"leave":              1,
	"status":             0,
	"registration_numbers_for_cars_with_colour": 1,
	"slot_numbers_for_cars_with_colour":         1,
	"slot_number_for_registration_number":       1,
	"exit":                                      0,
}

const (
	UNSUPPORTED_COMMAND           = "Unsupported Command"
	UNSUPPORTED_COMMAND_ARGUMENTS = "Unsupported Command Arguments"
)

// Process the command taken in from file/stdin
// Separate the command and arguments for command
// Validate the command and then do the necessary action
func processCommand(command string) error {
	commandDelimited := strings.Split(command, " ")
	lengthOfCommand := len(commandDelimited)
	arguments := []string{}
	if lengthOfCommand < 1 {
		err := errors.New(UNSUPPORTED_COMMAND)
		fmt.Println(err.Error())
		return err
	} else if lengthOfCommand == 1 {
		command = commandDelimited[0]
	} else {
		command = commandDelimited[0]
		arguments = commandDelimited[1:]
	}

	// check if command is one of the allowed commands
	if numberOfArguments, exists := allowedCommands[command]; exists {

		if len(arguments) != numberOfArguments {
			err := errors.New(UNSUPPORTED_COMMAND_ARGUMENTS)
			fmt.Println(err.Error())
			return err
		}

		// after validation of number of arguments per command, perform the necessary command
		switch command {
		case "create_parking_lot":
			if numberOfSlots, err := strconv.Atoi(arguments[0]); err != nil {
				fmt.Println(err.Error())
				return err
			} else {
				err := pkg.Create_parking_lot(numberOfSlots)
				if err != nil {
					fmt.Println(err.Error())
				}

				return err
			}

		case "park":
			regNo := arguments[0]
			color := arguments[1]
			err := pkg.Park(regNo, color)
			if err != nil {
				fmt.Println(err.Error())
			}
			return err

		case "leave":
			if slot, err := strconv.Atoi(arguments[0]); err != nil {
				fmt.Println(err.Error())
				return err
			} else {
				err := pkg.Leave(slot)
				if err != nil {
					fmt.Println(err.Error())
				}

				return err
			}

		case "status":
			err := pkg.Status()
			if err != nil {
				fmt.Println(err.Error())
			}

			return err

		case "registration_numbers_for_cars_with_colour":
			color := arguments[0]

			err := pkg.Registration_numbers_for_cars_with_colour(color)
			if err != nil {
				fmt.Println(err.Error())
			}
			return err

		case "slot_numbers_for_cars_with_colour":
			color := arguments[0]

			err := pkg.SlotNumberForRegistrationNumber(color)
			if err != nil {
				fmt.Println(err.Error())
			}
			return err

		case "slot_number_for_registration_number":
			regNo := arguments[0]

			err := pkg.SlotNumberForRegistrationNumber(regNo)
			if err != nil {
				fmt.Println(err.Error())
			}
			return err
		case "exit":
			os.Exit(0)

		}
		return errors.New("Not Reachable Code")
	} else {
		err := errors.New(UNSUPPORTED_COMMAND)
		fmt.Println(err.Error())

		return err
	}
}
