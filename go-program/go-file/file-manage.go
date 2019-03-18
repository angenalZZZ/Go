package go_file

import (
	"fmt"
	"os"
	"time"
)

// 文件管理：创建文件
func CreateFile() {
	fmt.Println("-------------------------\n文件管理：创建文件")

	fName := fmt.Sprintf("%s\\~%d.tmp", os.TempDir(), time.Now().Unix())
	f, e := os.Create(fName)
	if e != nil {
		fmt.Fprintf(os.Stderr, "  创建tmp文件异常：%+v", e)
		panic(e) // 无法创建文件时，进程退出code:2
	} else {
		fmt.Printf("  创建tmp文件成功：%s \n", f.Name())
		var fileContent = "文件内容\r\n内容..."
		defer os.Remove(f.Name())    // defer 当要退出func main()作用域时执行
		defer f.Close()              // defer 在作用域内 按倒序执行
		fmt.Fprintln(f, fileContent) // 写入文件内容
		fmt.Println("  写入tmp文件成功.", "删除tmp文件成功.")
	}
}
