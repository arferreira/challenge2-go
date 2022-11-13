package main

import "fmt"

func main() {
	slice := make([]int, 1, 10)
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			slice = append(slice, i)
		}
	}
	fmt.Println(slice)
}
