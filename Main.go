// GoLang learning material .. README
// CREDIT: https://www.youtube.com/watch?v=YS4e4q9oBaU&ab_channel=freeCodeCamp.org

/*
*
GoRoutines: allows us to work on multiple things concurrently.
GoLang enables an abstract light weight thread on OS level, which
is useful because you can define thousands of routines without much
overhead that other languages such as Java or C# create.
by marking a method with keyword "go" you have a routine.
*/
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg = sync.WaitGroup{} // designed to synch multiple go routines together
var counter = 0

// Example 4
var m = sync.RWMutex{} // mutext is a lock to make sure resources are used the way we want. RWMutext ensures that as many threads as we want can read this source, but as long as reading we cannot write to it

func main() {
	//Example 1 - Go Routine allows closure access to varibales, which is not a good practice, because go routine and the var will be on different stack traces that could create unexpected behaviours
	msg := "hello"
	go func() {
		fmt.Println("--Example1: ", msg) //the routine can understand what msg is refered to, although the routine and the var are on different stack traces
	}()
	msg = "goodbye"                    //the routine will usually pick this value, but that's not guaranteed because the routine and this var are on different stack traces
	time.Sleep(100 * time.Microsecond) //without this you don't give enough time to the OS scheduler to print the value of the routine. But this is a bad practice (because we are binding the application clock cycle with the real world clock which is unreliable) and used for demo only
	// to see the race condition clearly you can run the app like this: go run -race Main.go. this will return a DATA RACE message indicating that a routine (which is our func) has accessed a memory (our msg var) twice

	//Example 2 - if you pass the var as an arg to the routine, it will use the value at the time the routine was scheduled on the OS
	msg = "hello"
	go func(msg string) {
		fmt.Println("--Example2: ", msg)
	}(msg)
	msg = "goodbye"
	time.Sleep(100 * time.Microsecond)

	//Example 3 - run multiple go routines. you can see there is no synchronization between the routines. the two methods race against each other to get their work done
	fmt.Println("--Example3:")
	for i := 0; i < 10; i++ {
		wg.Add(2) // add two routines to the OS scheduler threads
		go sayHello()
		go increment()
	}
	wg.Wait() // wait for the results

	//Example 4 - use of mutext for managig resources access by each thread. this is basically destroying the use of parallelism; in fact the performanceo of running this without routines is higher (just a demo for mutext)
	counter = 0
	fmt.Println("--Example4:")
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello2()
		m.Lock()
		go increment2()
	}
	fmt.Println("Number of threads CPU has made available to the OS = ", runtime.GOMAXPROCS(-1)) // see the max num of threads available with param "-1". if you want to limit the number of threads use a positive int as a statement; i.e. runtime.GOMAXPROCS(1)
	wg.Wait()
}

// Example 3
func sayHello() {
	fmt.Printf("Hello #%v\n", counter)
	wg.Done() // tell the weight group that are routine is done
}
func increment() {
	counter++
	wg.Done()
}

// Example 4
func sayHello2() {
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	wg.Done() // tell the weight group that are routine is done
}
func increment2() {
	counter++
	m.Unlock()
	wg.Done()
}
