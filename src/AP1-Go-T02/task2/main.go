package main

import (
	"flag"
	"fmt"
)

func main() {
	k := flag.Int("k", 5, "Начальное значение (включительно)")
	n := flag.Int("n", 4, "Конечное значение (включительно)")

	flag.Parse()

	if *k > *n {
		fmt.Println("Конечное значение не может быть меньше начального")
		return
	}

	fmt.Printf("Начальное значение k = %d\n Конечное значение n = %d\n", *k, *n)

	numChan := generator(*k, *n)
	result := squarer(numChan)

	for num := range result {
		fmt.Println(num)
	}

}
func squarer(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for num := range in {
			out <- num * num
		}
	}()
	return out
}

func generator(k int, n int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for i := k; i <= n; i++ {
			out <- i
		}
	}()
	return out
}
