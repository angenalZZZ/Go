package go_file

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/rs/xid"
)

// 文件管理：创建文件
func TestCreateFile() {
	fmt.Println("-------------------------\n文件管理：创建文件")

	tmpDir, e := ioutil.TempDir(os.TempDir(), "")
	if e != nil {
		fmt.Fprintf(os.Stderr, "  创建tmp目录异常：%v", e)
		panic(e) // 无法创建目录时，进程退出code:2
	}

	f, e := os.Create(filepath.Join(tmpDir, fmt.Sprintf("_%s.tmp", xid.New().String()[12:])))
	//f, e := ioutil.TempFile("", "*.tmp") // dir = os.TempDir() 等同于上面>创建tmp文件
	if e != nil {
		fmt.Fprintf(os.Stderr, "  创建tmp文件异常：%v", e)
		panic(e) // 无法创建文件时，进程退出code:2
	}
	defer os.RemoveAll(tmpDir) // 不删除也可以，因为系统会定期清理

	fmt.Printf("  创建tmp文件成功：%s \n", f.Name())
	var fileContent = "文件内容\r\n内容..."
	defer os.Remove(f.Name())    // defer 当要退出func main()作用域时执行
	defer f.Close()              // defer 在作用域内 按倒序执行
	fmt.Fprintln(f, fileContent) // 写入文件内容
	fmt.Println("  写入tmp文件成功.", "删除tmp文件成功.")
}
