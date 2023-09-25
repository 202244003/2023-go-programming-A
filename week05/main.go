package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Input score : ")
	reader := bufio.NewReader(os.Stdin)
	inputNumberString, err := reader.ReadString('\n') // option 2
	if err != nil {
		log.Fatal(err)
	}

	inputNumberString = strings.TrimSpace(inputNumberString)
	inputNumber, err := strconv.ParseFloat(inputNumberString, 32)
	if err != nil {
		log.Fatal(err)
	}
	var grade string
	if inputNumber >= 90 {
		grade = "A grade!"
	} else {
		grade = "Under A grade~"
	}
	fmt.Println("You got " + grade)
}
