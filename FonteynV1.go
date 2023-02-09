package main

import (
	"fmt"
	"time"
)

func main() {
	greet := realTime()
	greeting(greet)
}

// geeft de realistische tijd weer in uren
func realTime() int {
	currentTime := time.Now()
	groet := currentTime.Hour()
	return groet
}

func greeting(groet int) {

	if groet >= 7 && groet <= 11 {
		fmt.Println("Goedenmorgen")
	} else if groet >= 12 && groet <= 18 {
		fmt.Println("Goedenmiddag")
	} else if groet >= 18 && groet <= 23 {
		fmt.Println("Goedenavond!")
	} else {
		fmt.Println("Sorry, de parkeerplaats is ’s nachts gesloten")
	}
}
