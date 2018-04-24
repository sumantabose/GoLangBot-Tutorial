package main

import "fmt"

func main() {
	var length, bredth int = 15, 10 // Multiple variable declaration
	fmt.Println("Area of rectangle is", length*bredth)

	var ( // Initiating multiple variables inside var
        name   = "Sumanta"
        age    = 28
        height int
    )
    fmt.Println("My name is", name, ",age is", age, "and height is", height)

    flower, number := "lotus", 1 // Short-hand declaration
    fmt.Println("My favourite flower is", flower, "and favourite number is", number)

    a, b := 145.8, 543.8
    c := math.Min(a, b)
    fmt.Println("minimum value is ", c)

}