package create

import "testing"
import "os"
import "log"
import "fmt"
import "path"
import "io/ioutil"
import "strings"

func TestCreateConfig(t *testing.T) {
	location := ConfigLocation{Path: "./test/config_fixtures/gitconfig"}
	config := GitUserConfig{Name: "Testuser", Email: "test@example.com"}

	err := CreateConfig(&config, &location)
	defer cleanUpDirectory(location)

	if err != nil {
		t.Fatal(fmt.Sprintf("Unable to create config directory: %v", err))
	}

	// Check files are present and containing the right information
	configFileContents, err := ioutil.ReadFile(location.Path)
	if err != nil {
		t.Fatal(fmt.Sprintf("Unable to read config file: %v", err))
	}

	got := strings.TrimSpace(string(configFileContents))

	expected := strings.TrimSpace(`
[user]
name  = Testuser
email = test@example.com
	`)

	if expected != got {
		t.Fatal(fmt.Sprintf("\nGot:\n%v \n\nExpected:\n%v", got, expected))
	}
}

func cleanUpDirectory(location ConfigLocation) {
	err := os.RemoveAll(path.Dir(location.Path))
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to remove test directory: %v", err))
	}
}
