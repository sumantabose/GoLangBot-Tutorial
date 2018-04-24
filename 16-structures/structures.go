package main

import (  
    "fmt"
    "./computer" 
)

type Employee struct {  
    firstName, lastName string
    age, salary         int
}

type Address struct {  
    city, state string
}

type Person struct {  
    name string
    age int
    address Address
}

func main() {

    //creating structure using field names
    emp1 := Employee{
        firstName: "Sam",
        age:       25,
        salary:    500,
        lastName:  "Anderson",
    }

    //creating structure without using field names
    emp2 := Employee{"Thomas", "Paul", 29, 6000}

    fmt.Println("Employee 1", emp1)
    fmt.Println("Employee 2", emp2)

    emp3 := &Employee{"Bill", "Gates", 55, 9000} // Pointers to a struct
    fmt.Println("Employee 3", *emp3)

    /////////////

    var p Person
    p.name = "Sumanta"
    p.age = 27
    p.address = Address {
        city: "NTU",
        state: "Singapore",
    }
    fmt.Println("Name:", p.name)
    fmt.Println("Age:",p.age)
    fmt.Println("City:",p.address.city)
    fmt.Println("State:",p.address.state)

    //////////////

    var spec computer.Spec
    spec.Maker = "Apple"
    // spec.model = "iMac" // will not work since model starts with a lower case
    spec.Price = 50000
    fmt.Println("Spec:", spec)
}