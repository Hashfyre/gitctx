package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"path/filepath"
	"reflect"

	"github.com/jfornoff/gitctx/operations/create"
	"github.com/jfornoff/gitctx/operations/use"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var useCmd = &cobra.Command{
	Use:   "use",
	Short: "Use a Git context in a repository",
	Long: `Allows using a Git context (i.e., username and email for committing) in a local Git repository.
  This is done by adding an [include] path in the repository-local Git config (.git/config).
  `,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := selectConfig()
		if err != nil {
			log.Fatal(err)
		}

		use.UseConfig(*config, ".")
	},
}

func selectConfig() (*create.ConfigLocation, error) {
	configDir, err := create.DefaultConfigDirectory()

	if err != nil {
		return nil, err
	}

	availableConfigs, err := listAvailableConfigs(*configDir)

	if err != nil {
		return nil, err
	}

	if len(availableConfigs) == 0 {
		return nil, fmt.Errorf("No configs found in %v", *configDir)
	}

	configsByName := make(map[string]create.ConfigLocation)

	for _, config := range availableConfigs {
		configsByName[filepath.Base(config.Path)] = config
	}

	prompt := promptui.Select{
		Label: "Which context to use?",
		Items: reflect.ValueOf(configsByName).MapKeys(),
	}

	_, configName, err := prompt.Run()

	if err != nil {
		return nil, err
	}

	config, exists := configsByName[configName]

	if !exists {
		return nil, errors.New("bug in config lookup, tried to fetch config name that is not there")
	}

	log.Printf("Selected config location %v", config)
	return &config, nil
}

func listAvailableConfigs(configDir string) ([]create.ConfigLocation, error) {
	files, err := ioutil.ReadDir(configDir)
	if err != nil {
		return nil, err
	}

	result := make([]create.ConfigLocation, len(files))
	for i, fileInfo := range files {
		result[i] = create.ConfigLocation{Path: path.Join(configDir, fileInfo.Name())}
	}

	return result, nil
}

func init() {
	rootCmd.AddCommand(useCmd)
}
