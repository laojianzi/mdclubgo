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
	HTTPSEnable              bool   `ini:"HTTPS_ENABLE"`
	HTTPAddr                 string `ini:"HTTP_ADDR"`
	HTTPPort                 string `ini:"HTTP_PORT"`
	CertFile                 string
	KeyFile                  string
	AccessControlAllowOrigin string `ini:"ACCESS_CONTROL_ALLOW_ORIGIN"`
}

// Server settings
var Server ServerOpts

// LogOpts log options
type LogOpts struct {
	RootPath string `ini:"ROOT_PATH"`
}

// Log settings
var Log LogOpts

// DatabaseOpts db options
type DatabaseOpts struct {
	Type         string
	Host         string
	Name         string
	User         string
	Password     string
	SSLMode      string `ini:"SSL_MODE"`
	Path         string
	MaxOpenConns int
	MaxIdleConns int
}

// Database settings
var Database DatabaseOpts

// Indicates which database backend is currently being used.
var (
	UseSQLite3    bool
	UseMySQL      bool
	UsePostgreSQL bool
	UseMSSQL      bool
)
