package go_type

import "time"

// 斐波那契数列
func (f *Fibonacci) Sequence(count int, timeout time.Duration, cb func([]int)) {
	c := make(chan int)       // 通道-循环结果
	q := make(chan bool)      // 通道-结束循环
	go f.set(c, q, timeout)   // 同步发送
	go f.get(count, c, q, cb) // 同步接收
}
func (f *Fibonacci) set(c chan<- int, q <-chan bool, d time.Duration) {
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
func (f *Fibonacci) get(n int, c <-chan int, q chan<- bool, cb func([]int)) {
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
