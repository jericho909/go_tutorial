package main

import (
	"errors"
	"fmt"
)

func main() {
	//variables
	myString := "I am a string"
	num1 := 42
	num2 := 7
	num3 := 32
	num4 := 9
	//calling functions
	printMe(myString)
	fmt.Println(intDivision(num1, num2))
	var result, remainder, err = intDivisionWithRemainder(num3, num4)
	//if loops
	if err != nil {
		fmt.Printf(err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result of the division is %v", result)
	} else {
		fmt.Printf("The result of the division is %v with remainder %v", result, remainder)
	}
	//switch statements
	var powerOfTwo = powerOfTwo(num3)
	switch {
	case powerOfTwo%2 == 0:
		fmt.Printf("\nThe power of two of %v is %v and is divisible by two", num3, powerOfTwo)
	case powerOfTwo%2 != 0:
		fmt.Printf("The power of two of %v is %v and is not divisible by two", num3, powerOfTwo)
	default:
		fmt.Printf("Unknown error")
	}
	//array slices
	var intArr [3]int32
	intArr[0] = 42
	fmt.Println(intArr[0])

	intArr2 := [...]int32{1, 2, 3, 4}
	fmt.Println(intArr2)

	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))
	//appending to array
	intSlice = append(intSlice, 7)
	fmt.Printf("\nThe length is %v with capacity %v\n", len(intSlice), cap(intSlice))
	fmt.Println(intSlice)

	var intSlice2 []int32 = []int32{9, 10, 11}
	intSlice = append(intSlice, intSlice2...)
	fmt.Printf("\nThe length is %v with capacity %v\n", len(intSlice), cap(intSlice))
	fmt.Println(intSlice)
	//maps that store key value pairs
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Kevin": 42, "Mary": 31, "Osman": 66}
	fmt.Println(myMap2["Kevin"])
	var age, ok = myMap2["Benjamin"]
	if ok {
		fmt.Printf("The age is %v", age)
	} else {
		fmt.Println("Invalid name")
	}
	//iterate over a map
	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Age: %v\n", name, age)
	}
	//iterate over an array
	for i, j := range intSlice {
		fmt.Printf("Index: %v, Value: %v \n", i, j)
	}
	//for loop
	for i := 0; i < 5; i++ {
		fmt.Printf("I am inside a for loop. This is loop number %v\n", i+1)
	}

}

// void function
func printMe(printValue string) {
	fmt.Println(printValue)
}

// function with parameter
func intDivision(num int, denum int) int {
	return num / denum
}

// function that returns multiple variables
func intDivisionWithRemainder(num int, denum int) (int, int, error) {

	var err error
	if denum == 0 {
		err = errors.New("Cannot divide by zero")
		return 0, 0, err
	}
	var result int = num / denum
	var remainder int = num % denum
	return result, remainder, err
}

func powerOfTwo(num int) int {
	return num * num
}
