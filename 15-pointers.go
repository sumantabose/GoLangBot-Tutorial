package main

import (  
    "fmt"
)

func squareVal(x *int) *int {  
    fmt.Println("square Value:", (*x)*(*x))
    *x = 3
    return x
}


func main() {  
	a := 2
    b := &a
    fmt.Println(a, &a)
    fmt.Println(b, *b)

    c := squareVal(b)

    fmt.Println(a, &a)
    fmt.Println(b, *b)
    fmt.Println(c, *c)
}