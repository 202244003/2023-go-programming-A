package main

import (
	"fmt"
	"strings"
)

func main() {
	Hotspurs := "hm ? j madi?"
	replacePlayer := strings.NewReplacer("?", "son")
	player := replacePlayer.Replace(Hotspurs)
	fmt.Println(player)
}
