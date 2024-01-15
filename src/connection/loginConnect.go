// migrateDockerRegistries
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/connection/loginConnect.go
// Original timestamp: 2024/01/13 12:31

package connection

import (
	"encoding/base64"
	"fmt"
	"github.com/docker/docker/client"
	"migrateDockerRegistries/env"
	"migrateDockerRegistries/helpers"
	"os"
)

func ClientConnect(showHostInfo bool) *client.Client {
	uri := buildConnectURI()
	cli, err := client.NewClientWithOpts(client.WithHost(uri), client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Printf("Unable to create docker client: %s\n", err)
		os.Exit(-1)
	}

	if showHostInfo {
		showHost(uri, showHostInfo)
	}
	return cli
}

func EncodeToken(regCreds env.DockerRegistry) string {
	return base64.StdEncoding.EncodeToString([]byte(regCreds.Username + ":" + helpers.DecodeString(regCreds.Password)))
}
