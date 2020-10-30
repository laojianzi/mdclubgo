package conf

var (
	// App for project basic
	App struct {
		Version string `ini:"-"`
		Name    string
		Debug   bool
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

// LogOpts log options
type LogOpts struct {
	RootPath string `ini:"ROOT_PATH"`
}

// Log settings
var Log LogOpts
