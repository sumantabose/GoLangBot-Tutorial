//rectCalculate.go
package rectangle // package name must be the same name as the directory name

import "math"
import "fmt"

/*
 * init function added
 */
func init() { // this initialise() function runs one time in the very beginning
	fmt.Println("rectangle package initialized")
}
func Area(len, wid float64) float64 {
	area := len * wid
	return area
}

func Diagonal(len, wid float64) float64 {
	diagonal := math.Sqrt((len * len) + (wid * wid))
	return diagonal
}
