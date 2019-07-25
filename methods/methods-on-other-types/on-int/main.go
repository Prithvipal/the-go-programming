package main

import "fmt"

func main() {
	var n number = 7
	fmt.Println(n.isPrime())
	n.increamentBy(3)
	fmt.Println(n)
}

type number int

func (n number) isPrime() bool {
	for i := 2; i < int(n); i++ {
		if int(n)%i == 0 {
			return false
		}
	}
	return true
}

func (n number) isDivisible(i number) bool {
	return (n % i) == 0
}

func (n *number) increamentBy(i number) {
	*n = *n + i
}
