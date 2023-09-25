package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Input score : ")
	reader := bufio.NewReader(os.Stdin)
	inputNumber, _ := reader.ReadString('\n') // option 1, Ignore The error return value with blank
	fmt.Println(inputNumber)
}
