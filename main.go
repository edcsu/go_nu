package main

import "fmt"

func main() {
	// arrays have fixed length
	var ages [3]int = [3]int{90, 60, 72}
	var ages2 = [3]int{65, 18, 120}
	ages3 := [4]int{80, 40, 52}

	fmt.Println(ages, len(ages))
	fmt.Println(ages2, len(ages2))
	fmt.Println(ages2, len(ages2))
	fmt.Println(ages3, len(ages3))

	// slices (arrays behind the scenes)
	var names = []string{"James", "Peter", "Ole"}

	fmt.Println(names, len(names))

	names = append(names, "Rashford")
	fmt.Println(names, len(names))

	// slice ranges

	range1 := names[1:3]
	range2 := names[2:]
	range3 := names[:3]

	fmt.Println(range1)
	fmt.Println(range2)
	fmt.Println(range3)
}
