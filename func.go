package main

import (
	"errors"
	"fmt"
)

func greetings(name string) (string, error) {
	if name == "" {
		return name, errors.New(" name cannot be empty!")
	}
	return "hello," + name, nil
}

func add(num ...int) int {
	sum := 0
	for _, i := range num {
		sum += i
	}

	return sum
}

func main() {
	greeting, err := greetings("chandu")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(greeting)
	}
	fmt.Println("the sum of 1,2,3,4,5 is ", add(1, 2, 3, 4, 5))
	fmt.Println("the sum of 1,2,3,4,5,6 is ", add(1, 2, 3, 4, 5, 6))
}
