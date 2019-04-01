package go_type

type Fibonacci struct {
}

// 斐波那契数列
func (f *Fibonacci) Sequence(count int, cb func([]int)) {
	c := make(chan int)       // 通道-循环结果
	q := make(chan bool)      // 通道-结束循环
	go f.set(c, q)            // 同步发送
	go f.get(count, c, q, cb) // 同步接收
}
func (f *Fibonacci) set(c chan<- int, q <-chan bool) {
	x, y := 1, 1
	for {
		select {
		case c <- x: // 未处理关闭时异常,因为下面退出时才会关闭
			x, y = y, x+y
		case <-q:
			close(c)
			return // 不使用break,因为它只能跳出select
		}
	}
}
func (f *Fibonacci) get(count int, c <-chan int, q chan<- bool, cb func([]int)) {
	s := make([]int, count, count) // 可省略第3个参数:容量(默认=第2个参数:长度)
	for i := 0; i < count; i++ {
		s[i] = <-c
	}
	q <- true
	close(q)
	cb(s)
}
