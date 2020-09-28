// Source: https://medium.com/rungo/string-formatting-in-go-dd752ff7f60
package main

import "fmt"

func main() {
	standardPrint()
	variadicFunction()
	printlnFunction()
	SprintReturnCase()
	stringInterpolation()
	vCaseInterpolation()
	vHastashInterpolation()
	tTypeFormat()
	tBooleanFormat()
	dBase10Format()
	bBinaryFormat()
	xhexadecimalformat()
	fFuntionFormat()
}

func standardPrint()  {
	// print to standard output
	fmt.Print("Hello World!")
}

func variadicFunction()  {
	fmt.Print("Hello World!")
	fmt.Print("How are you doing?")

	// start and end with a new line
	fmt.Print("\nI am", 26, "years old\n")

	// print non-string values
	fmt.Print(1, 2, false, 3, 4, nil, true)
}

func printlnFunction()  {
	fmt.Println("Hello World!")
	fmt.Println("How are you doing?")

	// start and end with a new line
	fmt.Println("I am", 26, "years old")

	// print non-string values
	fmt.Println(1, 2, false, 3, 4, nil, true)
}

func SprintReturnCase()  {
	s1 := fmt.Sprint("Hello World!")
	s2 := fmt.Sprint("How are you doing?")

	// start and end with a new line
	s3 := fmt.Sprint("\nI am", 26, "years old\n")

	// non-string values
	s4 := fmt.Sprint(1, 2, false, 3, 4, nil, true)

	// print
	fmt.Print(s1, s2, s3, s4)
}

func stringInterpolation()  {
	// write to std ouput
	fmt.Printf("Hi there, I am %v and I'm %v years old", "John", 26)
}

func vCaseInterpolation()  {
	// struct
	s := struct {
		name string
		age  int
	}{"John", 26}

	// channel
	c := make(chan int)

	// map
	m := map[string]int{
		"one": 1,
		"two": 2,
	}

	fmt.Printf("%v \n", true)           // boolean values
	fmt.Printf("%v \n", 132)            // signed and unsigned integers
	fmt.Printf("%v \n", 198.454)        // floating point numbers
	fmt.Printf("%v \n", 3+7i)           // complex numbers
	fmt.Printf("%v \n", "Hello World!") // strings
	fmt.Printf("%v \n", []int{1, 2, 3}) // slices and arrays
	fmt.Printf("%v \n", m)              // maps
	fmt.Printf("%v \n", s)              // structs
	fmt.Printf("%v \n", c)              // channels
	fmt.Printf("%v \n", &s)             // pointers
}

func vHastashInterpolation()  {
	// struct
	s := struct {
		name string
		age  int
	}{"John", 26}

	// channel
	c := make(chan int)

	// map
	m := map[string]int{
		"one": 1,
		"two": 2,
	}

	fmt.Printf("%#v \n", true)           // boolean values
	fmt.Printf("%#v \n", 132)            // signed and unsigned integers
	fmt.Printf("%#v \n", 198.454)        // floating point numbers
	fmt.Printf("%#v \n", 3+7i)           // complex numbers
	fmt.Printf("%#v \n", "Hello World!") // strings
	fmt.Printf("%#v \n", []int{1, 2, 3}) // slices and arrays
	fmt.Printf("%#v \n", m)              // maps
	fmt.Printf("%#v \n", s)              // structs
	fmt.Printf("%#v \n", c)              // channels
	fmt.Printf("%#v \n", &s)             // pointers
}

func tTypeFormat()  {
	// struct
	s := struct {
		name string
		age  int
	}{"John", 26}

	// channel
	c := make(chan int)

	// map
	m := map[string]int{
		"one": 1,
		"two": 2,
	}

	fmt.Printf("%T \n", true)           // boolean values
	fmt.Printf("%T \n", 132)            // signed and unsigned integers
	fmt.Printf("%T \n", 198.454)        // floating point numbers
	fmt.Printf("%T \n", 3+7i)           // complex numbers
	fmt.Printf("%T \n", "Hello World!") // strings
	fmt.Printf("%T \n", []int{1, 2, 3}) // slices and arrays
	fmt.Printf("%T \n", m)              // maps
	fmt.Printf("%T \n", s)              // structs
	fmt.Printf("%T \n", c)              // channels
	fmt.Printf("%T \n", &s)             // pointers
}

func tBooleanFormat()  {
	fmt.Printf("%t \n", true)
	fmt.Printf("%t \n", false)
}

func dBase10Format()  {
	fmt.Printf("[1] %d\n", 0xFF) // dexadecimal
	fmt.Printf("[2] %d\n", 011) // octal
	fmt.Printf("[3] %d\n", 'β') // rune
	fmt.Printf("[4] %+d\n", 11) // print sign
	fmt.Printf("[5] % d\n", 11) // leave space when sign is not mentioned
}

func bBinaryFormat()  {
	fmt.Printf("9 in base2 is %b \n", 9)
	fmt.Printf("128 in base2 is %b \n", 128)
	fmt.Printf("15 in base2 is %b \n", 15)
	fmt.Printf("0xF in base2 is %b \n", 0xF) // hexadecimal
}

func xhexadecimalformat()  {
	fmt.Printf("Hex of %v is %x\n", 124, 124)
	fmt.Printf("Hex UC of %v is %X\n", 124, 124) // with upper-case
	fmt.Printf("Hex of %v is %#x\n", 124, 124) // adds leading '0x'


	fmt.Printf("Oct of %v is %o\n", 9, 9)
	fmt.Printf("Oct of %v is %#o\n", 9, 9) // adds leading '0'
}

func fFuntionFormat()  {
	fmt.Printf("[1] %f\n",1225.04532684676) // defalt width and preicision
	fmt.Printf("[2] %.11f\n",-1225.04532684676) // defalt width
	fmt.Printf("[3] %20.12f\n",1225.04532684676)
	fmt.Printf("[4] %20f\n",1225.04532684676) // defalt precision
	fmt.Printf("[5] %20.f\n",1225.04532684676) // zero precision
	fmt.Printf("[6] %.f\n",1225.04532684676) // defalt width and zero precision
	fmt.Printf("[7] %#.f\n",1225.04532684676) // always show decimal point

	fmt.Printf("[8] %+.2F\n",1225.04532684676) // alias of %f (with sign)
}