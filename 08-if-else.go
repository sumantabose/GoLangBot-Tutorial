package main

import (  
    "fmt"
)

func main() {  
    num := 10
    if num % 2 == 0 { //checks if number is even
        fmt.Println("the number is even") 
    }  else { // The else statement should start in the same line after the closing curly brace } of the if statement. If not the compiler will complain.
        fmt.Println("the number is odd")
    }


    num = 99
    if num >= 0 && num <= 50 {
        fmt.Println("number is less than or equal to 50")
    } else if num >= 51 && num <= 100 {
        fmt.Println("number is between 51 and 100")
    } else { // The else statement should start in the same line after the closing curly brace } of the if statement. If not the compiler will complain.
        fmt.Println("number is greater than 100")
    }


    if newnum := 20; newnum % 2 == 0 { //checks if newnum is even // scope of newnum is limited to the if-else code block only
        fmt.Println(newnum,"is even") 
    }  else { // The else statement should start in the same line after the closing curly brace } of the if statement. If not the compiler will complain.
        fmt.Println(newnum,"is odd")
    }

    // fmt.Println(newnum) // will throw an error
}