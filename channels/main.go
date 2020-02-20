package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	channel := make(chan int, 1)
	// txOnlyChannel := make(chan int, 1)

	wg.Add(2)
	go rxChannel(channel, &wg)
	go TxChannel(channel, &wg)

	wg.Wait()

}

func rxChannel(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	r := <-ch
	fmt.Println(r)

}

func TxChannel(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- 1
	fmt.Println("End Tx")
}
