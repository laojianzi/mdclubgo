package conf

import (
	"os"
)

var (
	TestConf = os.Getenv("TEST_CONF")
)

// Build time and commit information.
//
// ⚠️ WARNING: should only be set by "-ldflags".
var (
	BuildTime   string
	BuildCommit string
)

// CustomConf returns the absolute path of custom configuration file that is used.
var CustomConf string

var (
	// App for project basic
	App = struct {
		Version string `ini:"-"`
		Name    string
		Debug   bool
	}{
		Version: "no version",
		Name:    "MDClubGo",
		Debug:   true,
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

// CacheOpts cache options
type CacheOpts struct {
	Type      string
	Host      string
	Namespace string
	Username  string
	Password  string
}

// Cache settings
var Cache CacheOpts

// EmailOpts email options
type EmailOpts struct {
	Type     string
	Host     string
	From     string
	Username string
	Password string
}

// Email settings
var Email EmailOpts

// StorageOpts storage options
type StorageOpts struct {
	Type string
}

// Storage settings
var Storage StorageOpts

// StorageLocalOpts storage local options
type StorageLocalOpts struct {
	URL string
}

// StorageLocal settings
var StorageLocal StorageLocalOpts
