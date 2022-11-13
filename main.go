package main

import "fmt"

func main() {
	slice := make([]int, 1, 10)
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println("Pin")
			slice = append(slice, i)
		} else if i%5 == 0 {
			fmt.Println("Pan")
		}
		fmt.Println(i)
	}
	fmt.Println(slice)
}
