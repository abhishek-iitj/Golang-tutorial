//Defining variables,
//types, fmt
// functions return
//loops
//Map data structure
package main

import (
	"fmt"
	"math"
	"math/rand"
)

func add(x float64, y float64) float64 {
	return x + y
}
func add2(x, y float32) float32 {
	return x + y
}

func multipleReturnType(a, b string) (string, string) {
	return a, b
}

func foo() {
	fmt.Println("called from foo ")
	var num1 = 10
	num1++
}
func main() {
	//Go Math lib functions
	fmt.Println("sqrt of 5 is", math.Sqrt(4), " ! ")
	fmt.Println("A random no. from 1-100", rand.Intn(100), " ! ")

	//Defining  a variable
	var num1 float64 = 5.6
	var num2 float64 = 9.56
	var num3, num4 float64 = 6.7, 9.888

	fmt.Println(add(num1, num2))
	fmt.Println(add(num3, num4))

	num5, num6 := 6.7, 8.0 //go gives it float 64
	//fmt.Println(add2(num5, num6)) //will thorw an error;
	fmt.Println(add(num5, num6))

	//multiple returns from a function
	w1, w2 := "Hey", "Rashi"
	fmt.Println(multipleReturnType(w1, w2))

	//type conversion
	var c int = 64
	var d float64 = float64(c)
	fmt.Println(c, d)

	//Pointers in GO..................
	x := 15
	a := &x //pointer to x
	fmt.Println(x, a, *a)
	*a = 5 //x will be changed;
	fmt.Println(x, a, *a)
	*a = *a * *a
	fmt.Println(x, a, *a)

	//Arrays

	//Loops.....
	//1
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	//2
	j := 0
	for j < 20 {
		fmt.Println(j)
		j++
	}
	//3
	y := 5
	for {
		fmt.Println("DO Stuff", y)
		y += 3
		if y > 25 {
			break
		}
	}

	//Map data structure key:value
	grades := make(map[string]float32) //key is string and value is float32
	grades["Timmy"] = 42
	grades["Jess"] = 92
	grades["Sam"] = 67
	fmt.Println(grades)

	TimsGrade := grades["Timmy"]
	fmt.Println(TimsGrade)

	delete(grades, "Timmy")
	fmt.Println(grades)

	//iterating through the map
	for k, v := range grades {
		fmt.Println(k, " : ", v)
	}

}
