package go_type

import (
	"angenalZZZ/go-program/api-models"
	"fmt"
)

// 类型检查
func TypeCheck() {
	var p = api_models.Point{}
	fmt.Println("-------------------------\n类型检查：")
	fmt.Println("  格式化p：%v %+v %T %#v")
	fmt.Printf("  格式化p：%v %+v %T %#v\n", p, p, p, p)
	fmt.Printf("  格式化i：%c %8.1f %8.2f %8x\n", 65, 12.5, 12.509, 54349)
}
