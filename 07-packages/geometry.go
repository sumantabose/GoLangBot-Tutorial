//geometry.go
package main

import (
	"./rectangle" //importing user defined package from directory ./rectangle/
	"fmt"
)

func init() { // this initialise() function runs one time in the very beginning
	println("main package initialized")
}

func main() {
	var rectLen, rectWidth float64 = 6.5, 7.5
	fmt.Println("Geometrical shape properties")
	/*Area function of rectangle package used
	 */
	fmt.Printf("area of rectangle %.2f\n", rectangle.Area(rectLen, rectWidth))
	/*Diagonal function of rectangle package used
	 */
	fmt.Printf("diagonal of the rectangle %.2f\n", rectangle.Diagonal(rectLen, rectWidth))
}
