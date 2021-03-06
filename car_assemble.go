/*In this exercise you'll be writing code to analyze the production of an assembly line in a car factory.
The assembly line's speed can range from `0` (off) to `10` (maximum).

At its default speed (`1`), `221` cars are produced each hour.
In principle, the production increases linearly.
So with the speed set to `4`, it should produce `4 * 221 = 884` cars per hour.
However, higher speeds increase the likelihood that faulty cars are produced, which then have to be discarded.

Also, there are times when the assembly line has an artificially imposed limit on the throughput (meaning no more than the limit can be produced per hour).

## 1. Calculate the success rate

Implement a function (`SuccessRate`) to calculate the ratio of an item being created without error for a given speed.
The following table shows how speed influences the success rate:

- `0`: 0% success rate.
- `1` - `4`: 100% success rate.
- `5` - `8`: 90% success rate.
- `9` - `10`: 77% success rate.

```go
rate := SuccessRate(6)
fmt.Println(rate)
// Output: 0.9
```

## 2. Calculate the production rate per hour

Implement a function to calculate the assembly line's production rate per hour:

```go
rate := CalculateProductionRatePerHour(7)
fmt.Println(rate)
// Output: 1392.3
```

> Note that the value returned is of type `float64`.

## 3. Calculate the number of working items produced per minute

Implement a function to calculate how many cars are produced each minute:

```go
rate := CalculateProductionRatePerMinute(5)
fmt.Println(rate)
// Output: 16
```

> Note that the value returned is of type `int`.

## 4. Calculate the artificially-limited production rate

Implement a function to calculate the assembly line's production rate per hour:

```go
rate := CalculateLimitedProductionRatePerHour(2, 1000.0)
fmt.Println(rate)
// Output: 442.0
rate := CalculateLimitedProductionRatePerHour(7, 1000.0)
fmt.Println(rate)
// Output: 1000.0
```*/
package main

import (
	"fmt"
)

const pr = 221

func successrate(i int) float64 {
	if i == 0 {
		return 0
	}
	if i == 1 || i == 2 || i == 3 || i == 4 {
		return 1.0
	}
	if i == 5 || i == 6 || i == 7 || i == 8 {
		return 0.9
	}
	if i == 9 || i == 10 {
		return 0.77
	}
	return 0
}
func prph(i int) float64 {
	return successrate(i) * float64(i) * pr
}
func prpm(i int) int {
	return int(prph(i)) / 60
}
func alp(i, a int) int {
	if int(prph(i)) <= a {
		return int(prph(i))
	} else {
		return a
	}
}
func main() {
	var i int
	var a int
	fmt.Print("enter the speed =")
	fmt.Scan(&i)
	fmt.Print("enter artificial limit=")
	fmt.Scan(&a)
	fmt.Println("the success rate is --------------->", successrate(i))
	fmt.Println("the production rate per hour is --->", prph(i))
	fmt.Println("the production rate per minute is ->", prpm(i))
	fmt.Println("the artificially limited rate is -->", alp(i, a))

}
