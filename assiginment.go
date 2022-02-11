package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var ord string
	var quantity int
	var sales = []int{}
	billbox := 0
	for true {
		fmt.Println("WELCOME TO GO CAFE")
		fmt.Println("      MENU")
		fmt.Print("C-->Coffee-----₹40\n D-->Dosa-----₹80 rs\n T-->Tomato Soup-----₹20\n I-->Idli-----₹20\n V-->Vada-----₹25\n B-->Bhature Chhole-----₹30\n P-->Paneer Pakoda-----₹30\n M-->Manchurian-----₹90\n H-->Hakka Noodle-----₹70\n F-->French fries-----30\n J-->Jalebi-----₹30\n L-->Lemonade-----₹15\n S-->Spring roll-----₹20\n")
		fmt.Println("________________________________________________")
		fmt.Println("Enter END To close the day.")
		fmt.Println("place your order: ")
		fmt.Scan(&ord)
		ord = strings.ToUpper(ord)
		if ord == "END" {
			total_income(sales)
		} else {
			fmt.Scan(&quantity)
		}
		bills := bill(quantity, ord)
		fmt.Println("_______________________")
		fmt.Println("Your Total bill for the order", quantity, ord, "is ₹", bills)
		fmt.Println("_______________________")
		sales = append(sales, bills)
		billbox++

	}
}

func total_income(sale []int) {
	var sum int = 0

	for i := 0; i < len(sale); i++ {
		sum = sale[i] + sum
	}
	fmt.Println("Today's total income is --> ₹ ", sum)
	os.Exit(0)
}

func bill(quantity int, ord string) int {
	a := map[string]int{
		"C": 40,
		"D": 80,
		"T": 20,
		"I": 20,
		"V": 25,
		"B": 30,
		"P": 30,
		"M": 90,
		"H": 70,
		"F": 30,
		"J": 30,
		"L": 15,
		"S": 20,
	}
	amount := quantity * a[ord]
	return amount
}
