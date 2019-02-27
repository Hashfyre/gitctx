package use

import (
	"errors"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/go-ini/ini"
	"github.com/jfornoff/gitctx/operations/create"
)

func UseConfig(configLocation create.ConfigLocation, projectDir string) error {
	currentPath := projectDir

	for descent := 1; descent <= 5; descent++ {
		log.Println("Checking", currentPath)

		if gitDirectoryPresent(currentPath) {
			log.Println("Found .git directory at", currentPath)
			targetPath, err := filepath.Abs(gitDirectoryPath(currentPath))
			if err != nil {
				return err
			}

			err = addConfigEntry(configLocation, targetPath)
			if err != nil {
				return err
			}

			return nil
		}

		currentPath = filepath.Join(currentPath, "..")
	}

	return errors.New("Did not find .git directory")
}

func gitDirectoryPresent(path string) bool {
	fileinfo, err := os.Stat(gitDirectoryPath(path))
	if err != nil {
		return false
	}

	return fileinfo.IsDir()
}

func gitDirectoryPath(path string) string {
	return filepath.Join(path, ".git")
}

func addConfigEntry(configLocation create.ConfigLocation, gitDirPath string) error {
	configFilePath := path.Join(gitDirPath, "config")
	log.Printf("Adding config entry to %v", configFilePath)
	cfg, err := ini.ShadowLoad(configFilePath)

	if err != nil {
		return err
	}

	cfg.Section("include").NewKey("path", configLocation.Path)

	return cfg.SaveTo(configFilePath)
}
