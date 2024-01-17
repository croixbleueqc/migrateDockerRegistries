// migrateDockerRegistries
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/img/listImages.go
// Original timestamp: 2024/01/17 08:47

package img

import (
	"fmt"
	"migrateDockerRegistries/env"
	"migrateDockerRegistries/helpers"
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
	fmt.Printf("Fetching image list from %s\n", helpers.Blue(registries.Source.Name))
	orgList, err = getRepoTags(registries.Source.URL)
	if err != nil {
		return helpers.CustomError{fmt.Sprintf("Unable to fetch %s repo/tags list: ",
			registries.Source.URL, err)}
	}

	// Fetches the dest registry's list
	fmt.Printf("Fetching image list from %s\n", helpers.Blue(registries.Dest.Name))
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
	fmt.Printf("\n%s\n", helpers.Blue("Comparing both lists"))
	finalList := compareLists(registries.Source.URL, orgList, destList)
	if err = saveListToFile(registries.Source.Name+"-"+registries.Dest.Name+".txt", finalList); err != nil {
		return helpers.CustomError{fmt.Sprintf("Unable to save the final list: %s", err)}
	}

	fmt.Printf("\nImage listing completed\n")
	return nil
}
