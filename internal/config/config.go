package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

const (
	ConfigFileName   = "config"
	ConfigFileType   = "yaml"
	ConfigHomeSubdir = ".doprax"
)

// TODO temporary solution until upstream https://github.com/spf13/viper/issues/433 is fixed
func WriteConfigFile() error {
	cf := viper.ConfigFileUsed()

	if cf == "" {
		fullname := ConfigFileName + "." + ConfigFileType

		configDir, err := GetConfigDir()
		if err != nil {
			return errors.New("Failed to acquire config directory name")
		}

		cf = filepath.Join(configDir, fullname)
		err = os.MkdirAll(configDir, os.ModePerm)
		if err != nil {
			return err
		}
		fmt.Printf("Doprax configuration file created at %s\n", cf)
	}
	if err := viper.WriteConfigAs(cf); err != nil {
		return err
	}
	return nil
}

func GetConfigDir() (string, error) {
	dirname, err := os.UserHomeDir()
	return filepath.Join(dirname, ConfigHomeSubdir), err
}
