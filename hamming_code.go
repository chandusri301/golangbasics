/*this is to find the total code difference in a sequence*/
package main

import (
	"errors"
	"fmt"
)

func distance(a, b string) (int, error) {
	hd := 0
	if len(a) == len(b) {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				hd++
			}

		}
		return hd, nil
	}
	return hd, errors.New("length should be same")
	panic("use same length")

}
func main() {
	a := ("ABCDEFGH")
	var b string
	fmt.Println("enter dna code (capsonly)=")
	fmt.Scan(&b)
	dist, errs := distance(a, b)
	if errs != nil {
		fmt.Println(errs)
	} else {
		fmt.Println("the length=", dist)
	}
}
