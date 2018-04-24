package main

import (  
    "fmt"
    "math"
)

type Employee struct {  
    name     string
    salary   int
    currency string
}


func (e Employee) displaySalary() { // displaySalary() method has Employee as the receiver type
    fmt.Printf("Salary of %s is %s%d\n", e.name, e.currency, e.salary)
}

func dispSalary(e Employee) { // dispSalary() is a function not a method
    fmt.Printf("Salary of %s is %s%d\n", e.name, e.currency, e.salary)
}

func (e Employee) changeName(newName string) { // Method with value receiver  
    e.name = newName
}

func (e *Employee) changeSalary(newSalary int) { // Method with pointer receiver  
    e.salary = newSalary
}

type Rectangle struct {  
    length int
    width  int
}

type Circle struct {  
    radius float64
}

func (r Rectangle) Area() int { // Methods with same name can be defined on different types whereas functions with the same names are not allowed.
    return r.length * r.width
}

func (c Circle) Area() float64 { // Methods with same name can be defined on different types whereas functions with the same names are not allowed.
    return math.Pi * c.radius * c.radius
}

func main() {  
    emp1 := Employee {
        name:     "Sam",
        salary:   5000,
        currency: "$",
    }
    emp1.displaySalary() //Calling displaySalary() method of Employee type
    dispSalary(emp1) // Calling dispSalary() function and passing Employee type variable

    // Changes made inside a method with a pointer receiver is visible to the caller, 
    // whereas this is not the case in value receiver
    fmt.Printf("\nEmployee name before change: %s\n", emp1.name)
    emp1.changeName("Michael Andrew")
    fmt.Printf("Employee name after change: %s\n", emp1.name)

    fmt.Printf("\nEmployee salary before change: %d\n", emp1.salary)
    (&emp1).changeSalary(6000) // Alternately we can write emp1.changeSalary(6000)
    fmt.Printf("Employee salary after change: %d\n", emp1.salary)


    // // Methods with same name can be defined on different types whereas functions with the same names are not allowed.
    r := Rectangle{
        length: 10,
        width:  5,
    }
    fmt.Printf("\nArea of rectangle %d\n", r.Area())
    c := Circle{
        radius: 12,
    }
    fmt.Printf("Area of circle %f\n", c.Area())
}