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
