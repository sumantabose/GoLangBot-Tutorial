package main

import "fmt"

func main() {  
        /*
        var defaultName = "Sam" //allowed
        type myString string
        var customName myString = "Sam" //allowed
        customName = defaultName //not allowed
		*/

        const a = 5
	    var intVar int = a
	    var int32Var int32 = a
	    var float64Var float64 = a
	    var complex64Var complex64 = a
	    fmt.Println("intVar",intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var",complex64Var)

	    var i = 5
	    var f = 5.6
    	var c = 5 + 6i
    	fmt.Printf("i's type %T, f's type %T, c's type %T\n", i, f, c)


}