package conf

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/ini.v1"
)

// Source is the configuration object.
var Source *ini.File

// Init initializes configuration from custom/conf/app.ini
func Init() error {
	content, err := ioutil.ReadFile("conf/app.ini")
	if err != nil {
		panic("can't read 'conf/app.ini'")
	}

	Source, err = ini.LoadSources(ini.LoadOptions{
		IgnoreInlineComment: true,
	}, content)
	if err != nil {
		return fmt.Errorf("parse 'conf/app.ini': %w", err)
	}

	Source.NameMapper = ini.SnackCase
	customConf := filepath.Join(CustomDir(), "conf", "app.ini")
	customConf, err = filepath.Abs(customConf)
	if err != nil {
		return fmt.Errorf("get absolute path: %w", err)
	}

	CustomConf = customConf

	file, err := os.Stat(customConf)
	if err != nil || file.IsDir() {
		log.Printf("Custom config %q not found. Ignore this warning if you're running for the first time", customConf)
	} else if err = Source.Append(customConf); err != nil {
		return fmt.Errorf("append %q: %w", customConf, err)
	}

	if err = Source.Section(ini.DefaultSection).MapTo(&App); err != nil {
		return fmt.Errorf("mapping default section: %w", err)
	}

	// server settings
	if err = Source.Section("server").MapTo(&Server); err != nil {
		return fmt.Errorf("mapping [server[ section: %w", err)
	}

	return nil
}
