/*
// defer function
package main

import (  
    "fmt"
)

func finished() {  
    fmt.Println("Finished finding largest")
}

func largest(nums []int) {  
    defer finished()    
    fmt.Println("Started finding largest")
    max := nums[0]
    for _, v := range nums {
        if v > max {
            max = v
        }
    }
    fmt.Println("Largest number in", nums, "is", max)
}

func main() {  
    nums := []int{78, 109, 2, 563, 300}
    largest(nums)
}
*/


/*
// defer method
package main

import (  
    "fmt"
)


type person struct {  
    firstName string
    lastName string
}

func (p person) fullName() {  
    fmt.Printf("%s %s\n",p.firstName,p.lastName)
}

func main() {  
    p := person {
        firstName: "Sumanta",
        lastName: "Bose",
    }
    defer p.fullName()
    fmt.Printf("Welcome ")  
}
*/





//The arguments of a deferred function are evaluated when the defer statement is executed 
// and not when the actual function call is done.
package main

import "fmt"

func printA(a int) {  
    fmt.Println("value of a in deferred function", a)
}
func main() {  

    // When a function has multiple defer calls, 
    // they are added on to a stack and executed in Last In First Out (LIFO) order.
    a := 5
    defer printA(a) // added to stack (a = 5)
    a = 10
    fmt.Println("value of a before deferred function call", a)
    defer printA(a) // added to stack on top of previous defer, hence executed first (a = 10)
}

