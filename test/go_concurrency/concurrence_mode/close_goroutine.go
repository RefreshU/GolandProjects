package concurrence_mode

import (
	"fmt"
	"math/rand"
	"time"
)

func closeGoroutineCase1(){
	doWork := func(
		done <-chan interface{},
		strings <-chan string,
		) <-chan interface{}{
		terminated := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			close(terminated)
			for {
				select {
				case s := <-strings:
					fmt.Println(s)//
				case <-done:
					return
				}
			}
		}()
		return terminated
	}

	done := make(chan interface{})
	terminated := doWork(done, nil)
	go func() {
		//cancel doWork after sleep 1 second
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling doWork goroutine...")
		close(done)
	}()
	<-terminated
	fmt.Println("Done.")
}

func closeGoroutineCase2(){
	newRandStream := func(done <-chan interface{}) <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exited.")
			defer close(randStream)
			for{
				select {
				case randStream <-rand.Int():
				case <-done:
					return
				}
			}
		}()
		return randStream
	}
	done := make(chan interface{})
	randStream := newRandStream(done)
	fmt.Println("3 random ints:")
	for i := 1; i < 3; i++ {
		fmt.Println("%d: %d\n", i, <-randStream)
	}
	close(done)
	//sleep 1 seconds for goroutine
	time.Sleep(1 * time.Second)
	fmt.Println("Done.")
}