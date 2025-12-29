package main

import "fmt"

func main() {
	const (
		add  string = "+"
		sub  string = "-"
		mult string = "*"
		div  string = "/"
		rem  string = "%"
	)
	var a, b float64
	var action string
	fmt.Scan(&a, &b, &action)
	switch action {
	case add:
		fmt.Println(a + b)
	case sub:
		fmt.Println(a - b)
	case mult:
		fmt.Println(a * b)
	case div:
		if b == 0 {
			fmt.Println("Делить на ноль нельзя!")
		} else {
			fmt.Println(a / b)
		}
	case rem:
		fmt.Println(int(a) % int(b))
	}
}
