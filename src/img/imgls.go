// migrateDockerRegistries
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/img/imgls.go
// Original timestamp: 2024/01/13 09:55

package img

import (
	"context"
	//"encoding/base64"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/registry"
	"migrateDockerRegistries/auth"
	"migrateDockerRegistries/env"
	"migrateDockerRegistries/helpers"
	"sort"
)

func CompareImagesLists() error {
	var registries env.DockerRegistryCreds
	var err error
	var srcImgs, dstImgs []string

	if registries, err = env.LoadEnvironmentFile(); err != nil {
		return err
	}

	if srcImgs, err = getImages(registries.Source); err != nil {
		return helpers.CustomError{fmt.Sprintf("%s: %s\n", helpers.Red("Unable to list images in source: "), err.Error())}
	}
	if dstImgs, err = getImages(registries.Dest); err != nil {
		return helpers.CustomError{fmt.Sprintf("%s: %s\n", helpers.Red("Unable to list images in destination: "), err.Error())}
	}
	missingImgs := compareImagesList(srcImgs, dstImgs)

	fmt.Printf("Images present in %s but not in %s:\n",
		helpers.Blue(registries.Source.URL), helpers.Blue(registries.Dest.URL))
	for _, image := range missingImgs {
		fmt.Println(image)
	}

	return nil
}

// Lists images from given registry
func getImages(regInfo env.DockerRegistry) ([]string, error) {
	cli := auth.ClientConnect(regInfo.URL)
	authConfig := registry.AuthConfig{
		Username: regInfo.Username,
		Password: regInfo.Password,
	}
	cli.RegistryLogin(context.Background(), authConfig)

	// Fetch a list of images from the registry
	imageList, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		return nil, err
	}

	var images []string
	for _, image := range imageList {
		for _, tag := range image.RepoTags {
			images = append(images, tag)
		}
	}
	return images, nil
}

// Compares the lists from both registries
func compareImagesList(srcLst, dstLst []string) []string {
	sort.Strings(srcLst)
	sort.Strings(dstLst)

	var missingImages []string

	// Compare source and target lists
	for _, sourceImage := range srcLst {
		if !contains(dstLst, sourceImage) {
			missingImages = append(missingImages, sourceImage)
		}
	}
	return missingImages
}

// checks if a given element is part of the list
// not really efficient, especially for large lists
func contains(list []string, element string) bool {
	for _, e := range list {
		if e == element {
			return true
		}
	}
	return false
}
