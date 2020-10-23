package conf

import (
	"os"
	"os/exec"
	"path/filepath"
	"sync"
)

var (
	appPath     string
	appPathOnce sync.Once
)

// AppPath returns the absolute path of the application's binary
func AppPath() string {
	appPathOnce.Do(func() {
		var err error
		appPath, err = exec.LookPath(os.Args[0])
		if err != nil {
			panic("look executable path: " + err.Error())
		}

		appPath, err = filepath.Abs(appPath)
		if err != nil {
			panic("get absolute executable path: " + err.Error())
		}
	})

	return appPath
}

var (
	workDir     string
	workDirOnce sync.Once
)

// WorkDir returns a work directory
func WorkDir() string {
	workDirOnce.Do(func() {
		workDir = os.Getenv("MDCLUBGO_WORK_DIR")
		if workDir != "" {
			return
		}

		workDir = filepath.Dir(AppPath())
	})

	return workDir
}

var (
	customDir     string
	customDirOnce sync.Once
)

// CustomDir returns a custom directory
func CustomDir() string {
	customDirOnce.Do(func() {
		customDir = filepath.Join(WorkDir(), "custom")
	})

	return customDir
}
