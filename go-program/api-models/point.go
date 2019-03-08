package api_models

///////////类型声明和定义///////////

type Point struct {
	x, y int
}

///////////函数声明和定义///////////

func (Point) New(x, y int) Point {
	return Point{x: x, y: y}
}
