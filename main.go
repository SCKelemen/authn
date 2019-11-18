package main

import (
	"exclaim"
	"fmt"
	"yell"
)

func main() {
	fmt.Println("Sanity Check")
	y := yell.Yell("Sanity Check")
	e := exclaim.Exclaim("Sanity Check")
	fmt.Println(y)
	fmt.Println(e)
}
