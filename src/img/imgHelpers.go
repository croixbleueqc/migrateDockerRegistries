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
	"path/filepath"
	"strings"
)

type RepoTagStruct struct {
	RepoName string `json:"reponame"`
	ImageTag string `json:"imagetag"`
}

var Retag, Push, Delete, LatestOnly bool

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

		if LatestOnly {
			if containsTag(tags["tags"].([]interface{}), "latest") {
				repoTags = append(repoTags, fmt.Sprintf("%s:latest", repo))
			}
		} else {
			// Include all tags
			for _, tag := range tags["tags"].([]interface{}) {
				repoTags = append(repoTags, fmt.Sprintf("%s:%s", repo, tag.(string)))
			}
		}
	}

	return repoTags, nil
}

func compareLists(url string, orgList, dstList []string) ([]string, []RepoTagStruct) {
	var finalList []string
	var finalrepotagStruct []RepoTagStruct
	//var finalrepotag RepoTagStruct

	// Identify repo+tags in orgList but not in dstList
	for _, repoTag := range orgList {
		if !containsList(dstList, repoTag) {
			repo := stripProtocol(url) + repoTag
			finalrepotag := RepoTagStruct{RepoName: url, ImageTag: repoTag}
			finalrepotagStruct = append(finalrepotagStruct, finalrepotag)
			finalList = append(finalList, repo)
		}
	}

	return finalList, finalrepotagStruct
}

func containsList(list []string, item string) bool {
	for _, val := range list {
		if val == item {
			return true
		}
	}
	return false
}

func containsTag(list []interface{}, item string) bool {
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

func saveJSON(filename string, repotaglistJson []RepoTagStruct) error {
	jStream, err := json.MarshalIndent(repotaglistJson, "", "  ")
	if err != nil {
		return err
	}
	rcFile := filepath.Join(filename)
	err = os.WriteFile(rcFile, jStream, 0600)

	return nil
}
