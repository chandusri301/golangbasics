package main

import "fmt"

func main() {
	i := 10
	a := 0
	for i = 0; i <= 10; i++ {
		a = a + i
		fmt.Println(i)
	}
	fmt.Println("the total of the loop=", a)
}
