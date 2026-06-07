package learngo

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)



func baseContext() {
	ctx := context.Background()
	fmt.Println("ctx", ctx)

	ctxTodo := context.TODO()
	fmt.Println("ctxTodo", ctxTodo)

	ctxWithValue := context.WithValue(ctx, "name", "John")
	val := ctxWithValue.Value("name")
	fmt.Println(val)

	ctxWithCancel, cancel := context.WithCancel(ctxWithValue)
	fmt.Println(ctxWithCancel.Err())
	cancel()
	fmt.Println(ctxWithCancel.Err())

	ctxWithDeadline, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second*3))
	defer cancel()
	ctxTime, ok := ctxWithDeadline.Deadline()
	fmt.Printf("time: %v, ok: %v\n", ctxTime, ok)
	fmt.Println(ctxWithDeadline.Err())
	fmt.Println(<-ctxWithDeadline.Done())

	ctxWithTimeOut, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	fmt.Println(ctxWithTimeOut.Done())

}

func workerPool() {
	// ctx, cancel := context.WithCancel(context.Background())
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*20)
	defer cancel()

	// var wg sync.WaitGroup
	wg := sync.WaitGroup{}
	// fmt.Printf("wg: %v", wg)
	// wg := &sync.WaitGroup{}

	numberToProcess, processedNumber := make(chan int, 5), make(chan int, 5)

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ctx, numberToProcess, processedNumber)
		}()
	}

	go func() {
		for i := 0; i < 1000; i++ {
			if i == 500 {
				cancel()
			}
			numberToProcess <- i
		}
		close(numberToProcess)
	}()

	go func() {
		wg.Wait()
		close(processedNumber)
	}()

	var counter int
	for resultValue := range processedNumber {
		counter++
		fmt.Println(resultValue)
	}

	fmt.Println(counter)
}

func worker(ctx context.Context, toProcess <-chan int, processed chan<- int) {
	for {
		select {
		case <-ctx.Done():
			// fmt.Println("Done")
			return
		case val, ok := <-toProcess:
			if !ok {
				return
			}
			time.Sleep(time.Millisecond)
			processed <- val * val

		}
	}
}

func withoutErrorGroup() {
	var err error

	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}

	wg.Add(3)

	go func() {
		time.Sleep(time.Second)
		defer wg.Done()

		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("first started")
			time.Sleep(time.Second)
		}

	}()

	go func() {
		defer wg.Done()

		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("second started")
			err = fmt.Errorf("any Error")
			cancel()
		}

	}()

	go func() {
		defer wg.Done()

		select {
		case <-ctx.Done():
			return
		default:
			time.Sleep(time.Second * 10)
			fmt.Println("third started")
		}

	}()

	wg.Wait()
	fmt.Println(err)
}

func withErrorGroup() {
	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		time.Sleep(time.Second)

		select {
		case <-ctx.Done():
			return nil
		default:
			fmt.Println("First started")
			time.Sleep(time.Second)
			return nil
		}
	})

	g.Go(func() error {
		fmt.Println("Start second")
		return fmt.Errorf("Unexpected error in 2")
	})

	g.Go(func() error {
		select {
		case <-ctx.Done():
		default:
			fmt.Println("thrid started")
			time.Sleep(time.Second)
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		fmt.Println("Error", err)
	}

}