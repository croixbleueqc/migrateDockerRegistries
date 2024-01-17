// migrateDockerRegistries
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/img/listImages.go
// Original timestamp: 2024/01/17 08:47

package img

import (
	"fmt"
	"migrateDockerRegistries/env"
	"migrateDockerRegistries/helpers"
	"os"
	"strings"
)

func CompareImagesLists() error {
	var registries env.DockerRegistryCreds
	var err error
	var orgList, destList []string

	// load config from environment file
	if registries, err = env.LoadEnvironmentFile(); err != nil {
		return err
	}

	// Fetches the source registry's list
	orgList, err = getRepoTags(registries.Source.URL)
	if err != nil {
		return helpers.CustomError{fmt.Sprintf("Unable to fetch %s repo/tags list: ",
			registries.Source.URL, err)}
	}

	// Fetches the dest registry's list
	destList, err = getRepoTags(registries.Dest.URL)
	if err != nil {
		return helpers.CustomError{fmt.Sprintf("Unable to fetch %s repo/tags list: ",
			registries.Dest.URL, err)}
	}
	// Save both lists to files
	if err = saveListToFile(registries.Source.Name+".txt", orgList); err != nil {
		return helpers.CustomError{fmt.Sprintf("Unable to save the source registry's list: %s", err)}
	}
	if err = saveListToFile(registries.Dest.Name+".txt", destList); err != nil {
		return helpers.CustomError{fmt.Sprintf("Unable to save the destination registry's list: %s", err)}
	}
	finalList := compareLists(orgList, destList)
	if err = saveListToFile(registries.Source.Name+"-"+registries.Dest.Name+".txt", finalList); err != nil {
		return helpers.CustomError{fmt.Sprintf("Unable to save the final list: %s", err)}
	}
	return nil
}

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

func compareLists(orgList, dstList []string) []string {
	var finalList []string

	// Identify repo+tags in orgList but not in dstList
	for _, repoTag := range orgList {
		if !contains(dstList, repoTag) {
			finalList = append(finalList, repoTag)
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
