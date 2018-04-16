package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func cleanup() {
	defer wg.Done() //we really wanna make sure we defer to done
	if r := recover(); r != nil {
		fmt.Println("Recovered in cleanup", r)
	}
	//wg.Done()
}
func say(s string) {
	defer cleanup()
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
		if i == 2 {
			panic("Oh dear , it's a panic " + s)
		}
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
