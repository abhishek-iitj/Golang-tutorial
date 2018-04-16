//Sending and receiving values over the channel and these are blocking
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func foo(c chan int, someVal int) {
	defer wg.Done()
	c <- someVal * 5 //put somevalue times 5 to channel
}
func main() {
	fooVal := make(chan int, 10) //fooVal is a channel of int with a buffer for 10 items
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go foo(fooVal, i)
	}
	wg.Wait() //before closing the channel we need to wait before all to complete.
	close(fooVal)

	for item := range fooVal {
		fmt.Println(item)
	}
	// fmt.Println(v1, v2)
}
