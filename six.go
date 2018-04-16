//The defer makes ure for something to happen.
//In case of panic it makes sure recover.
//The defer statements are last ib first out manner.

package main

import "fmt"

func foo() {
	//it will be evaluated but run only when the function is done.
	defer fmt.Println("Done! from foo")
	defer fmt.Println("Are we done?")
	fmt.Println("Doing some stuff, who knows what?")
}

func bar() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
func main() {
	foo()
	bar()
}
