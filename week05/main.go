package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("Input score : ")
	reader := bufio.NewReader(os.Stdin)
	inputNumber, err := reader.ReadString('\n') // option 2
	if err != nil {
		log.Fatal(err)
	}
	if inputNumber >= 90 {
		grade := "A grade!"
	} else {
		grade := "BCDE grade~"
	}
	fmt.Println(grade)
}
