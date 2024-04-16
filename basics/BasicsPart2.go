package mypackage

import "fmt"

const PI = 3.4

// typed constant
const A1 int = 1

// multiple constant declaration
const (
	A int = 1
	B     = 3.14
	C     = "Hi Go!"
)

func BasicsPart2() {
	fmt.Println(PI)
	fmt.Println("typed constant")
	fmt.Println(A1)
	fmt.Println("multiple constant declaration")
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println("Print example")
	var i, j = 10, 20
	fmt.Print(i, j)
	fmt.Println("\nPrint example")
	var i1, j1 string = "Hello", "GO"

	fmt.Print(i1, " ", j1)
	fmt.Print("\n", i1, "\n", j1)
	fmt.Println("\nPrintf example")
	var i2 string = "Hello"
	var j2 int = 15

	fmt.Printf("i2 has value: %v and type: %T\n", i2, i2)
	fmt.Printf("j2 has value: %v and type: %T", j2, j2)

	fmt.Println("\nGeneral Formatting Verbs")
	var i3 = 15.5
	var txt = "Hello World!"

	fmt.Printf("%v\n", i3)
	fmt.Printf("%#v\n", i3)
	fmt.Printf("%v%%\n", i3)
	fmt.Printf("%T\n", i3)

	fmt.Printf("%v\n", txt)
	fmt.Printf("%#v\n", txt)
	fmt.Printf("%T\n", txt)

	fmt.Println("\nInteger Formatting Verbs")
	var i4 = 15

	fmt.Printf("%b\n", i4)
	fmt.Printf("%d\n", i4)
	fmt.Printf("%+d\n", i4)
	fmt.Printf("%o\n", i4)
	fmt.Printf("%O\n", i4)
	fmt.Printf("%x\n", i4)
	fmt.Printf("%X\n", i4)
	fmt.Printf("%#x\n", i4)
	fmt.Printf("%4d\n", i4)
	fmt.Printf("%-4d\n", i4)
	fmt.Printf("%04d\n", i4)

	fmt.Println("\nString Formatting Verbs")
	var txt1 = "Hello"

	fmt.Printf("%s\n", txt1)
	fmt.Printf("%q\n", txt1)
	fmt.Printf("%8s\n", txt1)
	fmt.Printf("%-8s\n", txt1)
	fmt.Printf("%x\n", txt1)
	fmt.Printf("% x\n", txt1)

	fmt.Println("\nBoolean Formatting Verbs")
	var i5 = true
	var j5 = false

	fmt.Printf("%t\n", i5)
	fmt.Printf("%t\n", j5)

	fmt.Println("\nFloat Formatting Verbs")
	var i6 = 3.141

	fmt.Printf("%e\n", i6)
	fmt.Printf("%f\n", i6)
	fmt.Printf("%.2f\n", i6)
	fmt.Printf("%6.2f\n", i6)
	fmt.Printf("%g\n", i6)
}
