package main

import (
	"exclaim"
	"fmt"
	"password"
	"yell"
)

func main() {
	fmt.Println("Sanity Check")
	y := yell.Yell("Sanity Check")
	e := exclaim.Exclaim("Sanity Check")
	h := password.HashPassword("sanity check")
	fmt.Println(y)
	fmt.Println(e)
	fmt.Println(h)
}
