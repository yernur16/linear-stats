package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	arg := os.Args[1:]

	if len(arg) != 1 {
		log.Println("error, there must be only (1)one argument")
		return
	}
	dataByte, err := ioutil.ReadFile(arg[0])
	if err != nil {
		log.Println("error, file does not exist")
		return
	}
	var numbers []int

	strFromBytes := strings.Fields(string(dataByte))
	for _, j := range strFromBytes {
		conversion, err := strconv.Atoi(j)
		if err != nil {
			log.Println("error, in the appending data into array")
			return
		}
		numbers = append(numbers, conversion)
	}

	linreg(numbers)
}

// sum of 12502 lines == 78 156 253

func linreg(arr []int) {
	sumX := 0.0
	sumX2 := 0.0
	sumY := 0.0
	sumY2 := 0.0

	xy := 0.0

	for i := 0; i < len(arr); i++ {
		sumX += float64(i)
		sumX2 += float64((i * i))
		sumY += float64(arr[i])
		sumY2 += float64(arr[i] * arr[i])
		xy += float64(i * arr[i])

	}
	xSquare := sumX * sumX
	ySquare := sumY * sumY

	l := float64(len(arr))
	// Linear regression line formula
	a := (sumY*sumX2 - sumX*xy) / (l*sumX2 - xSquare)
	b := ((l * xy) - (sumX * sumY)) / (l*sumX2 - xSquare)

	// Pearson Correlation Coefficient formula
	r := ((l * xy) - sumX*sumY) / math.Sqrt((l*sumX2-xSquare)*(l*sumY2-ySquare))
	fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\n", b, a)
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", r)
}
