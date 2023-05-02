package main

import (
	"fmt"
	"math"
)

func SayGreeting(name string) {
	fmt.Printf("Good morning %v \n", name)
}

func SayBye(name string) {
	fmt.Printf("Good bye %v \n", name)
}

func cycleNames(names []string, f func(string)) {
	for _, v := range names {
		f(v)
	}
}

func circleArea(radius float64) float64 {
	return math.Pi * radius * radius
}

func main() {
	SayGreeting("Jane")

	SayBye("Doe")

	cycleNames([]string{"Messi", "Mbappe", "Neymar"}, SayGreeting)
	cycleNames([]string{"Messi", "Mbappe", "Neymar"}, SayBye)

	a1 := circleArea(5)
	fmt.Println(a1)
	fmt.Printf("Area for circle is: %0.4f", a1)
}
