package main

import (
	"errors"
	"fmt"
)

func main() {
	myString := "I am a string"
	num1 := 42
	num2 := 7
	num3 := 32
	num4 := 9
	printMe(myString)
	fmt.Println(intDivision(num1, num2))
	var result, remainder, err = intDivisionWithRemainder(num3, num4)
	if err != nil {
		fmt.Printf(err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result of the division is %v", result)
	} else {
		fmt.Printf("The result of the division is %v with remainder %v", result, remainder)
	}

	var powerOfTwo = powerOfTwo(num3)
	switch {
	case powerOfTwo%2 == 0:
		fmt.Printf("\nThe power of two of %v is %v and is divisible by two", num3, powerOfTwo)
	case powerOfTwo%2 != 0:
		fmt.Printf("The power of two of %v is %v and is not divisible by two", num3, powerOfTwo)
	default:
		fmt.Printf("Unknown error")
	}

}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(num int, denum int) int {
	return num / denum
}

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
