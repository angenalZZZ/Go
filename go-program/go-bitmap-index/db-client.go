package go_bitmap_index

import "github.com/pilosa/go-pilosa"

var Client *pilosa.Client // 数据库连接客户端

func init() {
	// 配置数据库连接客户端
	Client = pilosa.DefaultClient()
}
