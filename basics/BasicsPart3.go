package mypackage

import "fmt"

func BasicsPart3() {
	var a bool = true       // Boolean
	var b int = 5           // Integer
	var c float32 = 3.14    // Floating point number
	var d string = "Hi GO!" // String

	fmt.Println("Boolean: ", a)
	fmt.Println("Integer: ", b)
	fmt.Println("Float:   ", c)
	fmt.Println("String:  ", d)
	fmt.Println("Boolean Data Type")
	var b1 bool = true // typed declaration with initial value
	var b2 = true      // untyped declaration with initial value
	var b3 bool        // typed declaration without initial value
	b4 := true         // untyped declaration with initial value

	fmt.Println(b1) // Returns true
	fmt.Println(b2) // Returns true
	fmt.Println(b3) // Returns false
	fmt.Println(b4) // Returns true
	fmt.Println("Signed Integers")
	var x int = 500
	var y int = -4500
	fmt.Printf("Type: %T, value: %v", x, x)
	fmt.Printf("\nType: %T, value: %v", y, y)
	fmt.Println("\nfloat32 Keyword")
	var x2 float32 = 123.78
	var y2 float32 = 3.4e+38
	fmt.Printf("Type: %T, value: %v\n", x2, x2)
	fmt.Printf("Type: %T, value: %v", y2, y2)
	fmt.Println("\nfloat64 Keyword")
	var x3 float64 = 1.7e+308
	fmt.Printf("Type: %T, value: %v", x3, x3)
	fmt.Println("\nString Data Type")
	var txt1 string = "Hello!"
	var txt2 string
	txt3 := "World 1"

	fmt.Printf("Type: %T, value: %v\n", txt1, txt1)
	fmt.Printf("Type: %T, value: %v\n", txt2, txt2)
	fmt.Printf("Type: %T, value: %v\n", txt3, txt3)

}
