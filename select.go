package learngo

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func baseSelect() {
	bufferedChan := make(chan string, 2)
	bufferedChan <- "First"
	select {
	case str := <-bufferedChan:
		fmt.Println("read:", str)
	case bufferedChan <- "Second":
		fmt.Println("write:", <-bufferedChan)
	}

	unbuffChan := make(chan int)

	go func() {
		time.Sleep(time.Second)
		unbuffChan <- 1
	}()

	select {
	case bufferedChan <- "thrid":
		fmt.Println("Write third")
	case value := <-unbuffChan:
		fmt.Println("block reading:", value)
	case time := <-time.After(time.Millisecond * 1000):
		fmt.Println("time's up:", time)
	default:
		fmt.Println("default case")
	}

	resultChan := make(chan int)
	timer := time.After(time.Second) // timer outside loop

	go func() {
		defer close(resultChan)

		for i := 0; i < 10000; i++ {
			select {
			case <-timer:
				fmt.Println("Time's up")
				return
			default:
				time.Sleep(time.Nanosecond)
				resultChan <- i
			}
		}
	}()

	for value := range resultChan {
		fmt.Println("Result:", value)
	}
}

func graceFullShoutDown() {
	signalChan := make(chan os.Signal, 1)
	timer := time.After(10 * time.Second)

	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-timer:
		fmt.Println("time up")
		return
	case val := <-signalChan:
		fmt.Println("Stopped signal type:", val)
	}
}