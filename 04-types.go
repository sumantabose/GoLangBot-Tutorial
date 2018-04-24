package main

import (  
    "fmt"
    "unsafe"
)

func main() {  
    var a int = 89
    b := 95
    fmt.Println("value of a is", a, "and b is", b)
    fmt.Printf("type of a is %T, size of a is %d bytes\n", a, unsafe.Sizeof(a)) //type and size of a
    fmt.Printf("type of b is %T, size of b is %d bytes\n", b, unsafe.Sizeof(b)) //type and size of b
    // %T is the format specifier to print the type and %d is used to print the size.


    aa := true // Variable type is Bool
    bb := false
    fmt.Println("aa:", aa, "bb:", bb)
    cc := aa && bb // Logical function
    fmt.Println("cc:", cc)
    dd := aa || bb // Logical function
    fmt.Println("dd:", dd)

    c1 := complex(5, 7) // Variable type in complex number
    c2 := 8 + 27i
    cadd := c1 + c2
    fmt.Println("sum:", cadd)
    cmul := c1 * c2
    fmt.Println("product:", cmul)

    first := "Sumanta"
    last := "Bose"
    name := first +" "+ last
    fmt.Println("My name is",name)

    i := 55      //int
    j := 67.8    //float64
    sum1 := i + int(j) //j is converted to int
    sum2 := float64(i) + j
    fmt.Printf("sum1 is %d and sum2 is %f\n", sum1, sum2)
}