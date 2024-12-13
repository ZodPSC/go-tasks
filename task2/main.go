package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {

	//symbol
	charset := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	countCharset := len(charset)

	//const countBase = 10000
	const countBase = 100

	//getting random character
	var genomeOne string
	for i := 0; i < countBase; i++ {
		genomeOne += string(charset[rand.Intn(countCharset)])
	}
	fmt.Printf("%s\n", genomeOne)

	var genomeTwo string
	for i := 0; i < countBase; i++ {
		genomeTwo += string(charset[rand.Intn(countCharset)])
	}
	fmt.Printf("%s\n", genomeTwo)

	fmt.Println("Generated genomes is completed")

	coincidence := 0

	for i := 0; i < countBase-1; i++ {
		currentPair := string(genomeOne[i]) + string(genomeOne[i+1])
		isContains := strings.Contains(genomeTwo, currentPair)
		if isContains {
			coincidence++
		}
		//fmt.Printf("%s is %t \n", currentPair, isContains)
	}
	fmt.Println()
	fmt.Printf("coincidence = %d\n", coincidence)
}
