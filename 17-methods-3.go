package main

import (  
    "fmt"
)

type rectangle struct {  
    length int
    width  int
}

func perimeter(r *rectangle) {  
    fmt.Println("perimeter function output:", 2*(r.length+r.width))

}

func (r *rectangle) perimeter() {  
    fmt.Println("perimeter method output:", 2*(r.length+r.width))
}

func main() {  
    r := rectangle{
        length: 10,
        width:  5,
    }
    p := &r //pointer to r
    perimeter(p)
    p.perimeter()

    /*
        cannot use r (type rectangle) as type *rectangle in argument to perimeter
    */
    //perimeter(r)

    r.perimeter()//calling pointer receiver with a value

}

/*

Pointer receivers in methods vs pointer arguments in functions.

Similar to value arguments, functions with pointer arguments will accept 
only pointers whereas methods with pointer receivers will accept both value and pointer receiver.

Line no. 12 of the above program defines a function perimeter which accepts a pointer argument 
and line no. 17 defines a method which has a pointer receiver.

In line no. 27 we call the perimeter function with a pointer argument and in line no.28 
we call the perimeter method on a pointer receiver. All is good.

In the commented line no.33, we try to call the perimeter function with a value argument r. 
This is not allowed since a function with pointer argument will not accept value arguments. 
If this line is uncommented and the program is run, the compilation will fail 
with error main.go:33: cannot use r (type rectangle) as type *rectangle in argument to perimeter.

In line no.35 we call the pointer receiver method perimeter with a value receiver r. 
This is allowed and the line of code 
r.perimeter() will be interpreted by the language as (&r).perimeter() for convenience. 

        This program will output,

        perimeter function output: 30  
        perimeter method output: 30  
        perimeter method output: 30  

*/