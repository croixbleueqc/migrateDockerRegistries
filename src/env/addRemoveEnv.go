// nxrmuploader
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/env/addRemoveEnv.go
// Original timestamp: 2023/12/31 14:49

package env

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Remove an environment file
func RemoveEnvFile(envfiles []string) error {

	for _, envfile := range envfiles {
		if !strings.HasSuffix(envfile, ".json") {
			envfile += ".json"
		}
		if err := os.Remove(filepath.Join(os.Getenv("HOME"), ".config", "JFG", "migrateDockerRegistries", envfile)); err != nil {
			return err
		}
		if err := os.Remove(filepath.Join(os.Getenv("HOME"), ".config", "JFG", "migrateDockerRegistries", envfile)); err != nil {
			return err
		}
		fmt.Printf("%s removed succesfully\n", envfile)
	}
	return nil
}

// Create an environment file, you will be prompted for the necessary values
func AddEnvFile(envfile string) error {
	var env DockerRegistryCreds
	var err error

	if envfile == "" {
		envfile = EnvConfigFile
	}
	if !strings.HasSuffix(envfile, ".json") {
		envfile += ".json"
	}

	env = prompt4EnvironmentValues()

	if err = env.SaveEnvironmentFile(envfile); err != nil {
		return err
	}
	return nil
}

func prompt4EnvironmentValues() DockerRegistryCreds {
	var env DockerRegistryCreds
	var source, destination DockerRegistry

	// Fetch source repo info
	source = fetchRepoInfo("source")
	//source.Direction = "source"

	// Fetch dest repo info
	destination = fetchRepoInfo("destination")
	//destination.Direction = "destination"

	// We now add those repos into the super-struct
	env.Source = source
	env.Dest = destination
	return env
}
