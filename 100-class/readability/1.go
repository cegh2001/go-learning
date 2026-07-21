package main

import "fmt"

func main() {

	if batteryLevel := 10; batteryLevel < 15 {
		fmt.Println("Please plug in your charger")
	} else if batteryLevel < 20 {
		fmt.Println("Battery level is low, consider charging soon")
	} else {
		fmt.Println("Battery level is sufficient")
	}
}