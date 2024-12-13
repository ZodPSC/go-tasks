package main

import (
	"fmt"
	"sort"
)

func main() {

	//Finding the middle element(-s)
	fmt.Println("Finding the middle element(-s)")
	var countPlanet int
	fmt.Print("Enter the number of planets: ")
	fmt.Scan(&countPlanet)

	var countHumanOnPlanet = make([]int, countPlanet)

	for i := 0; i < countPlanet; i++ {
		fmt.Printf("Input human on planet #%d: ", i+1)
		fmt.Scan(&countHumanOnPlanet[i])
	}

	sort.Ints(countHumanOnPlanet)

	for _, count := range countHumanOnPlanet {
		fmt.Printf("%d ", count)
	}
	fmt.Println()

	indexMiddleElement := countPlanet / 2

	//fmt.Println(indexMiddleElement)

	if indexMiddleElement%2 == 0 {
		fmt.Printf("The middle elements is: %d and %d", countHumanOnPlanet[indexMiddleElement-1], countHumanOnPlanet[indexMiddleElement])
	} else {
		fmt.Printf("The middle element is: %d", countHumanOnPlanet[indexMiddleElement])
	}
}
