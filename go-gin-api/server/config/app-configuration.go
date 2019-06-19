package config

// App 配置 变量
var appConfig *AppConfiguration

// App 配置 结构
type AppConfiguration struct {
	Server struct {
		ListenAddr string `default:"" env:"GI_API_ADDR"`
		Port       int    `default:"80" env:"GI_API_PORT" required:"true"`
		SSL        struct {
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
		ResponseHeaders map[string]string
		Stream          struct {
			AllowedOrigins []string
		}
	}
	Database struct {
		Dialect    string `default:"sqlite3"`
		Connection string `default:"data/app.db"`
	}
	DefaultUser struct {
		Name string `default:"admin"`
		Pass string `default:"admin"`
	}
	UploadedImagesDir string `default:"data/images"`
	PluginsDir        string `default:"data/plugins"`
}

// 获取 App 配置
func Get() *AppConfiguration {
	if appConfig == nil {
		config := Config{&Configuration{EnvironmentPrefix: "GI"}}
		appConfig, files := new(AppConfiguration), []string{"config.yml", "/etc/app/config.yml"}

		// 读取配置例子文件
		files = append(files, "config.example.yml")

		err := config.Load(appConfig, files...)
		if err != nil {
			panic(err)
		}
	}
	return appConfig
}
