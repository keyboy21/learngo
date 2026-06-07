package learngo

import (
	"fmt"
	"sync"
	"time"
)

func channels() {
	var nilChannel chan int

	fmt.Printf("Type: %T, Value: %v\n", nilChannel, nilChannel)
	fmt.Printf("Len: %v, Cap: %v\n", len(nilChannel), cap(nilChannel))

	// write to chanel blocks forever
	// nilChannel <- 1

	// read from chanel
	// <- nilChannel

	// close chanel
	// close(nilChannel)

	unBufferedChanel := make(chan int)
	fmt.Printf("Type: %T, Value: %v\n", unBufferedChanel, unBufferedChanel)
	fmt.Printf("Len: %v, Cap: %v\n", len(unBufferedChanel), cap(unBufferedChanel))

	// block until smb reads
	// unBufferedChanel <- 1
	// <- unBufferedChanel

	// block on reading then write
	go func(chanForWriting chan<- int) {
		time.Sleep(time.Second)
		chanForWriting <- 1
	}(unBufferedChanel)

	val := <-unBufferedChanel
	fmt.Println(val)

	go func(chanForReading <-chan int) {
		time.Sleep(time.Second)
		value := <-chanForReading
		fmt.Println(value)
	}(unBufferedChanel)

	unBufferedChanel <- 10

	// panic
	close(unBufferedChanel)
	close(unBufferedChanel)
}

func bufferedChannel() {
	bufferedChan := make(chan int, 2)
	fmt.Printf("Len: %v, Cap: %v\n", len(bufferedChan), cap(bufferedChan))

	// doesn't block while buffer not full
	bufferedChan <- 2
	bufferedChan <- 4

	fmt.Printf("Len: %v, Cap: %v\n", len(bufferedChan), cap(bufferedChan))

	// blocks to write, buffer is full
	// bufferedChan <- 6

	fmt.Println(<-bufferedChan)
	fmt.Println(<-bufferedChan)

	fmt.Printf("Len: %v, Cap: %v\n", len(bufferedChan), cap(bufferedChan))

	// go func(chanForWriting chan<- int) {
	// 	chanForWriting <- 20
	// 	chanForWriting <- 21
	// }(bufferedChan)

	// block to read, buffer is empty, no write, deadlock
	fmt.Println(<-bufferedChan)
	fmt.Println(<-bufferedChan)

}

func chanWithRange() {
	bufferedChan := make(chan int, 3)
	numbers := []int{1, 2, 4, 3}

	go func() {
		for _, num := range numbers {
			bufferedChan <- num
		}
		close(bufferedChan)
	}()

	for {
		// v := <-bufferedChan
		v, ok := <-bufferedChan
		fmt.Println(v, ok)
		if !ok {
			break
		}
	}

	bufferedChan = make(chan int, 10)

	go func() {
		for _, num := range numbers {
			bufferedChan <- num
		}
		close(bufferedChan)
	}()

	for v := range bufferedChan {
		fmt.Println(v)
	}

	unBufferedChan := make(chan int)

	go func() {
		for _, val := range numbers {
			unBufferedChan <- val
		}
		close(unBufferedChan)
	}()

	for v := range unBufferedChan {
		fmt.Println(v)
	}
}

func chanAsMutex() {
	var counter int

	mutexChan := make(chan struct{}, 1)
	wg := sync.WaitGroup{}
	// mu := sync.Mutex{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			mutexChan <- struct{}{}

			// mu.Lock()
			counter++
			// mu.Unlock()

			<-mutexChan
		}()
	}

	wg.Wait()

	fmt.Println(counter)

}




// MARK: channels
	// var channel chan int // by default channel is nil
	// fmt.Println("channel is nil:", channel == nil)
	// fmt.Printf("Type: %T , Value: %#v\n", channel, channel)
	// fmt.Printf("Length: %d , Capacity: %d\n", len(channel), cap(channel))

	// write to nil channel blocks forever: deadlock
	// channel <- 1

	// read from nil channel blocks forever: deadlock
	// <-channel

	// close nil channel will raise panic
	// close(channel)

	// var unBufferedChannel = make(chan int)
	// fmt.Println("unBufferedChannel is nil:", unBufferedChannel == nil)
	// fmt.Printf("Length: %d , Capacity: %d\n", len(unBufferedChannel), cap(unBufferedChannel))

	// blocks until write to unbuffered channel
	// unBufferedChannel <- 1

	// blocks until read from unbuffered channel
	// <-unBufferedChannel

	// only read from unbuffered channel
	// go func(chanToWrite chan<- int) {
	// 	time.Sleep(time.Second)
	// 	chanToWrite <- 3
	// }(unBufferedChannel)

	// value := <-unBufferedChannel

	// fmt.Println("value:", value)
