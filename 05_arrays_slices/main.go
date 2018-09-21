package main

import "fmt"

func main() {
	// Arrays
	//var fruitArr [2]string

	// Assign values
	//fruitArr[0] = "apple"
	//fruitArr[1] = "Orange"

	// Declare and assign
	// fruitArr := [2]string{"Apple", "Orange"}

	// fmt.Println(fruitArr)
	// fmt.Println(fruitArr[1])

	fruitSlice := []string{"apple", "Orange", "grape", "Cherry"}

	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])
}
