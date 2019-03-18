package api_models

import "fmt"

///////////类型声明和接口定义///////////

type Point struct {
	X, Y int
	IPoint
}

type IPoint interface {
	Sum() int
	Complex() complex128
	RealX() int
	ImagY() int
	Value() string
	Values() [2]int
	String() string
}

///////////函数声明和定义///////////

func (p *Point) Sum() int {
	return p.X + p.Y
}

func (p *Point) Complex() complex128 {
	return complex(float64(p.X), float64(p.Y))
}
func (p *Point) RealX() int {
	return int(real(p.Complex()))
}
func (p *Point) ImagY() int {
	return int(imag(p.Complex()))
}

func (p *Point) Value() string {
	return fmt.Sprintf("{X: %d, Y: %d}", p.X, p.Y)
}
func (p *Point) Values() [2]int {
	return [2]int{p.X, p.Y}
}
func (p *Point) String() string {
	return fmt.Sprintf("{X: %d, Y: %d} %p", p.X, p.Y, p)
}
