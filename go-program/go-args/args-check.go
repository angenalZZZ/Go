package go_args

import (
	"fmt"
	"os"
	"strings"
)

// 命令行参数
func ArgsCheck() {
	fmt.Println("-------------------------\n命令行参数：")
	//println("os.Args:", strings.Join(os.Args[1:], " "))
	fmt.Println("  os.Args:", os.Args[1:])

	// 循环并转换os.Args
	//var args = []string(nil)
	args := make(map[string]string)
	for _, v := range os.Args[1:] {
		i := strings.Index(v, "=")
		if i > 0 {
			//args = append(args, v[i+1:])
			args[strings.Trim(v[0:i], "-")] = v[i+1:]
		}
	}
	fmt.Println("  转换后的args:", args)
}
