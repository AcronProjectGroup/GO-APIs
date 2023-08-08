// Source : https://linuxhint.com/golang-map-of-maps/

package main

import (
	"fmt"
)

func main() {
	myMap := map[string]map[string]string{
		"UserInfo": {
			"Name":     "Sina",
			"SureName": "Lalehbakhsh",
			"Age":      "29",
		},
		"Scores": {
			"Python": "100",
			"Linux":  "100",
			"Golang": "100",
		},
		"Targets": {
			"First Target":  "imegrate To Canada",
			"Second Target": "Get Job in Google",
			"Third Target":  "Chinese Food",
		},
	}

	fmt.Println("User info:", myMap["UserInfo"])
	fmt.Println("User scores:", myMap["Scores"])
	fmt.Println("User targets:", myMap["Targets"])


}
