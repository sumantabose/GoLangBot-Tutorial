package main

import (  
    "fmt"
)

func calculateBill(price, no int) int {  // One return value
    var totalPrice = price * no
    return totalPrice
}

func rectProps1(length, width float64)(float64, float64) {  // Multiple return values
    var area = length * width
    var perimeter = (length + width) * 2
    return area, perimeter
}


func rectProps2(length, width float64)(area, perimeter float64) {  // Named return values
    area = length * width
    perimeter = (length + width) * 2
    return //no explicit return value
}



func main() {  
    price, no := 90, 6
    totalPrice := calculateBill(price, no)
    fmt.Println("Total price is", totalPrice)

    area, perimeter := rectProps1(10.8, 5.6)
    fmt.Printf("Area %f Perimeter %f\n", area, perimeter) 

    area, perimeter = rectProps2(11.8, 4.6)
    fmt.Printf("Area %f Perimeter %f\n", area, perimeter)

    ar,_ := rectProps2(11.8, 4.6) // Perimeter is discarded by the use of underscore
    fmt.Printf("Area %f\n", ar)
}