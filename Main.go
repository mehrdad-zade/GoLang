// GoLang learning material .. README
// CREDIT: https://www.youtube.com/watch?v=YS4e4q9oBaU&ab_channel=freeCodeCamp.org

/*
concurrency and parallelism are core to GoLang. Channels synchronize data transmission between multiple routines. If
we are working with unbuffered channels, then we need to have a receiver for every sender.
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

// Example 5
type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logCH = make(chan logEntry, 50)
var doneCH = make(chan struct{}) // the common way to use a channel for ending a routin that is infinitly looping over a channel

const (
	logInfo     = " INFO "
	logWarnning = " WARNING "
	logError    = " ERROR "
)

func main() {
	//Example 1 - create a channel with make function
	fmt.Println("--Example1:")
	ch := make(chan int) // data type int will fllow into the channel. this is strongly typed, so only int can go through
	wg.Add(2)
	go func() {
		i := <-ch // receiving data from the channel
		fmt.Println(i)
		wg.Done()
	}()
	go func() {
		ch <- 42 // sending data to the channel
		wg.Done()
	}()
	wg.Wait()

	//Example 2: two routines, each acting as both sender and receiver. usually each routine will only do either though
	fmt.Println("--Example2:")
	wg.Add(2)
	go func() {
		i := <-ch // sender
		fmt.Println(i)
		ch <- 27 // receiver
		wg.Done()
	}()
	go func() {
		ch <- 42          // receiver
		fmt.Println(<-ch) //sender
		wg.Done()
	}()
	wg.Wait()

	//Example 3: A way to make the routine a sender-only and receiver-only is by passing the channel parameter in the way expected
	fmt.Println("--Example3:")
	wg.Add(2)
	go func(ch <-chan int) { // ch is bound to send int only; cannot be a sender
		i := <-ch // receiver
		fmt.Println(i)
		//ch <- 27 // sender is not allowed
		wg.Done()
	}(ch)
	go func(ch chan<- int) { // ch can only receive int; cannot send
		ch <- 42 // sender
		//fmt.Println(<-ch) //receiver is not allowed
		wg.Done()
	}(ch)
	wg.Wait()

	//Example 4: Buffered channels: sometimes your sender routine geenrates data faster than receiver-routine can consume because it might be busy processing them. As we know
	//we should have a receiver for every sender, so in such cases we need to have a buffered channel. With buffered channels, if we have less receivers, it will only be able
	//handle the messages in a FIFO manner, and the remaining will be lost. The workaround/solution is to have a loop on the receiver side, and close the channel on the sender.
	fmt.Println("--Example4:")
	bufferedCH := make(chan int, 50)
	wg.Add(2)
	go func(bufferedCH <-chan int) {
		for i := range bufferedCH { // pay attention to i. it's not an index, it's the actual value send to ch
			fmt.Println(i)
		}
		wg.Done()
	}(bufferedCH)
	go func(bufferedCH chan<- int) {
		bufferedCH <- 1
		bufferedCH <- 2
		bufferedCH <- 3
		bufferedCH <- 4
		bufferedCH <- 50  // send as many data as the buffer size
		close(bufferedCH) // make sure to close the ch, otherwise the for loop above won't know we are done sending. Also make sure not to send a msg to a closed channel.
		wg.Done()
	}(bufferedCH)
	wg.Wait()

	//Example 5: a coomon way to keep the channel running and waiting on data, but to close it properly when we are done
	fmt.Println("--Example 5:")
	go logger()
	logCH <- logEntry{time.Now(), logInfo, " - App is starting"}

	logCH <- logEntry{time.Now(), logInfo, " - App is shutting down"}
	time.Sleep(100 * time.Millisecond)
	doneCH <- struct{}{} //breaking out of the logger for loop. when the selectblock (listener of the logger func) receives this instruction, it will terminate the listener/loop/routine
}

// Example 5
func logger() {
	for {
		select { // this acts like a listener waiting for a message to be received by either of the channels
		case entry := <-logCH:
			fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
		case <-doneCH:
			break
			// you can use a default too, if you want to run some logic when you haven't received any messages by either of the channels
		}
	}
}
