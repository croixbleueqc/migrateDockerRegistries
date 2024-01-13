// migrateDockerRegistries
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/auth/clientHelpers.go
// Original timestamp: 2024/01/13 12:31

package auth

import (
	"encoding/base64"
	"fmt"
	"github.com/docker/docker/client"
	"migrateDockerRegistries/env"
	"migrateDockerRegistries/helpers"
	"os"
)

func ClientConnect(uri string) *client.Client {
	cli, err := client.NewClientWithOpts(client.WithHost(uri), client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Printf("Unable to create docker auth: %s\n", err)
		os.Exit(-1)
	}

	return cli
}

func EncodeToken(regCreds env.DockerRegistry) string {
	return base64.StdEncoding.EncodeToString([]byte(regCreds.Username + ":" + helpers.DecodeString(regCreds.Password)))
}
