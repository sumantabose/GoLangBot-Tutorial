package main

import (  
    "fmt"
)

func main() {  
    for i := 1; i <= 10; i++ {
        fmt.Printf("%d ",i)
    }
    fmt.Printf("\n")


    for j := 1; j <= 10; j++ {
        if j > 5 {
            break //loop is terminated if j > 5
        }
        fmt.Printf("%d ", j)
    }
    fmt.Printf("\nline after for loop\n")

    for k := 1; k <= 10; k++ {
        if k%2 == 0 {
            continue
            // The continue statement is used to skip the current iteration of the for loop.
            // All code present in a for loop after the continue statement will not be executed for the current iteration.
            // The loop will move on to the next iteration.
        }
        fmt.Printf("%d ", k)
    }
    fmt.Printf("\n")


    ii := 0
    for ;ii <= 10; { // initialisation and post are omitted
        fmt.Printf("%d ", ii)
        ii += 2
    }

    for no, m := 10, 1; m <= 10 && no <= 19; m, no = m+1, no+1 { //multiple initialisation and increment
        fmt.Printf("%d * %d = %d\n", no, m, no*m)
    }

    for { // Infinite loop
        fmt.Println("Hello World")
    }
}