package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Input score : ")
	reader := bufio.NewReader(os.Stdin)
	inputNumber := reader.ReadString('\n') // 1 variable but reader.ReadString returns 2 values
	fmt.Println(inputNumber)
}
