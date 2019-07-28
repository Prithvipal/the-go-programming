package main

import "fmt"

func main() {

	var l list = []string{"A", "B", "C"}
	var input list = []string{"B", "C", "D"}
	intersResult := l.intersact(input)
	fmt.Println("Intersaction:", intersResult)

	l.remove("B")
	fmt.Println("After deletion:", l)
}

type list []string

func (l list) intersact(input list) list {
	mp := make(map[string]bool)
	for _, ele := range l {
		mp[ele] = true
	}
	var result list
	for _, ele := range input {
		if mp[ele] {
			result = append(result, ele)
		}
	}
	return result
}

func (l *list) remove(in string) {
	index := l.indexOf(in)
	*l = append((*l)[:index], (*l)[index+1:]...)
}

func (l list) indexOf(in string) int {
	for index, ele := range l {
		if ele == in {
			return index
		}
	}
	return -1
}
