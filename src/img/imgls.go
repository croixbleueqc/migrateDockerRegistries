// migrateDockerRegistries
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/img/imgls.go
// Original timestamp: 2024/01/13 09:55

package img

import (
	"context"
	"encoding/json"
	"github.com/docker/docker/api/types/registry"
	"io/ioutil"
	"net/http"

	//"encoding/base64"
	"fmt"
	"migrateDockerRegistries/connection"
	"migrateDockerRegistries/env"
	"migrateDockerRegistries/helpers"
)

func CompareImagesLists() error {
	var registries env.DockerRegistryCreds
	var err error
	var srcImgs, dstImgs []string

	// load config from environment file
	if registries, err = env.LoadEnvironmentFile(); err != nil {
		return err
	}

	// fetches all images+tags from source registry
	if srcImgs, err = listImages(registries.Source); err != nil {
		return helpers.CustomError{fmt.Sprintf("%s: %s\n", helpers.Red("Unable to list images in source: "), err.Error())}
	}

	// fetches all images+tags from destination registry
	if dstImgs, err = listImages(registries.Dest); err != nil {
		return helpers.CustomError{fmt.Sprintf("%s: %s\n", helpers.Red("Unable to list images in destination: "), err.Error())}
	}

	// find images that are in first registry but not in second
	//missingImgs := compareImagesList(srcImgs, dstImgs)

	// for now: simply print the results
	fmt.Printf("Images present in %s but not in %s:\n",
		helpers.Blue(registries.Source.URL), helpers.Blue(registries.Dest.URL))
	for _, image := range missingImgs {
		fmt.Println(image)
	}
	return nil
}

// Lists images from given registry
func listImages(regInfo env.DockerRegistry) ([]Repository, error) {
	cli := connection.ClientConnect(false)
	authConfig := registry.AuthConfig{
		Username:      regInfo.Username,
		Password:      helpers.DecodeString(regInfo.Password),
		ServerAddress: regInfo.URL,
	}
	authToken := connection.EncodeToken(regInfo)

	cli.RegistryLogin(context.Background(), authConfig)
	cli.NegotiateAPIVersion(context.Background())

	// build the http request, with Authorization http header
	repoListURL := regInfo.URL + "/v2/_catalog"
	req, err := http.NewRequest("GET", repoListURL, nil)
	if err != nil {
		return nil, helpers.CustomError{fmt.Sprintf("Error fetching list from %s: %s", regInfo.URL, err)}
	}
	req.Header.Set("Authorization", "Basic "+authToken)

	//
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Decode response: JSON->repoList
	var repoList RepositoryList
	body, err := os.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &repoList)
	if err != nil {
		return nil, err
	}

	// Fetch tags for each repository
	for i, repo := range repoList.Repositories {
		tags, err := fetchTags(regInfo.URL, repo, authToken)
		if err != nil {
			return nil, err
		}
		repoList.Repositories[i].Tags = tags
	}

	return repoList.Repositories, nil
}

func fetchTags(registryURL, repository, encodedAuth string) ([]string, error) {
	//ctx := context.Background()

	// Make a request to list tags for a repository
	tagsListURL := registryURL + "/v2/" + repository + "/tags/list"
	req, err := http.NewRequest("GET", tagsListURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Basic "+encodedAuth)

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Decode the response into a list of tags
	var tagsList TagsList
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &tagsList)
	if err != nil {
		return nil, err
	}

	return tagsList.Tags, nil
}
