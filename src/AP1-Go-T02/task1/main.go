package main

import (
	"flag"
	"fmt"
	"math/rand"
	"runtime"
	"sort"
	"sync"
	"time"
)

type Result struct {
	ID       int
	Duration int
}

func main() {
	n := flag.Int("n", 0, "Количество горутин")
	m := flag.Int("m", 0, "Максимальное время сна в миллисекундах")
	flag.Parse()

	fmt.Println("Количество горутин: ", *n, "Максимальное время сна в миллисекундах: ", *m)

	results := make([]Result, *n)
	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup
	wg.Add(*n)

	for i := 0; i < *n; i++ {
		go func(i int) {
			defer wg.Done()
			defer fmt.Printf("Go рутина #%d - Done\n", i)
			sleep := rand.Intn(*m + 1)
			time.Sleep(time.Duration(sleep) * time.Millisecond)
			results[i] = Result{ID: i, Duration: sleep}
		}(i)
	}

	wg.Wait()
	fmt.Println("До сортировки:")
	for _, r := range results {
		fmt.Printf("Горутина #%d спала %d мс\n", r.ID, r.Duration)
	}
	fmt.Println(runtime.NumCPU())

	sort.Slice(results, func(i int, j int) bool {
		return results[i].Duration > results[j].Duration
	})

	for _, r := range results {
		fmt.Println("Номер рутины: ", r.ID)
		fmt.Println("Время сна: ", r.Duration)

	}
}
