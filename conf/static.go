package conf

// CustomConf returns the absolute path of custom configuration file that is used.
var CustomConf string

var (
	// App for project basic
	App struct {
		Version string `ini:"-"`
		Name    string
		Debug   string
	}
)

// ServerOpts Server options
type ServerOpts struct {
	HTTPSEnable bool   `ini:"HTTPS_ENABLE"`
	HTTPAddr    string `ini:"HTTP_ADDR"`
	HTTPPort    string `ini:"HTTP_PORT"`
	CertFile    string
	KeyFile     string
}

// Server settings
var Server ServerOpts
