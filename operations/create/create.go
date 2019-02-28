package create

import (
	"os"
	"os/user"
	"path"

	"gopkg.in/ini.v1"
)

// ConfigLocation : Where to store a context configuration
type ConfigLocation struct {
	Path string
}

// GitUserConfig : Contents of a context configuration
type GitUserConfig struct {
	Name  string
	Email string
}

func DefaultConfigDirectory() (*string, error) {
	user, err := user.Current()

	if err != nil {
		return nil, err
	}

	result := path.Join(user.HomeDir, ".config", "gitctx")
	return &result, nil
}

func DefaultConfigLocation(name string) (*ConfigLocation, error) {
	defaultConfigDirectory, err := DefaultConfigDirectory()

	if err != nil {
		return nil, err
	}

	return &ConfigLocation{Path: path.Join(*defaultConfigDirectory, name)}, nil
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
