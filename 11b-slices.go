package main

import (  
    "fmt"
)

func main() {  
    a := [7]int{71, 74, 78, 79, 80, 85, 92}
    var b []int = a[1:3] //creates a slice from a[1] to a[2]
    // The syntax a[start:end] creates a slice from array  a starting from index (start) to index (end - 1)
    fmt.Println(b)


    darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
    dslice := darr[2:5]
    fmt.Println("array before:",darr)
    for i := range dslice {
        dslice[i]++
        // A slice does not own any data of its own. It is just a representation of the underlying array.
        // Any modifications done to the slice will be reflected in the underlying array.
    }
    fmt.Println("array after:",darr)


    numa := [3]int{78, 79 ,80}
    nums1 := numa[:] //creates a slice which contains all elements of the array
    nums2 := numa[:]
    // When a number of slices share the same underlying array, the changes that each one makes will be reflected in the array.
    fmt.Println("array before change 1",numa)
    nums1[0] = 100
    fmt.Println("array after modification to slice nums1", numa)
    nums2[1] = 101
    fmt.Println("array after modification to slice nums2", numa)


    fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
    fruitslice := fruitarray[1:3]
    fmt.Printf("length of slice %d capacity %d\n", len(fruitslice), cap(fruitslice)) //length of is 2 and capacity is 6
    // The length of the slice is the number of elements in the slice.
    // The capacity of the slice is the number of elements in the underlying array starting from the index from which the slice is created.
    fruitslice = fruitslice[:cap(fruitslice)] //re-slicing furitslice till its capacity
    fmt.Println("After re-slicing length is",len(fruitslice), "and capacity is",cap(fruitslice))


    cars := []string{"Ferrari", "Honda", "Ford"}
    fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) //capacity of cars is 3
    cars = append(cars, "Toyota")
    fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 6

    veggies := []string{"potatoes","tomatoes","brinjal"}
    fruits := []string{"oranges","apples"}
    food := append(veggies, fruits...) // It is also possible to append one slice to another using the ... operator. 
    fmt.Println("food:",food)

    // xx := [...]int{} // Here xx would be an array similar to line 14 & 36
    xx := []int{} // Here xx is a slice
    fmt.Println(xx)
    xx = append(xx,1)
    fmt.Println(xx)
    xx = append(xx,2)
    fmt.Println(xx)
}