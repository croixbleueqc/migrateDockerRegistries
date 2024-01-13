// migrateDockerRegistries
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/env/envHelpers.go

package env

import (
	"bufio"
	"encoding/json"
	"fmt"
	"migrateDockerRegistries/helpers"
	"os"
	"path/filepath"
	"strings"
)

var EnvConfigFile string

type DockerRegistry struct {
	Name     string `json:name"`
	URL      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type DockerRegistryCreds struct {
	Source DockerRegistry `json:"Source Docker Registry"`
	Dest   DockerRegistry `json:"Destination DockerRegistry"`
	//ALPINE []Repository `json:"APK"`
}

// Load the JSON environment file in the user's .config/certificatemanager directory, and store it into a data type (struct)
func LoadEnvironmentFile() (DockerRegistryCreds, error) {
	var payload DockerRegistryCreds
	var err error

	if !strings.HasSuffix(EnvConfigFile, ".json") {
		EnvConfigFile += ".json"
	}
	rcFile := filepath.Join(os.Getenv("HOME"), ".config", "JFG", "migrateDockerRegistries", EnvConfigFile)
	jFile, err := os.ReadFile(rcFile)
	if err != nil {
		return DockerRegistryCreds{}, err
	}
	err = json.Unmarshal(jFile, &payload)
	if err != nil {
		return DockerRegistryCreds{}, err
	} else {
		return payload, nil
	}
}

// Save the above structure into a JSON file in the user's .config/certificatemanager directory
func (e DockerRegistryCreds) SaveEnvironmentFile(outputfile string) error {
	if outputfile == "" {
		outputfile = EnvConfigFile
	}
	jStream, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return err
	}
	rcFile := filepath.Join(os.Getenv("HOME"), ".config", "JFG", "migrateDockerRegistries", outputfile)
	err = os.WriteFile(rcFile, jStream, 0600)

	return err
}

// We fetch the following info to populate the DockerRegistry structure
func fetchRepoInfo(prompt string) DockerRegistry {
	var repo DockerRegistry

	fmt.Printf("\nPlease enter the %s docker registry credentials\n\n", helpers.Blue(prompt))

	repo.Name = getStringVal("Please enter the friendly repo name (ENTER to quit): ")
	repo.URL = getStringVal("Please enter the repo URL: ")
	if !strings.HasSuffix(repo.URL, "/") {
		repo.URL += "/"
	}
	repo.Username = getStringVal("Please enter the username needed to login: ")
	repo.Password = helpers.EncodeString(helpers.GetPassword("Please enter that user's password: "))

	return repo
}

func getStringVal(prompt string) string {
	fmt.Print(prompt)
	inputVal := bufio.NewReader(os.Stdin)
	input, _ := inputVal.ReadString('\n')

	return strings.TrimSpace(input)
}
