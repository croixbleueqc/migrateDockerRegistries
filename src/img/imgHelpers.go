// migrateDockerRegistries
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/img/imgHelpers.go
// Original timestamp: 2024/01/15 16:28

package img

import (
	"encoding/json"
	"fmt"
	"io"
	"migrateDockerRegistries/helpers"
	"net/http"
	"os"
	"strings"
)

// fetchJSON() : generic function used to either pick the image list or an image available tag
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

// getRepoTags() : fetch all images, and then for each image, all of its tags
func getRepoTags(registryURL string) ([]string, error) {
	// Fetch list of repo+tags from the registry
	catalogPath := "/v2/_catalog"
	repos, err := fetchJSON(registryURL + catalogPath)
	if err != nil {
		return nil, err
	}

	var repoTags []string

	// Fetch tags for each repository
	for _, repo := range repos["repositories"].([]interface{}) {
		tagsPath := fmt.Sprintf("/v2/%s/tags/list", repo.(string))
		tags, err := fetchJSON(registryURL + tagsPath)
		if err != nil {
			return nil, err
		}

		// Construct repo+tags and add to the list
		for _, tag := range tags["tags"].([]interface{}) {
			repoTags = append(repoTags, fmt.Sprintf("%s:%s", repo, tag.(string)))
		}
	}

	return repoTags, nil
}

func compareLists(url string, orgList, dstList []string) []string {
	var finalList []string

	// Identify repo+tags in orgList but not in dstList
	for _, repoTag := range orgList {
		if !contains(dstList, repoTag) {
			repo := stripProtocol(url) + repoTag
			finalList = append(finalList, repo)
		}
	}

	return finalList
}

func contains(list []string, item string) bool {
	for _, val := range list {
		if val == item {
			return true
		}
	}
	return false
}

func saveListToFile(filename string, list []string) error {
	data := []byte(strings.Join(list, "\n"))
	err := os.WriteFile(filename, data, 0644)

	return err
}

func stripProtocol(url string) string {
	// The following might seem nonsensical, but is needed to ensure we have a well-formatted URL
	url = strings.TrimPrefix(url, "https://")
	url = strings.TrimPrefix(url, "http://")
	url = strings.TrimSuffix(url, "/")
	return url + "/"
}
