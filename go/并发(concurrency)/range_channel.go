package main

import "fmt"

func main() {
	ch := make(chan int)

	for i := 0; i < 5; i++ {
		go func() {
			ch <- i
		}()
	}

	for v := range ch {
		fmt.Println(v)
	}


/*
5
5
5
5
5
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()
	/tmp/sandbox982202598/main.go:14 +0x1e0
*/


	{

		ch := make(chan int)
		limit := 5

		for i := 0; i < limit; i++ {
			go func(i int) {
				ch <- i
			}(i)
		}

		cn := 0
		for v := range ch {
			fmt.Println(v)
			cn++
			if cn == limit {
				close(ch)
			}
		}
		// 0
		// 1
		// 2
		// 3
		// 4

		v, ok := <-ch
		fmt.Println(v, ok) // 0 false
		// any value received from closed channel succeeds without blocking
		// , returning the zero value of channel type and false.

	}
}
