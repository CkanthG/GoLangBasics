package mypackage

import "fmt"

func BasicsPart1() {
	var number = 100
	var userName string = "Sreekanth"
	x := 10
	fmt.Println("Hello Go World!")
	fmt.Println("Number : ", number)
	fmt.Println("UserName : ", userName)
	fmt.Println("X : ", x)
	// declare without initial value
	var a string
	var b int
	var c bool
	fmt.Println("declare without initial value")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	// multiple variable declarations
	var a2, b2 = 6, "Hello"
	c2, d2 := 7, "Go!"
	fmt.Println("multiple variable declarations")
	fmt.Println(a2)
	fmt.Println(b2)
	fmt.Println(c2)
	fmt.Println(d2)
	// block declaration
	var (
		a1 int
		b1 int    = 1
		c1 string = "hello"
	)
	fmt.Println("block declaration")
	fmt.Println(a1)
	fmt.Println(b1)
	fmt.Println(c1)

}
