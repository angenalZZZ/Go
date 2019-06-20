package app

// App 配置 结构
type Config struct {
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
