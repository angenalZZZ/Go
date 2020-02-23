package config

import "github.com/lexkong/log"

// App 配置 结构
type Config struct {
	// WebApi服务
	Server struct {
		ListenAddr string `default:"" env:"API_SERVER_ADDR"`
		Port       int    `default:"80" env:"API_SERVER_PORT" required:"true"`
		// HTTPS设置
		SSL struct {
			Enabled         *bool  `default:"false"`
			RedirectToHTTPS *bool  `default:"true"`
			ListenAddr      string `default:""`
			Port            int    `default:"443"`
			CertFile        string `default:""`
			CertKey         string `default:""`
			LetsEncrypt     struct {
				Enabled   *bool  `default:"false"`
				AcceptTOS *bool  `default:"false"`
				Cache     string `default:"data/certs"`
				Hosts     []string
			}
		}
		// 输出响应头
		ResponseHeaders map[string]string
		// 跨域访问限制
		Stream struct {
			AllowedOrigins []string
		}
	}
	// 连接数据库
	Database struct {
		Dialect    string `default:"sqlite3"`
		Connection string `default:"data/app.db"`
	}
	// 系统默认账号(内置账号)
	DefaultUser struct {
		Name string `default:"admin"`
		Pass string `default:"admin"`
	}
	// 链路追踪
	Tracker struct {
		Enabled    *bool  `default:"false"`
		JaegerAddr string `default:""`
	}
	// 上传图片目录
	UploadedImagesDir string `default:"data/images"`
	// 插件目录
	PluginsDir string `default:"data/plugins"`
	// 日志跟踪
	Log log.PassLagerCfg
}
