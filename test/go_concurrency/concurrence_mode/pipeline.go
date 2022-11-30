package concurrence_mode

import "fmt"

func echo(nums []int) <-chan int{
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int{
	out := make(chan int)
	go func() {
		for n := range in{
			out <- n * n
		}
		close(out)
	}()
	return out
}

func odd(in <-chan int) <-chan int{
	out := make(chan int)
	go func() {
		for n := range in{
			if n % 2 != 0{
				out <- n
			}
		}
		close(out)
	}()
	return out
}

func sum(in <-chan int)<-chan int{
	out := make(chan int)
	go func() {
		var sum	= 0
		for n := range in{
			sum =+ n
		}
		out <- sum
		close(out)
	}()
	return out
}

type EchoFunc func([]int) <-chan int
type Pipeline func(<-chan int) <-chan int
func pipleline(nums []int, echo EchoFunc, pipeFns ...Pipeline) <-chan int{
	ch := echo(nums)
	for i := range pipeFns{
		ch = pipeFns[i](ch)
	}
	return ch
}

func piplelineMode(){
	var nums = []int{1,2,3,4,5,6,7,8,9,10}
	for n := range pipleline(nums,echo, odd, sq, sum){
		fmt.Println(n)
	}

}