// GoLang learning material .. README
// CREDIT: https://www.youtube.com/watch?v=YS4e4q9oBaU&ab_channel=freeCodeCamp.org

package main

/*

 */

import (
	"fmt"
)

func main() {
	//Example 1
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Let's Go"))

	//Example 2
	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i:=0; i<10; i++ {
		fmt.Println(inc.Increment())
	}

}

// -------------------------------Interfaces
// Example 1
type Writer interface { //interfaces don't describe data, they describe behaviours. structs define data types, interfaces define methods.
	Write([]byte) (int, error)
}

type ConsoleWriter struct{} // a way to implement the interface

func (cs ConsoleWriter) Write(data []byte) (int, error) { //instead of using something like "implements" keyword, we implicitly implement Writer interface by creating a method on our ConsoleWriter construct that has the signiture of a Writer interface
	n, err := fmt.Println(string(data))
	return n, err
}

//End of Example 1

//Example 2
type Incrementer interface {
	Increment() int
}
type IntCounter int
func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}