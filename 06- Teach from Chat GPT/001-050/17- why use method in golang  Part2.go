// why use method in golang Part2

// Notion Source: https://sinalalenakhsh.notion.site/8-why-use-method-in-golang-1e880de696b84938836a160ecf599902


package main

import (
	"fmt"
)

type MrRobot struct {
	Price int
	Power int
	Autopilot bool
}

func (robot MrRobot) BuyBread() {
	fmt.Printf("With %d Energy, go and buy bread.\n", robot.Price)
}

func (robot MrRobot) HelpMe() {
	fmt.Printf("with %d help me to clean house.\n", robot.Power)
}

func (robot MrRobot) ChargingTime() {
	fmt.Printf("Autopilot is %t\n", robot.Autopilot)
	fmt.Println("Charging Time ... ")
}

func main() {
	myRobot := MrRobot{
		Price: 100,
		Power: 21,
		Autopilot: false,
	}

	if myRobot.Power < 20 {
		myRobot.Autopilot = true
		myRobot.ChargingTime()
	}

	if myRobot.Price >= 10 && myRobot.Power > 20 {
		myRobot.BuyBread()
		myRobot.Power = myRobot.Power - 20
		fmt.Printf("for buy bread I lose 20 energy and now my energy is %d\n", myRobot.Power)
		if myRobot.Power < 20 {
			myRobot.Autopilot = true
			myRobot.ChargingTime()
		}
	} 
	




}
