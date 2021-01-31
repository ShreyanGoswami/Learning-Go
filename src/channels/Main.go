package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int) // Channels are strongy typed
	wg.Add(2)
	go func() {
		i := <-ch // Receiving data from channel
		fmt.Println(i)
		wg.Done()
	}()
	go func() {
		ch <- 42 // Sending data to channel. Only one value can be sent at a time
		wg.Done()
	}()
	wg.Wait()
	// Creating specific methods for receiving and sending data on channels
	wg.Add(2)
	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		ch <- 42
		wg.Done()
	}(ch)
	wg.Wait()
	// Buffered channel
	bufferedChannel := make(chan int, 50) // Declaring the size of the buffer, can store 50 integers
	wg.Add(2)
	go func(ch <-chan int) {
		// for i := range ch {
		// 	fmt.Println(i)
		// }
		for {
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				fmt.Println("Channel is closed")
				break
			}
		}
		wg.Done()
	}(bufferedChannel)
	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		close(ch) // On the sending side there is currently no way to check if a channel is closed through code
		wg.Done()
	}(bufferedChannel)
	wg.Wait()
}
