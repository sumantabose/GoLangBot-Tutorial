package main

import (  
    "fmt"
)

func number() int {  
        num := 15 * 5 
        return num
}

func main() {  
    switch finger := 4; finger { // variable finger is declared and initialised in the switch statement itself
        // The scope of finger is limited to the switch code block
    case 1:
        fmt.Println("Thumb")
    case 2:
        fmt.Println("Index")
    case 3:
        fmt.Println("Middle")
    case 4:
        fmt.Println("Ring")
    case 5:
        fmt.Println("Pinky")
    default: //default case
        fmt.Println("Incorrect finger number")
    }

    letter := "i"
    switch letter {
    case "a", "e", "i", "o", "u": //multiple expressions in case
        fmt.Println("vowel")
    default:
        fmt.Println("consonant")
    }

    num := 75
    switch { // expression is omitted
    // The expression in a switch is optional and it can be omitted.
        // If the expression is omitted, the switch is considered to be 'switch true' and each of the case expression is evaluated for truth and the corresponding block of code is executed.
    case num >= 0 && num <= 50:
        fmt.Println("num is greater than 0 and less than 50")
    case num >= 51 && num <= 100:
        fmt.Println("num is greater than 51 and less than 100")
    case num >= 101:
        fmt.Println("num is greater than 100")
    }

    switch num2 := number(); { //num2 is not a constant. It is a return value from a function
    case num2 < 50:
        fmt.Printf("%d is lesser than 50\n", num2)
        fallthrough // fallthrough should be the last statement in a case. If it present somewhere in the middle, the compiler will throw error fallthrough statement out of place
    case num2 < 100:
        fmt.Printf("%d is lesser than 100\n", num2)
        fallthrough
    case num2 < 200:
        fmt.Printf("%d is lesser than 200\n", num2)
    }
}