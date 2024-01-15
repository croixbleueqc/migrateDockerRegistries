// migrateDockerRegistries
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/connection/authHelpers.go
// Original timestamp: 2024/01/15 15:48

package connection

import (
	"fmt"
	"migrateDockerRegistries/helpers"
	"strings"
)

var ConnectURI = ""

func buildConnectURI() string {
	if strings.HasPrefix(ConnectURI, "unix://") {
		return ConnectURI
	}
	if !strings.HasPrefix(ConnectURI, "tcp://") {
		if !strings.Contains(ConnectURI, ":") {
			ConnectURI = "tcp://" + ConnectURI + ":2375"
		} else {
			ConnectURI = "tcp://" + ConnectURI
		}
	} else {
		if !strings.Contains(ConnectURI, ":") {
			ConnectURI += ":2375"
		}
	}
	return ConnectURI
}

func showHost(uri string, showNow bool) string {
	//if uri == "" {
	//	uri = BuildConnectURI()
	//}
	if strings.HasPrefix(uri, "unix://") {
		uri = "localhost (unix socket)"
	}
	if showNow {
		fmt.Printf("\nDocker host is: %s.\n", helpers.White(uri))
	}
	return uri
}
