// GoLang learning material .. README
// CREDIT: https://www.youtube.com/watch?v=YS4e4q9oBaU&ab_channel=freeCodeCamp.org

package main

/*

 */

import (
	"fmt"
)

func main() {
	arg1, arg2 := "A", "B"
	myFunc1(arg1, arg2) // pass by value
	fmt.Println("func is pass by val: ", arg1)

	myFunc2(&arg1, &arg2) // pass by reference. useful when you are passing a huge datastructure. this way you don't pass the data every time, but only the address to them
	fmt.Println("func is pass by val: ", arg1)

	funcSum("custom msg", 1, 2, 3, 4, 5)

	fmt.Println("Func with return: ", myFunc3(2, 3))
	fmt.Println("Func with named return: ", myFunc4(2, 3))

	divRes, err := myDiv(5, 0)
	fmt.Printf("get multiple returns: %v, %v\n", divRes, err)

	//dealing with methods

	g := greeter {
		name : "me",
		greet : "hi",
	}
	g.getGreet()

}

// ----------------------------------------------------Functions
func myFunc1(arg1, arg2 string) { //when the type of all args are the same, you don't have to declare it every time
	fmt.Println("Fun1------")
	fmt.Println("functioan call - myFunc1:", arg1, arg2)
	arg1 = "C"
	fmt.Println("func is pass by val: ", arg1)
}

func myFunc2(arg1, arg2 *string) { //when the type of all args are the same, you don't have to declare it every time
	fmt.Println("Fun1------")
	fmt.Println("functioan call - myFunc1:", *arg1, *arg2)
	*arg1 = "C"
	fmt.Println("func is pass by ref: ", *arg1)
}

func funcSum(msg string, values ...int) { //this is a way to assign multiple values to a slice
	fmt.Printf("vals: %v %v, type: %T\n", msg, values, values)
}

func myFunc3(x, y int) int {
	return x + y
}

func myFunc4(x, y int) (res int) {
	res = x + y
	return // named return, helps with not having to declare the returned var twice
}

func myDiv(x, y float64) (float64, error) {
	if y == 0 {
		return 0.0, fmt.Errorf("Denom is zero, cannot complete the devision")
	}
	return x / y, nil
}


// ----------------------------------------------------Methods
//Methods are similar to functions, however, a datatype gets context
type greeter struct {
	name string
	greet string

}
func (g greeter) getGreet() {
	fmt.Println("Methods: ", g.name, g.greet)
}