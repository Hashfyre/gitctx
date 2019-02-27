package use

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/jfornoff/gitctx/operations/create"
	"gopkg.in/ini.v1"
)

func TestUse(t *testing.T) {
	setUpTestProject()
	defer resetTestProject()

	location := create.ConfigLocation{Path: "./test/testctx"}
	projectDir, err := filepath.Abs("./test")
	if err != nil {
		log.Fatal(err)
	}

	err = UseConfig(location, projectDir)
	if err != nil {
		log.Fatal(err)
	}

	cfg, err := ini.ShadowLoad("./test/.git/config")
	if err != nil {
		log.Fatal("Could not load git config for assertion!")
	}

	pathSet := cfg.Section("include").Key("path").ValueWithShadows()
	expectedPathSet := []string{location.Path}
	if !reflect.DeepEqual(pathSet, expectedPathSet) {
		t.Fatalf("Expected include path to be %#v, but was %#v", expectedPathSet, pathSet)
	}
}

func setUpTestProject() {
	copyFile("./test/.git/config", "./test/.git/backupconfig")
}

func resetTestProject() {
	copyFile("./test/.git/backupconfig", "./test/.git/config")

	err := os.Remove("./test/.git/backupconfig")
	if err != nil {
		log.Fatal(err)
	}
}

func copyFile(sourceFile string, destinationFile string) {
	input, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = ioutil.WriteFile(destinationFile, input, 0644)
	if err != nil {
		fmt.Println("Error creating", destinationFile)
		log.Fatal(err)
		return
	}
}
