package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func say(s string) {
	defer wg.Done() //what if we dont actually get to get here so we need Defer.

	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}
	// wg.Done() //what if we dont actually get to get here so we need Defer.

}

func main() {
	wg.Add(1)
	go say("Hey")
	wg.Add(1)
	go say("There")
	//wg.Wait()
	wg.Add(1)
	go say("Hi")
	wg.Wait()
}
