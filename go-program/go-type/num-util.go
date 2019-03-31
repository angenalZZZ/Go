package go_type

import "fmt"

// 斐波那契数列
func FibonacciSequence(n int) (s []int) {
	s = make([]int, n, n) // 可省略第3个参数:容量(默认=第2个参数:长度)
	c := make(chan int)   // 同步通道c 用于n次循环
	q := make(chan bool)  // 同步通道q 用于结束循环

	go getFibonacciSequence(s, c, q) // 同步接收
	go setFibonacciSequence(c, q)    // 同步发送
	return
}
func setFibonacciSequence(c chan<- int, q <-chan bool) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-q:
			return
		}
	}
}
func getFibonacciSequence(s []int, c <-chan int, q chan<- bool) {
	for i, n := 0, cap(s); i < n; i++ {
		s[i] = <-c
	}
	q <- true
	fmt.Printf("  斐波那契数列: %v", s)
}
