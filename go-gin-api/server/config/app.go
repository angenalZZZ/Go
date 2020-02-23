package config

import "github.com/lexkong/log"

// Must panic if err.
func Must(err error) {
	if err != nil {
		panic(err)
	}
}

// App 配置 结构
type AppConfigModel struct {
	// App名称
	AppName string `default:"API"`
	AppMode string `default:"debug"` //debug or release
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
		// 响应超时
		ReadTimeout  int `default:"120"`
		WriteTimeout int `default:"120"`
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
	// 链路追踪(uber/jaeger)
	Tracker struct {
		Enabled   *bool  `default:"true"`
		ServeAddr string `default:"127.0.0.1:6831"`
	}
	// 系统告警
	NotifyUser struct {
		Enabled *bool `default:"false"`
		// 告警接收人
		ErrorNotifyUser string `default:"angenal2008@163.com"`
		// 邮箱服务器信息
		Smtp struct {
			Port int    `default:"465"`
			Host string `default:"smtp.163.com"`
			User string `default:"angenal@163.com"`
			Pass string `default:""` // 密码或授权码
		}
	}
	// 上传图片目录
	UploadedImagesDir string `default:"data/images"`
	// 插件目录
	PluginsDir string `default:"data/plugins"`
	// 日志跟踪
	Log log.PassLagerCfg
}
