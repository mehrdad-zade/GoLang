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


}

//-------------------------------Interfaces
//Example 1
type Writer interface { //interfaces are data containers; they don't describe data, they describe behaviours
	Write([]byte) (int, error)
}

type ConsoleWriter struct {} // a way to implement the interface

func (cs ConsoleWriter) Write(data []byte) (int, error){
	n, err := fmt.Println(string(data))
	return n, err
}
//End of Example 1