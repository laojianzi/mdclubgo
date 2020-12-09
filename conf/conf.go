package conf

//go:generate go-bindata -nomemcopy -pkg=conf -ignore="\\.DS_Store|README.md|TRANSLATORS" -prefix=../ -debug=false -o=conf_gen.go ../conf/app.ini

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/labstack/gommon/color"
	"gopkg.in/ini.v1"

	"github.com/laojianzi/mdclubgo/log"
)

// Source is the configuration object.
var Source *ini.File

// Init initializes configuration from custom/conf/app.ini
func Init(customConf string) error {
	log.Init(App.Name, Log.RootPath, App.Debug)

	var err error
	defer func() {
		if err == nil {
			log.Init(App.Name, Log.RootPath, App.Debug)
		}

		log.Info(color.Green(fmt.Sprintf("%s %s", App.Name, App.Version)))
		log.Info(color.Green(fmt.Sprintf("Work directory: %s", WorkDir())))
		log.Info(color.Green(fmt.Sprintf("Custom path: %s", CustomDir())))
		log.Info(color.Green(fmt.Sprintf("Custom config: %s", CustomConf)))
		log.Info(color.Green(fmt.Sprintf("Log path: %s", Log.RootPath)))
		log.Info(color.Green(fmt.Sprintf("Email Use: %s", Email.Type)))
		log.Info(color.Green(fmt.Sprintf("Build time: %s", BuildTime)))
		log.Info(color.Green(fmt.Sprintf("Build commit: %s", BuildCommit)))
	}()

	Source, err = ini.LoadSources(ini.LoadOptions{
		IgnoreInlineComment: true,
	}, MustAsset("conf/app.ini"))
	if err != nil {
		return fmt.Errorf("parse 'conf/app.ini': %w", err)
	}

	Source.NameMapper = ini.SnackCase
	customConf, err = filepath.Abs(customConf)
	if err != nil {
		return fmt.Errorf("get absolute path: %w", err)
	}

	CustomConf = customConf
	file, err := os.Stat(customConf)
	if err != nil || file.IsDir() {
		log.Warn("Custom config %q not found. Ignore this warning if you're running for the first time", customConf)
	} else if err = Source.Append(customConf); err != nil {
		return fmt.Errorf("append %q: %w", customConf, err)
	}

	if err = Source.Section(ini.DefaultSection).MapTo(&App); err != nil {
		return fmt.Errorf("mapping default section: %w", err)
	}

	// server settings
	if err = Source.Section("server").MapTo(&Server); err != nil {
		return fmt.Errorf("mapping [server] section: %w", err)
	}

	// log settings
	if err = Source.Section("log").MapTo(&Log); err != nil {
		return fmt.Errorf("mapping [log] section: %w", err)
	}

	if Log.RootPath == "" {
		Log.RootPath = filepath.Join(WorkDir(), "tmp", "mdclubgo.log")
	}

	if Log.RootPath[0] != '/' {
		Log.RootPath = filepath.Join(WorkDir(), Log.RootPath)
	}

	// database settings
	if err = Source.Section("database").MapTo(&Database); err != nil {
		return fmt.Errorf("mapping [database] section: %w", err)
	}

	// cache settings
	if err = Source.Section("cache").MapTo(&Cache); err != nil {
		return fmt.Errorf("mapping [cache] section: %w", err)
	}

	// email settings
	if err = Source.Section("email").MapTo(&Email); err != nil {
		return fmt.Errorf("mapping [cache] section: %w", err)
	}

	// storage settings
	if err = Source.Section("storage").MapTo(&Storage); err != nil {
		return fmt.Errorf("mapping [storage] section: %w", err)
	}

	var storageV interface{}
	switch Storage.Type {
	case "local":
		storageV = &StorageLocal
	default:
		return fmt.Errorf("[storage] TYPE is invalid")
	}

	if err = Source.Section("storage.local").MapTo(storageV); err != nil {
		return fmt.Errorf("mapping [storage.local] section: %w", err)
	}

	return nil
}
