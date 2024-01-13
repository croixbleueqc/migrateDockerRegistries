package main

import (
	"migrateDockerRegistries/cmd"
	"migrateDockerRegistries/env"
	"os"
	"path/filepath"
)

func main() {
	// Before anything else, we need to create the configs directory
	os.MkdirAll(filepath.Join(os.Getenv("HOME"), ".config", "JFG", "migrateDockerRegistries"), os.ModePerm)

	// Then we create a sample environment file
	//s := env.DockerRegistry{Direction: "source", Name: "Source registry example", URL: "myurl", Username: "myusername", Password: "my (encrypted) password"}
	//d := env.DockerRegistry{Direction: "destination", Name: "Destination registry example", URL: "myurl", Username: "myusername", Password: "my (encrypted) password"}
	s := env.DockerRegistry{Name: "Source registry example", URL: "myurl", Username: "myusername", Password: "my (encrypted) password"}
	d := env.DockerRegistry{Name: "Destination registry example", URL: "myurl", Username: "myusername", Password: "my (encrypted) password"}
	e := env.DockerRegistryCreds{Source: s, Dest: d}

	e.SaveEnvironmentFile("sample.json")

	// Finally we start the command loop
	cmd.Execute()
}
