package utils

import "time"

// 斐波那契数列: 并行运算
func (f *Fibonacci) FibonacciToDo(count int, timeout time.Duration, do func([]int)) {
	c := make(chan int)                    // 通道-循环结果
	q := make(chan bool)                   // 通道-结束循环
	go f.setFibonacciChan(c, q, timeout)   // 同步发送
	go f.getFibonacciChan(count, c, q, do) // 同步接收
}

func (f *Fibonacci) setFibonacciChan(c chan<- int, q <-chan bool, d time.Duration) {
	x, y := 1, 1
	for {
		select {
		case c <- x: // 未处理关闭时异常,因为下面退出时才会关闭
			x, y = y, x+y
		case <-time.After(d): // 超时
		case <-q: // 结束
			close(c) // Indicate to our routine to exit cleanly upon return.
			return   // 不使用break,因为它只能跳出select
		}
	}
}

func (f *Fibonacci) getFibonacciChan(n int, c <-chan int, q chan<- bool, cb func([]int)) {
	s := make([]int, n, n) // 可省略第3个参数:容量(默认=第2个参数:长度)
	for i := 0; i < n; i++ {
		if x, ok := <-c; !ok { // 超时或结束
			break
		} else {
			s[i] = x
		}
	}
	q <- true
	close(q) // Indicate to our routine to exit cleanly upon return.
	cb(s)
}
