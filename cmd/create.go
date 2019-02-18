package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/jfornoff/gitctx/operations/create"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new Git context.",
	Long:  `Adds a new config file in ~/.config/gitctx that can be included in project gitconfig-Files.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctxName := promptNonBlank("Context name")
		name := promptNonBlank("User name")
		email := promptNonBlank("Email")

		locationConfig, err := create.DefaultConfig(ctxName)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		create.CreateConfig(&create.GitUserConfig{Name: name, Email: email}, locationConfig)
		println(fmt.Sprintf("Created config entry in %v!", locationConfig.Path))
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}

func promptNonBlank(label string) string {
	validate := func(input string) error {
		if input == "" {
			return errors.New("Cannot be blank")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    label,
		Validate: validate,
	}

	input, err := prompt.Run()
	if err != nil {
		log.Fatal(fmt.Sprintf("Prompt failed %v", err))

		os.Exit(1)
	}

	return input
}
