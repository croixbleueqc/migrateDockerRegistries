// migrateDockerRegistries
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/img/imgHelpers.go
// Original timestamp: 2024/01/15 16:28

package img

import (
	"encoding/json"
	"io"
	"migrateDockerRegistries/helpers"
	"net/http"
)

// RepositoryList represents the response from the /v2/_catalog endpoint
type RepositoryList struct {
	Repositories []Repository `json:"repositories"`
}

// Repository represents a Docker repository
type Repository struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

// TagsList represents the response from the /v2/<repository>/tags/list endpoint
type TagsList struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

func fetchJSON(url string) (map[string]interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, helpers.CustomError{"Error getting JSON payload from url endpoint: " + err.Error()}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, helpers.CustomError{"Error reading JSON payload from url endpoint: " + err.Error()}
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
