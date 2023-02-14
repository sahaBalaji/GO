package main

import "fmt"

func main() {

	var num, i int

	fmt.Print("\nEnter the Maximum Natural Number Limit = ")
	fmt.Scanln(&num)

	fmt.Println("\nNatural Numbers from 1 to ", num, " are = ")
	for i = 1; i <= num; i++ {
		fmt.Print(i, "\t")
	}
	fmt.Println()
}
