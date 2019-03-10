package api_models

///////////类型声明和定义///////////

type Point struct {
	X, Y int
	IPoint
}

type IPoint interface {
	Add(x, y int)
}

///////////函数声明和定义///////////

func (Point) New(x, y int) Point {
	return Point{X: x, Y: y}
}
