package config

// App 配置 变量
var appConfig *AppConfiguration

// App 配置 结构
type AppConfiguration struct {
	Server struct {
		ListenAddr string `default:""`
		Port       int    `default:"80"`
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
	PassStrength      int    `default:"10"`
	UploadedImagesDir string `default:"data/images"`
	PluginsDir        string `default:"data/plugins"`
}

// 获取 App 配置
func Get() *AppConfiguration {
	if appConfig == nil {
		appConfig, files := new(AppConfiguration), []string{"config.yml", "/etc/app/config.yml"}

		// 读取配置例子文件
		files = append(files, "config.example.yml")

		err := Config{&Configuration{EnvironmentPrefix: "GI"}}.Load(appConfig, files...)
		if err != nil {
			panic(err)
		}
	}
	return appConfig
}
