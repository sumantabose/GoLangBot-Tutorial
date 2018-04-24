package main

import (  
    "fmt"
)

func changeLocal(num [5]int) {  
    num[0] = 55
    fmt.Println("inside function ", num)

}

func printarray(a [3][2]string) {  
    for _, v1 := range a {
        for _, v2 := range v1 {
            fmt.Printf("%s ", v2)
        }
        fmt.Printf("\n")
    }
}

func main() {  
    var a [3]int //int array with length 3
    fmt.Println(a)
    fmt.Println("length of a is",len(a))

    a[0] = 12 // array index starts at 0
    a[1] = 78
    a[2] = 50
    fmt.Println(a)

    b := [3]int{52, 18, 70} // short hand declaration to create array
    fmt.Println(b)
    fmt.Println("length of b is",len(b))

    c := [...]int{22, 48, 10} // ... makes the compiler determine the length
    fmt.Println(c)

    x := [...]string{"USA", "China", "India", "Germany", "France"}
    y := x // a copy of a is assigned to b
    y[0] = "Singapore"
    fmt.Println("x is ", x)
    fmt.Println("y is ", y)
    fmt.Println("length of x is",len(x))
    fmt.Println("length of y is",len(y))

    num := [...]int{5, 6, 7, 8, 8}
    fmt.Println("before passing to function ", num)
    changeLocal(num) //num is passed by value
    fmt.Println("after passing to function ", num)
    for i := 0; i < len(num); i++ { //looping from 0 to the length of the array
        fmt.Printf("Element %d of num is %d\n", i, num[i])
    }
    sum := int(0)
    for j, v := range num { // range returns both the index and value
        fmt.Printf("Element %d of num is %d\n", j, v)
        sum += v
    }
    fmt.Println("\nsum of all elements of num is",sum)
    fmt.Println("")



    ax := [3][2]string{
        {"lion", "tiger"},
        {"cat", "dog"},
        {"pigeon", "peacock"}, //this comma is necessary. The compiler will complain if you omit this comma
    }
    printarray(ax)
    var bx [3][2]string
    bx[0][0] = "apple"
    bx[0][1] = "samsung"
    bx[1][0] = "microsoft"
    bx[1][1] = "google"
    bx[2][0] = "AT&T"
    bx[2][1] = "T-Mobile"
    fmt.Printf("\n")
    printarray(bx)
}