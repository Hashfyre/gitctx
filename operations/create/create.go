package create

import (
	"log"
	"os"
	"os/user"
	"path"

	"gopkg.in/ini.v1"
)

type ConfigLocation struct {
	Path string
}

type GitUserConfig struct {
	Name  string
	Email string
}

func DefaultConfig(name string) (*ConfigLocation, error) {
	user, err := user.Current()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &ConfigLocation{Path: path.Join(user.HomeDir, ".config", "gitctx", name)}, nil
}

func CreateConfig(config *GitUserConfig, location *ConfigLocation) error {
	err := os.MkdirAll(path.Dir(location.Path), os.ModeDir|0755)
	if err != nil {
		return err
	}

	err = createConfigFile(config, location)
	if err != nil {
		return err
	}

	return nil
}

func createConfigFile(config *GitUserConfig, location *ConfigLocation) error {
	cfg := ini.Empty()
	userSection := cfg.Section("user")
	userSection.Key("name").SetValue(config.Name)
	userSection.Key("email").SetValue(config.Email)

	return cfg.SaveTo(location.Path)
}
