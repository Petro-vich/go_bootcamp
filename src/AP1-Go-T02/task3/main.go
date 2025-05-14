package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	signChan := make(chan os.Signal, 1)
	signal.Notify(signChan, syscall.SIGINT, syscall.SIGTERM)
	k := flag.Int("k", 1, "<UNK>")
	flag.Parse()

	fmt.Printf("Тикер будет срабатывать каждые %d секунд\n", *k)

	start := time.Now()

	done := make(chan struct{})

	go func() {
		tick := 1
		for {
			select {
			case <-done:
				return
			default:
				fmt.Println("Номер тика", tick)
				fmt.Println("Время с начала работы программы:", time.Since(start))
				time.Sleep(time.Duration(*k) * time.Second)
				tick++
			}
		}
	}()
	<-signChan
	fmt.Println("Termination")
	close(done)
}
